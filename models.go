package main

import (
	"time"

	"github.com/gocraft/dbr"
)

type Note struct {
	Id      int64          `json:"id" db:id`
	Title   dbr.NullString `json:"title" db:title`
	Body    dbr.NullString `json:"body" db:body`
	Created time.Time      `json:"created" db:created`
	Updated time.Time      `json:"updated" db:updated`
}

type NoteTitle struct {
	Id      int64          `json:"id"`
	Title   dbr.NullString `json:"title"`
	Created time.Time      `json:"created"`
	Updated time.Time      `json:"updated"`
}
