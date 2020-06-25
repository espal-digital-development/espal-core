package runner

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Listen for syscalls to execute wind-down calls.
func (r *Runner) listenToSystemSignals() {
	go func(appRunner *Runner) {
		sigc := make(chan os.Signal, 1)
		signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		sig := <-sigc
		if sig == os.Interrupt {
			appRunner.services.logger.Infof("Received signal `%s`", sig.String())
			appRunner.services.logger.Infof("The Server ran for %v", time.Since(appRunner.serverStartTime))
		}
		os.Exit(0)
	}(r)
}
