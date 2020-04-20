package stories

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.Template
var defaultHandlerTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Choose your own adventure</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
        <p>{{.}}</p>
    {{end}}
    <ul>
    {{range .Options}}
        <li><a href="/{{.Arc}}">{{.Text}}</a></li>
    {{end}}    
    </ul>
</body>
</html>
`

func NewHandler(s Story) http.Handler {
	return handler{s: s}
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, h.s["intro"])
	if err != nil {
		log.Fatal(err)
	}
}


type Story map[string]Chapter

type Chapter struct {
	Title   string    `json:"title"`
	Paragraphs   []string  `json:"story"`
	Options []Options `json:"options"`
}

type Options struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func JsonStory(s string) (Story, error) {
	file, err := ioutil.ReadFile(s)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var myStory Story
	err = json.Unmarshal(file, &myStory)
	if err != nil {
		return nil, err
	}

	return myStory, nil
}