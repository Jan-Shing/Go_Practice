package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"./Story"
)

func main(){
	
	filename := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := StoryTool.JanStory(f)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.New("").Parse(storyTmpl))

	handler := StoryTool.NewHandler(story, StoryTool.Use_template(tmpl), StoryTool.Use_path(pathfn))

	mux := http.NewServeMux()
	mux.Handle("/story/", handler)
	mux.Handle("/", StoryTool.NewHandler(story))

	fmt.Printf("Starting the server on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

}

func pathfn(r *http.Request)string{
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}

var storyTmpl = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Content}}
        <p>{{.}}</p>
      {{end}}
      <ul>
      {{range .Opt}}
        <li><a href="/story/{{.Arc}}">{{.Text}}</a></li>
      {{end}}
      </ul>
    </section>
    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FCF6FC;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #797;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: underline;
        color: #555;
      }
      a:active,
      a:hover {
        color: #222;
      }
      p {
        text-indent: 1em;
      }
    </style>
  </body>
</html>`