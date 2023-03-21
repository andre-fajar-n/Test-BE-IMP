// Package cmd contains command line that can be used in this app
package cmd

import (
	"os"

	"test-be-IMP/internal"
	server "test-be-IMP/internal/app"
	"test-be-IMP/internal/pkg/migrations"

	"github.com/urfave/cli/v2"
)

// Start function to start cli
func Start() error {
	app := cli.NewApp()
	app.Name = internal.Get().Name

	app.Commands = []*cli.Command{
		{
			Name:  "migrations",
			Usage: "Migration Up",
			Action: func(ctx *cli.Context) error {
				if err := migrations.Up(); nil != err {
					return err
				}

				return nil
			},
		},
		{
			Name:  "rollback",
			Usage: "Migration Down",
			Action: func(ctx *cli.Context) error {
				if err := migrations.Down(); nil != err {
					return err
				}

				return nil
			},
		},
		{
			Name:  "drop",
			Usage: "Drop All Tables",
			Action: func(ctx *cli.Context) error {
				if err := migrations.Drop(); nil != err {
					return err
				}

				return nil
			},
		},
		{
			Name:  "start",
			Usage: "Start the Application",
			Action: func(ctx *cli.Context) error {
				return server.StartServer()
			},
		},
		{
			Name:  "launch",
			Usage: "Migration Up, Start Application",
			Action: func(ctx *cli.Context) error {
				if err := migrations.Up(); nil != err {
					return err
				}

				return server.StartServer()
			},
		},
	}

	return app.Run(os.Args)
}
