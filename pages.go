package main

import (
	"github.com/russross/blackfriday"
	"html/template"
	"strings"
)

type Page struct {
	Title string
	Body  template.HTML
	File  string
}

func NewPage(lines []string, file string) Page {
	body := strings.Join(lines[1:], "\n")
	body = string(blackfriday.MarkdownCommon([]byte(body)))
	page := Page{
		string(lines[0]),
		template.HTML(body),
		file,
	}
	return page
}
