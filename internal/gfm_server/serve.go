package gfm_server

import (
	"embed"
	"fmt"
	"log"
	"path/filepath"
	"regexp"

	"github.com/shurcooL/github_flavored_markdown/gfmstyle"
	"github.com/urfave/cli/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed serve.min.css
var ServeCSS string

//go:embed templates/*
var templatesDir embed.FS

var excludeRegex = regexp.MustCompile(`^(?:\.git|.DS_Store|\.Trash)`)

func ServeGfm(c *cli.Context) error {
	port := c.Int("port")
	host := c.String("addr")
	directory := c.String("directory")

	full_path, err := filepath.Abs(directory)
	if err != nil {
		return err
	}

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		ServerHeader:          "GFM-Server",
		AppName:               "GFM-Server v0.1.0",
	})

	app.Static("/", directory, fiber.Static{
		Browse:         true,
		ModifyResponse: modifyResponse,
	})

	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: gfmstyle.Assets,
	}))

	fmt.Printf("Serving files from `%s` on http://%s:%d\n", full_path, host, port)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", host, port)))
	return nil
}
