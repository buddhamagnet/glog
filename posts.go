package main

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
	Extract string
	File  string
}

func (post Post) Teaser() template.HTML {
	if len(post.Body) >= 100 {
		return post.Body[:100] + "..."
	}
	return post.Body
}

func (post Post) Excerpt() string {
	if len(post.Extract) >= 100 {
		return post.Extract[:100] + "..."
	}
	return post.Extract
}

func GetPosts() []Post {
	a := []Post{}
	files, _ := filepath.Glob("posts/*")
	for i := len(files) - 1; i >= 0; i-- {
		file := strings.Replace(files[i], "posts/", "", -1)
		file = strings.Replace(file, ".md", "", -1)
		fileread, _ := ioutil.ReadFile(files[i])
		lines := strings.Split(string(fileread), "\n")
		a = append(a, NewPost(lines, file))
	}
	return a
}

func NewPost(lines []string, file string) Post {
	body := strings.Join(lines[2:], "\n")
	processed := string(blackfriday.MarkdownCommon([]byte(body)))
	post := Post{
		string(lines[0]),
		string(lines[1]),
		template.HTML(processed),
		body,
		file,
	}
	return post
}
