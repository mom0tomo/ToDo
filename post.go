package main

import (
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func post(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	title := r.FormValue("title")

	content := r.FormValue("content")

	task := &Task {
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}

	key := datastore.NewIncompleteKey(ctx, "Task", nil)

	if _, err := datastore.Put(ctx, key, task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/", http.StatusFound)
}