package gfm_server

import (
	"bytes"
	"os"

	"text/template"

	"github.com/gofiber/fiber/v2"
	gfm "github.com/shurcooL/github_flavored_markdown"
)
func renderMarkdown(c *fiber.Ctx) error {
	markdownData, err := os.ReadFile(c.Path()[1:])
	if err != nil {
		return err
	}
	outBuf := bytes.Buffer{}
	output := gfm.Markdown(markdownData)

	tmpl, err := template.New("mdpage").ParseFS(templatesDir, "templates/mdpage.tmpl")
	if err != nil {
		return err
	}
	tmpl.ExecuteTemplate(&outBuf, "mdpage", string(output))

	c.Set("Content-Type", "text/html; charset=utf-8")
	c.Response().SetBodyRaw(outBuf.Bytes())

	return nil
}
