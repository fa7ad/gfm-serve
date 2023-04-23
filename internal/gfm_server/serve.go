package gfm_server

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
)

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
		ModifyResponse: decideMdOrFile,
	})

	fmt.Printf("Serving files from `%s` on http://%s:%d\n", full_path, host, port)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", host, port)))
	return nil
}

func decideMdOrFile(c *fiber.Ctx) error {
	body := c.Response().Body()
	newbody := string(body)
	mtime := regexp.MustCompile(`, (?:file|dir),.*?last modified.*?\<`)

	newbody = mtime.ReplaceAllString(newbody, `<`)

	c.Send([]byte(newbody))
	return nil
}
