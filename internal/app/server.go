// Package app contains server of this app
package app

import (
	logger "log"
	"os"
	"test-be-IMP/gen/restapi"
	"test-be-IMP/gen/restapi/operations"
	"test-be-IMP/internal"
	"test-be-IMP/internal/app/handlers"
	"test-be-IMP/internal/app/rest"
	"test-be-IMP/pkg/db"
	"test-be-IMP/pkg/log"

	"github.com/casualjim/middlewares"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"
	"github.com/justinas/alice"
)

var mainFlags = struct {
	AppConfig string `long:"config" description:"Main application configuration YAML path"`
}{}

// start server
func StartServer() error {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Get().Err(err)
	}

	api := operations.NewTestBeIMPServerAPI(swaggerSpec)
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "App Flags",
			LongDescription:  "",
			Options:          &mainFlags,
		},
	}

	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Test BE IMP Open API"
	parser.LongDescription = "Test BE IMP Open API"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Get().Err(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	h := handlers.NewHandler(db.Open())

	rest.Route(api, h)

	rest.Authorization(api)

	logger := logger.New(os.Stdout, "Log -> ", logger.LstdFlags)

	handler := alice.New(
		middlewares.NewRecoveryMW("testbeIMPopenapi", nil),
		middlewares.NewProfiler,
	).Then(api.Serve(nil))

	server.SetHandler(handler)
	api.Logger = func(s string, i ...interface{}) {
		logger.Println(s, i)
	}
	server.Port = internal.Get().Port

	return server.Serve()
}
