package main

import (
	"flag"
	_ "net/http/pprof" // nolint:gosec // The registration is safely forced to a randomly generated path
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/espal-digital-development/espal-core/runner"
	"github.com/juju/errors"
	_ "github.com/lib/pq"
)

// TODO :: 7777 Make some more tooling adjustable:
// - The HTTP engine
// - Email engines
// - The storages
// - The database (engine?)

func main() {
	path := flag.String("config-dir", "", "Destination directory where the config.yml can be found")
	flag.Parse()
	if err := Start(*path); err != nil {
		if _, err := spew.Println(errors.ErrorStack(err)); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
	select {}
}

// Start executes the full application runner.
func Start(path string) error {
	app := runner.New()
	// TODO :: 7777 Implement and test this:
	// - Themes, routes, pages
	// - Inject ful new services, storages, databases, forms, etc.
	// app.RegisterTheme()
	// app.RegisterRoute()
	return app.Run(path)
}
