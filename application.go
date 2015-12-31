package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/getsentry/raven-go"
	"github.com/go-zoo/bone"
	"github.com/rs/xhandler"
	"golang.org/x/net/context"
)

func main() {
	// set up root context
	appConfig, err := NewAppConfig("./configs/development.toml")
	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}
	db, err := NewDB(appConfig)
	if err != nil {
		log.Fatalf("failed to open database: %s", err)
	}
	raven.SetDSN(os.Getenv("SENTRY_DSN"))
	rootCtx := context.Background()
	rootCtx = context.WithValue(rootCtx, "config", appConfig)
	rootCtx = context.WithValue(rootCtx, "db", db)

	// middleware chaining
	c := xhandler.Chain{}
	c.Use(recoverMiddleware)
	c.Use(loggingMiddleware)
	c.UseC(func(next xhandler.HandlerC) xhandler.HandlerC {
		return requestIdMiddleware{next: next}
	})
	c.UseC(xhandler.CloseHandler)
	c.UseC(xhandler.TimeoutHandler(2 * time.Second))

	// application routing
	mux := bone.New()
	mux.Get("/note", c.HandlerCtx(rootCtx, xhandler.HandlerFuncC(getNoteTitles)))
	mux.Get("/note/:id", c.HandlerCtx(rootCtx, xhandler.HandlerFuncC(getNote)))
	mux.Delete("/note/:id", c.HandlerCtx(rootCtx, xhandler.HandlerFuncC(deleteNote)))
	mux.Put("/note/:id", c.HandlerCtx(rootCtx, xhandler.HandlerFuncC(updateNote)))
	mux.Post("/note", c.HandlerCtx(rootCtx, xhandler.HandlerFuncC(createNote)))
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
