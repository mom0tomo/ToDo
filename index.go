package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

var indexTmpl = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html>
	<head>
		<title>ToDo List</title>
	</head>
	<body>
	<h1>ToDo List</h1>
	<form action="/post">
		<input type="text" name="title">
		<input type="text" name="content">
		<input type="submit">
	</form>
	<div class="tasks">{{range .}}
		<div class="task">
			<h2 class="task-title">{{.Title}}</h2>
			<p class="task-content">{{.Content}}</p>
		</div>
	{{end}}</div>
	</body>
</html>`))
func index(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	var tasks []*Task

	q := datastore.NewQuery("Task").Order("-createdAt").Limit(10)
	if _, err := q.GetAll(ctx, &tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := indexTmpl.Execute(w, tasks); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
