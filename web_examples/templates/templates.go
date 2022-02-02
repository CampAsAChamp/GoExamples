package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

// The html/template package provides templating language for HTMl templates
// It is mostly use in web applications to display data in a structured way in a client's browswer
// One benefit of Go's templating language is the automatic escaping of data
// There is no need to worry about XSS attacks as Go parses the HTML template and escapes all inputs before displaying it to the browser
func main() {
	// Templates can either be parsed from a string or a file on disk
	tmplate, _ := template.ParseFiles("layout.html")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		data := TodoPageData{
			PageTitle: "My TODO List",
			Todos: []Todo{
				{Title: "Task 1", Done: true},
				{Title: "Task 2", Done: false},
				{Title: "Task 3", Done: false},
			},
		}

		// Execute a Template in a Request Handler
		// Once a template is parsed from disk it's ready to be used in the request handler
		// When the function is called on a http.ResponseWriter the Content-type in the header is automatically set in teh HTTP reponse to 'content type: text/html; charset=utf-8'
		tmplate.Execute(rw, data)
	})

	// This shows a TODO list written as an unordered list in HTML
	// When rendering templates, the data passed in can be any kind of Go data structure
	// To access the data in a template the top most variable is accessed by {{.}}
	// The dot inside the curly braces is called the pipeline and the root element of the data
	/*
		<h1>{{.PageTitle}}</h1>
		<ul>
			{{range .Todos}}
				{{if .Done}}
					<li class = "done">{{.Title}}</li>
				{{else}}
					<li>{{.Title}}</li>
				{{end}}
			{{end}}
		</ul>
	*/

	http.ListenAndServe(":80", nil)
	fmt.Println("End of Main")
}
