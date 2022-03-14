package story

import (
	"fmt"
	"net/http"
	"path/filepath"
	"text/template"
)

type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func (s *Story) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file := filepath.Join(fmt.Sprintf("%s/src/templates", filepath.Dir("story.html")), "story.html")
	tmpl := template.Must(template.ParseFiles(file))
	tmpl.Execute(w, s)
}
