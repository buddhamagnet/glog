package models

import (
	"github.com/russross/blackfriday"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Post struct {
	Title string
	Date  string
	Body  template.HTML
	File  string
}

func (post Post) Teaser() template.HTML {
	return post.Body[:100] + "..."
}

func GetPosts() []Post {
	a := []Post{}
	files, _ := filepath.Glob("posts/*")
	for _, f := range files {
		file := strings.Replace(f, "posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fileread, _ := ioutil.ReadFile(f)
		lines := strings.Split(string(fileread), "\n")
		a = append(a, NewPost(lines, file))
	}
	return a
}

func NewPost(lines []string, file string) Post {
	body := strings.Join(lines[2:], "\n")
	body = string(blackfriday.MarkdownCommon([]byte(body)))
	post := Post{
		string(lines[0]),
		string(lines[1]),
		template.HTML(body),
		file,
	}
	return post
}
