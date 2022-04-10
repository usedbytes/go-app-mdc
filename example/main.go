package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/usedbytes/go-app-mdc/mdc"
)

type example struct {
	app.Compo

	buttonText string
}

func (e *example) Render() app.UI {
	return app.Div().Body(
		&mdc.Button{ Label: "Text Button" },
		&mdc.Button{ Label: "Outlined Button", Outlined: true },
		&mdc.Button{ Label: "Icon Button", Icon: "favorite" },
		&mdc.Button{ Label: "Raised Icon Button", Raised: true, Icon: "warning" },
	)
}

func run() error {
	app.Route("/", &example{})
	app.RunWhenOnBrowser()

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/").Handler(&app.Handler{
		Name:        "MDC Example",
		Description: "Example app for using MDC components",
		Scripts: []string{
			"https://unpkg.com/material-components-web@latest/dist/material-components-web.min.js",
		},
		Styles: []string{
			"/web/style.css",
			"https://fonts.googleapis.com/icon?family=Material+Icons",
			"https://unpkg.com/material-components-web@latest/dist/material-components-web.min.css",
		},
	})

	httpSrv := &http.Server{
		Handler: router,
		Addr:    ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	return httpSrv.ListenAndServe()
}

func main() {
	err := run()
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		os.Exit(1)
	}
}
