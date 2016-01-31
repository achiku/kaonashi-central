package main

import (
	"encoding/json"
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-zoo/bone"
	"github.com/gocraft/dbr"
)

func getNote(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	noteID := bone.GetValue(r, "id")
	sess := ctx.Value("db").(*DB).conn.NewSession(nil)
	var note Note
	sess.Select("id", "title", "body", "created", "updated").
		From("note").
		Where("id = ?", noteID).
		Load(&note)

	if note.Id == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	res := NoteResponse{Data: note}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func getNoteTitles(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	sess := ctx.Value("db").(*DB).conn.NewSession(nil)
	var titles []NoteTitle
	sess.Select("id", "title", "created", "updated").
		From("note").
		Load(&titles)
	res := NoteTitlesResponse{Data: titles}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func createNote(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	title := dbr.NewNullString(r.Form["title"][0])
	body := dbr.NewNullString(r.Form["body"][0])
	note := &Note{
		Title:   title,
		Body:    body,
		Created: time.Now(),
		Updated: time.Now(),
	}
	sess := ctx.Value("db").(*DB).conn.NewSession(nil)
	_, err := sess.InsertInto("note").
		Columns("title", "body", "created", "updated").
		Record(note).
		Exec()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := MessageResponse{Data: StatusMessage{Message: "created"}}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func deleteNote(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	noteID := bone.GetValue(r, "id")
	sess := ctx.Value("db").(*DB).conn.NewSession(nil)
	var note Note
	sess.Select("id").
		From("note").
		Where("id = ?", noteID).
		Load(&note)
	if note.Id == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	sess.DeleteFrom("note").Where(dbr.Eq("id", note.Id)).Exec()
	res := MessageResponse{Data: StatusMessage{Message: "deleted"}}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func updateNote(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	title := dbr.NewNullString(r.Form["title"][0])
	body := dbr.NewNullString(r.Form["body"][0])

	noteID := bone.GetValue(r, "id")
	sess := ctx.Value("db").(*DB).conn.NewSession(nil)
	var note Note
	sess.Select("id, title, body, created, updated").
		From("note").
		Where("id = ?", noteID).
		Load(&note)
	if note.Id == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	sess.Update("note").
		Set("title", title).
		Set("body", body).
		Set("updated", time.Now()).
		Where(dbr.Eq("id", note.Id)).
		Exec()
	res := MessageResponse{Data: StatusMessage{Message: "updated"}}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(res)
}
