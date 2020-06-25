package runner

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

// TODO :: 777 Reuse port again? And split for Windows too

type redirectRouter struct {
	serverHost string
	serverPort int
}

// ServeHTTP registers all the required redirect invokers.
func (r *redirectRouter) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// TODO :: 77 Redirect to default port should be left out? And what about internal ports vs external normal ports?
	if len(request.URL.RequestURI()) == 1 && request.URL.RequestURI()[0] == '/' {
		http.Redirect(response, request, fmt.Sprintf("https://%s:%d", r.serverHost, r.serverPort), http.StatusMovedPermanently)
	} else {
		http.Redirect(response, request, fmt.Sprintf("https://%s:%d%s", r.serverHost, r.serverPort, request.URL.RequestURI()), http.StatusMovedPermanently)
	}
}

// Listen to non-TLS to redirect to the TLS version.
func (r *Runner) startRedirectNonTLSServer() {
	go func(appRunner *Runner) {
		server := &http.Server{
			Addr: fmt.Sprintf("%s:%d", appRunner.services.config.ServerHost(), appRunner.services.config.ServerHTTPRedirectPort()),
			Handler: &redirectRouter{
				serverHost: appRunner.services.config.ServerHost(),
				serverPort: appRunner.services.config.ServerPort(),
			},
		}

		if err := server.ListenAndServe(); err != nil {
			appRunner.services.logger.Errorf("error in server.ListenAndServe: %s", err)
		}
	}(r)
}

func (r *Runner) startTLSServer() {
	go func(appRunner *Runner) {
		cfg := &tls.Config{
			MinVersion:               tls.VersionTLS13, // TODO :: wrk needs 12, but should use 13 for running
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
		}
		server := &http.Server{
			Addr:         fmt.Sprintf("%s:%d", appRunner.services.config.ServerHost(), appRunner.services.config.ServerPort()),
			Handler:      appRunner.services.router,
			TLSConfig:    cfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
		}

		if err := server.ListenAndServeTLS(appRunner.services.config.ServerSSLCertificateFilePath(), appRunner.services.config.ServerSSLKeyFilePath()); err != nil {
			appRunner.services.logger.Errorf("error in server.ListenAndServeTLS: %s", err)
		}
	}(r)

	r.services.logger.Infof("Server running TLS on `%s` port %d and redirecting non-TLS from port %d.", r.services.config.ServerHost(), r.services.config.ServerPort(), r.services.config.ServerHTTPRedirectPort())
}
