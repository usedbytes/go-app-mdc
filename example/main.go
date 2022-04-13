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
	clickHandler := func(c app.Context, e app.Event) {
		app.Log(c.JSSrc().Get("id"), "clicked")
	}

	return app.Div().Body(
		app.Div().Style("padding", "5px").Body(
			mdc.Button().
				ID("text-button").
				Label("Text Button").
				OnClick(clickHandler),
		),
		app.Div().Style("padding", "5px").Body(
			mdc.Button().
				Label("Outlined Button").
				Outlined(true).
				OnClick(clickHandler),
		),
		app.Div().Style("padding", "5px").Body(
			mdc.Button().
				Label("With icon").
				Outlined(true).
				LeadingIcon("favorite").
				OnClick(clickHandler),
		),
		app.Div().Style("padding", "5px").Body(
			mdc.Button().
				Label("Trailing icon").
				Outlined(true).
				TrailingIcon("settings").
				OnClick(clickHandler),
		),
		app.Div().Style("padding", "5px").Body(
			mdc.Button().
				Label("Raised").
				Raised(true).
				OnClick(clickHandler),
		),
		app.Div().Style("padding", "5px").Body(
			mdc.Button().
				Label("Both icons").
				Raised(true).
				LeadingIcon("chevron_left").
				TrailingIcon("chevron_right").
				OnClick(clickHandler),
		),
		app.Div().Style("padding", "5px").Body(
			mdc.IconButton().
				Icon("favorite").
				OnClick(clickHandler),
			mdc.IconButtonToggle().
				IconOff("visibility_off").
				IconOn("visibility_on").
				OnClick(clickHandler),
			mdc.IconButtonToggle().
				On(true).
				IconOff("arrow_back").
				IconOn("arrow_forward").
				OnClick(clickHandler),
		),
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
