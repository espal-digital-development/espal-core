package router

import (
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/espal-digital-development/espal-core/stores/setting"
	"github.com/juju/errors"
)

const dataIntegrityViolated = "The data integrity has been violated"

// ServeHTTP functions as a callback for server routing binding.
// nolint:funlen,gocognit
func (r *HTTPRouter) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	var start time.Time
	if r.configService.Logging() {
		start = time.Now()
	}

	defer func(router *HTTPRouter, respWriter http.ResponseWriter, req *http.Request, startTime time.Time) {
		if rec := recover(); rec != nil {
			router.loggerService.Errorf("recovered from Dispatch %s\n%s", rec, string(debug.Stack()))
			context := router.contextsFactory.NewContext(req, respWriter, nil, nil)
			context.RenderInternalServerError(errors.Errorf("recovered from Dispatch %s\n%s", rec, string(debug.Stack())))

			router.loggerService.Error(string(debug.Stack()))
			router.log(startTime, context)
		}
	}(r, responseWriter, request, start)

	domain, ok, err := r.domainStore.GetOneActiveByHost(request.Host)
	if err != nil {
		context := r.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.RenderInternalServerError(errors.Trace(err))

		r.loggerService.Errorf("domain `%s` fetch threw an error `%s`. %s", request.Host, err.Error(),
			dataIntegrityViolated)
		return
	}
	if !ok {
		context := r.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.SetStatusCode(http.StatusInternalServerError)
		context.RenderInternalServerError(errors.Trace(err))

		r.loggerService.Errorf("domain `%s` must exist. %s", request.Host, dataIntegrityViolated)
		return
	}
	if !domain.Active() {
		context := r.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.SetStatusCode(http.StatusServiceUnavailable)
		context.RenderNon200()
		return
	}
	site, ok, err := r.siteStore.GetOneOnlineByID(domain.SiteID())
	if err != nil {
		context := r.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.RenderInternalServerError(errors.Trace(err))

		r.loggerService.Errorf("site `%s` fetch threw an error `%s`. %s", request.Host, err.Error(),
			dataIntegrityViolated)
		return
	}
	if !ok {
		context := r.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.RenderInternalServerError(errors.Trace(err))

		r.loggerService.Errorf("domain `%s` must have a site. %s", request.Host, dataIntegrityViolated)
		return
	}
	if !site.Online() {
		context := r.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.SetStatusCode(http.StatusServiceUnavailable)
		context.RenderNon200()
		return
	}

	context := r.contextsFactory.NewContext(request, responseWriter, domain, site)

	theme, err := r.settingStore.GetOneForSite(setting.SettingTheme, "", domain.ID(), site.ID())
	if err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}
	if err := context.SetTheme(theme); err != nil {
		context.RenderInternalServerError(errors.Trace(err))
		return
	}

	// TODO :: Make a config option that can disguise admin routes and make them appear as they don't exist (404)

	if request.Method == http.MethodOptions {
		context.SetHeader("Access-Control-Allow-Origin", domain.Host())
		context.SetHeader("Access-Control-Allow-Headers", "Authorization")
		context.SetStatusCode(http.StatusOK)
		if r.configService.Logging() {
			r.log(start, context)
		}
		return
	}

	// TODO :: 777777 The Access-Control-Allow-Origin/Credentials, Content-Security-Policy should be a field on the
	// Config db object.
	context.SetHeader("Referrer-Policy", r.configService.SecurityHTTPReferrerPolicy())
	context.SetHeader("Content-Security-Policy", r.configService.SecurityHTTPContentSecurityPolicy())
	context.SetHeader("Access-Control-Allow-Origin", domain.Host())

	route, routeFound := r.getRoute(context.Path())
	if routeFound { // nolint:nestif
		// TODO :: 7 Assets for the Auth/Login page should also be allowed to be served
		if r.configService.SecurityGlobalAuthentication() && request.URL.Path != "/Auth" && !context.IsLoggedIn() {
			context.Redirect("/Auth", http.StatusTemporaryRedirect)
			return
		}
		route.Handle(context)
	} else {
		slug, ok, err := r.slugStore.GetOneByDomainIDAndPath(domain.ID(), context.Path())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			r.loggerService.Errorf("slug `%s` fetch threw an error `%s`. The data integrity has been violated",
				context.Path(), err.Error())
			return
		}

		var routeFound bool
		var route handler
		if ok {
			route, routeFound = r.getRoute("/" + slug.RerouteTo())
		}
		if routeFound {
			// TODO :: 7 Test if it works with POSTing (non-GET requests) like /Login
			if r.configService.SecurityGlobalAuthentication() && !context.IsLoggedIn() {
				context.Redirect("/Auth", http.StatusTemporaryRedirect)
				return
			}
			route.Handle(context)
		} else {
			context.RenderNotFound()
		}
		// TODO :: Slug; bounce it against the actual route (it needs to start to honor being a slug tho;
		// not suddenly go back from /Inloggen to /Login again).
	}

	if context.StatusCode() < http.StatusMultipleChoices {
		// TODO :: 7777 Maybe don't wait here and fire a routine to handle the save? But might cause not being saved
		// yet on next page load.
		if err := context.SaveSessionIfNeeded(); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}

	if r.configService.Logging() {
		if r.configService.Development() && strings.HasPrefix(context.Path(), "/health") {
			return
		}
		r.log(start, context)
	}
}

func (r *HTTPRouter) log(start time.Time, context contexts.RequestContext) {
	pathColor := 37
	statusColor := 32
	queryStringColor := 94

	statusCode := context.StatusCode()
	if statusCode == 0 {
		statusCode = http.StatusOK
	}

	switch {
	case statusCode >= http.StatusInternalServerError:
		statusColor = 31
	case statusCode >= http.StatusBadRequest:
		statusColor = 33
	case statusCode >= http.StatusMultipleChoices:
		statusColor = 35
	}

	if !strings.Contains(context.RequestURI(), ".") {
		pathColor = 34
	}

	var queryString []byte
	if len(context.QueryString()) > 0 {
		queryString = make([]byte, 0, len(context.QueryString())+1)
		queryString = append(queryString, '?')
		queryString = append(queryString, context.QueryString()...)
	}

	r.loggerService.Customf("\033[0;%dm%d\033[m %6s %13.3f \033[0;%dm%s\033[m\033[0;%dm%s\033[m", func(s string) string {
		return s
	}, statusColor, statusCode, context.Method(), float64(time.Since(start).Nanoseconds())/1e3, pathColor, // nolint:gomnd
		context.Path(), queryStringColor, queryString)
}
