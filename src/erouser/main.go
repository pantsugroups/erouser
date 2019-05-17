package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

func init() {
	// 想了想这个似乎用不上？
}

func main() {
	app := cli.NewApp()
	app.Name = "Ero User API"
	app.Usage = "a command user api from ero.ink"
	app.Commands = []cli.Command{
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "run the server.",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "confuse",
					Value:       "",
					Usage:       "Confuse the backend url.Such as /Confuse/api",
					Destination: &confuse,
				},
				cli.StringFlag{
					Name:        "port",
					Value:       "80",
					Usage:       "web server listening port.",
					Destination: &port,
				},
			},
			Action: func(c *cli.Context) error {
				StartEchoHandle()
				return nil
			},
		},
		{
			Name:    "Hijack",
			Aliases: []string{"h"},
			Usage:   "Hijack the http  and database handle.Please use other Execute to quote as package.",
			Action: func(c *cli.Context) error {
				Hijack(nil, nil)
				log.Println("The Function Prototype:\n\tfunc Hijack(db2 *gorm.DB,e2 *echo.Echo){ ... }\n" +
					"And you must used LoadRoutes() to Load the routes for your Echo Routes")
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
