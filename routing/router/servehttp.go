package router

import (
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/espal-digital-development/espal-core/routing/router/contexts"
	"github.com/juju/errors"
)

// ServeHTTP functions as a callback for server routing binding.
func (httpRouter *HTTPRouter) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	var start time.Time
	if httpRouter.configService.Logging() {
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
	}(httpRouter, responseWriter, request, start)

	domain, ok, err := httpRouter.domainStore.GetOneActiveByHost(request.Host)
	if err != nil {
		context := httpRouter.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.RenderInternalServerError(errors.Trace(err))

		httpRouter.loggerService.Errorf("domain `%s` fetch threw an error `%s`. The data integrity has been violated", request.Host, err.Error())
		return
	}
	if !ok {
		context := httpRouter.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.SetStatusCode(http.StatusInternalServerError)
		context.RenderInternalServerError(errors.Trace(err))

		httpRouter.loggerService.Errorf("domain `%s` must exist. The data integrity has been violated", request.Host)
		return
	}
	if !domain.Active() {
		context := httpRouter.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.SetStatusCode(http.StatusServiceUnavailable)
		context.RenderNon200()
		return
	}
	site, ok, err := httpRouter.siteStore.GetOneOnlineByID(domain.SiteID())
	if err != nil {
		context := httpRouter.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.RenderInternalServerError(errors.Trace(err))

		httpRouter.loggerService.Errorf("site `%s` fetch threw an error `%s`. The data integrity has been violated", request.Host, err.Error())
		return
	}
	if !ok {
		context := httpRouter.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.RenderInternalServerError(errors.Trace(err))

		httpRouter.loggerService.Errorf("domain `%s` must have a site. The data integrity has been violated", request.Host)
		return
	}
	if !site.Online() {
		context := httpRouter.contextsFactory.NewContext(request, responseWriter, nil, nil)
		context.SetStatusCode(http.StatusServiceUnavailable)
		context.RenderNon200()
		return
	}

	context := httpRouter.contextsFactory.NewContext(request, responseWriter, domain, site)

	// TODO :: Make a config option that can disguise admin routes and make them appear as they don't exist (404)

	route, routeFound := httpRouter.getRoute(context.Path())
	if routeFound {
		// TODO :: 7 Assets for the Auth/Login page should also be allowed to be served
		if httpRouter.configService.SecurityGlobalAuthentication() && request.URL.Path != "/Auth" && !context.IsLoggedIn() {
			context.Redirect("/Auth", http.StatusTemporaryRedirect)
			return
		}

		route.Handle(context)
	} else {
		slug, ok, err := httpRouter.slugStore.GetOneByDomainIDAndPath(domain.ID(), context.Host())
		if err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			httpRouter.loggerService.Errorf("slug `%s` fetch threw an error `%s`. The data integrity has been violated", context.Host(), err.Error())
			return
		}
		var routeFound bool
		var route handler
		if ok {
			route, routeFound = httpRouter.getRoute(slug.Path())
		}
		if routeFound {
			// TODO :: 7 Test if it works with POSTing (non-GET requests) like /Login
			if httpRouter.configService.SecurityGlobalAuthentication() && !context.IsLoggedIn() {
				context.Redirect("/Auth", http.StatusTemporaryRedirect)
				return
			}
			route.Handle(context)
		} else {
			context.RenderNotFound()
		}
		// TODO :: Slug; bounce it against the actual route (it needs to start to honor being a slug tho; not suddenly go back from /Inloggen to /Login again)
	}

	if context.StatusCode() == 0 || context.StatusCode() < 300 {
		// TODO :: 7777 Maybe don't wait here and fire a routine to handle the save?
		if err := context.SaveSessionIfNeeded(); err != nil {
			context.RenderInternalServerError(errors.Trace(err))
			return
		}
	}

	if httpRouter.configService.Logging() {
		httpRouter.log(start, context)
	}
}

func (httpRouter *HTTPRouter) log(start time.Time, context contexts.RequestContext) {
	pathColor := 37
	statusColor := 32
	queryStringColor := 94

	statusCode := context.StatusCode()
	if statusCode == 0 {
		statusCode = 200
	}

	switch {
	case statusCode >= 500:
		statusColor = 31
	case statusCode >= 400:
		statusColor = 33
	case statusCode >= 300:
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

	httpRouter.loggerService.Customf("\033[0;%dm%d\033[m %6s %12v \033[0;%dm%s\033[m\033[0;%dm%s\033[m", func(s string) string {
		return s
	}, statusColor, statusCode, context.Method(), float64(time.Since(start).Nanoseconds())/1e3, pathColor, context.Path(), queryStringColor, queryString)
}
