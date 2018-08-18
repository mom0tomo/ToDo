package main

import "time"

type Task struct {
	ID int64  `datastore:"-"`
	Title string `datastore:"title"`
	Content string `datastore:"content"`
	CreatedAt time.Time `datastore:"createdAt"`
}