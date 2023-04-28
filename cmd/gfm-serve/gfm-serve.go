package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/fa7ad/gfm-serve/internal/gfm_serve"
)

func main() {
	app := &cli.App{
		Name:  "gfm-serve",
		Usage: "Serve GitHub Flavored Markdown files",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p", "P"},
				Value:   8080,
				Usage:   "Port number to listen on",
			},
			&cli.StringFlag{
				Name:    "directory",
				Aliases: []string{"d", "D"},
				Value:   ".",
				Usage:   "Path to serve files from",
			},
			&cli.StringFlag{
				Name:    "addr",
				Aliases: []string{"a", "A"},
				Value:   "localhost",
				Usage:   "Hostname/Address to listen on",
			},
		},
		Action:                 gfm_serve.ServeGfm,
		UseShortOptionHandling: true,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
