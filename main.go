package main

//go:generate go get -u github.com/valyala/quicktemplate/qtc
//go:generate qtc -dir=./pages

import (
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof" // nolint:gosec // The registration is safely forced to a randomly generated path
	"os"

	"github.com/espal-digital-development/espal-core/app/modules/core"
	"github.com/espal-digital-development/espal-core/runner"
	"github.com/juju/errors"
	_ "github.com/lib/pq"
)

func main() {
	path := flag.String("config-dir", "", "Destination directory where the config.yml can be found")
	flag.Parse()
	if err := Start(*path); err != nil {
		if _, err := fmt.Println(errors.ErrorStack(err)); err != nil {
			panic(err)
		}
		os.Exit(1)
	}
	select {}
}

// Start executes the full application runner.
func Start(path string) error {
	app, err := runner.New()
	if err != nil {
		log.Fatal(err)
	}
	app.SetPath(path)

	// TODO :: 777777 :: This registration goes into the actual app style
	coreModule, err := core.New()
	if err != nil {
		log.Fatal(err)
	}
	if err := app.RegisterModule(coreModule); err != nil {
		log.Fatal(err)
	}

	return app.RunNonBlocking()
}
