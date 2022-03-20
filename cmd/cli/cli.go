package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	defaultCliConfFiles = []string{"./cli_conf.json", "/etc/cli_conf.json"}
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{

		{
			Name:     "schema",
			Usage:    "options for fundb schmea opetaion",
			Category: "schema",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new schema",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "del",
					Usage: "del a schema",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "fetch",
					Usage: "feth a schema meta",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},

		{
			Name:     "kv",
			Aliases:  []string{"t", "tmpl"},
			Usage:    "options for task templates",
			Category: "kv",
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a for data in schema ",
					Action: func(c *cli.Context) error {
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "del",
					Usage: "remove a data from schema",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "get",
					Usage: "get a data from schema",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}

	app.Name = "fundb_cli"
	app.Usage = ""
	app.Description = "fundb_cli for fundbserver"
	app.Version = "v1.0.1"

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
