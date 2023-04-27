package gfm_server

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	_ "embed"

	"github.com/antchfx/htmlquery"
	"github.com/gofiber/fiber/v2"
	"github.com/urfave/cli/v2"
)

//go:embed serve.min.css
var ServeCSS string
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
		ModifyResponse: serveMiddleware,
	})

	fmt.Printf("Serving files from `%s` on http://%s:%d\n", full_path, host, port)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", host, port)))
	return nil
}

func serveMiddleware(c *fiber.Ctx) error {
	mime := c.GetRespHeader("Content-Type")
	if !strings.Contains(mime, "text/html") {
		return nil
	}

	_body := c.Response().Body()
	body := make([]byte, len(_body))
	copy(body, _body)

	c.Response().ResetBody()

	document, err := htmlquery.Parse(bytes.NewReader(body))
	if err != nil {
		return err
	}

	styleTag := h("style", nil, s(ServeCSS))

	h1, err := htmlquery.Query(document, "//h1")
	if err != nil || h1 == nil {
		return err
	}

	bodyPtr := h1.Parent

	bodyPtr.AppendChild(styleTag)

	bodyPtr.AppendChild(h("header", nil, h1))

	h1txt := h1.LastChild.Data
	h1.LastChild.Data = ""
	itag := h("i", nil, s("Index of "))

	anchor := h("a", a(map[string]string{"href": "#"}), s(h1txt))

	h1.AppendChild(itag)
	h1.AppendChild(anchor)

	ul, err := htmlquery.Query(document, "//ul")
	if err != nil || ul == nil {
		return err
	}

	bodyPtr.AppendChild(h("main", nil, ul))

	listItems, err := htmlquery.QueryAll(document, "//li")
	if err != nil {
		return err
	}

	for _, listItem := range listItems {
		if excludeRegex.MatchString(listItem.FirstChild.LastChild.Data) {
			listItem.Parent.RemoveChild(listItem)
			continue
		}

		if strings.HasPrefix(listItem.LastChild.Data, ",") {
			listItem.LastChild.Data = ""
		}
	}

	output := htmlquery.OutputHTML(document, true)
	c.Response().SetBody([]byte(output))
	return nil
}
