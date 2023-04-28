package gfm_server

import (
	"bytes"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/gofiber/fiber/v2"
)

func modifyResponse(c *fiber.Ctx) error {
	path := c.Path()
	if strings.HasSuffix(path, ".md") {
		return renderMarkdown(c)
	}
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

	title, err := htmlquery.Query(document, "//title")
	if err != nil || title == nil {
		return err
	}

	title.LastChild.Data = `Index of ` + title.LastChild.Data

	h1, err := htmlquery.Query(document, "//h1")
	if err != nil || h1 == nil {
		return err
	}

	bodyPtr := h1.Parent

	bodyPtr.AppendChild(h("style", nil, s(ServeCSS)))
	bodyPtr.AppendChild(h("header", nil, h1))

	h1txt := h1.LastChild.Data
	h1.LastChild.Data = ""

	h1.AppendChild(h("i", nil, s("Index of ")))
	h1.AppendChild(h("a", a(map[string]string{"href": "#"}), s(h1txt)))

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
