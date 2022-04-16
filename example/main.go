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

	content app.UI
	buttons buttonExample
	lists   listExample
}

func (e *example) Render() app.UI {
	return app.Div().
		ID("container").
		Body(
			mdc.AppBarTop().
				ID("app-bar").
				Title("MDC Example").
				NavIcon(
					mdc.IconButton().
						Class(mdc.AppBarNavigationClass).
						Icon("menu"),
				).Actions(
					mdc.IconButton().
						Class(mdc.AppBarActionClass).
						Icon("help"),
					mdc.IconButton().
						Class(mdc.AppBarActionClass).
						Icon("settings"),
				),
			mdc.Drawer().
				Class(mdc.AppBarMainClass).
				Title("Nav Title").
				Subtitle("Nav subtitle").
				Content(
					mdc.NavList().Items(
						mdc.NavListItem().
							Text("Buttons").
							Meta(
								app.Span().
								Class(mdc.ListItemMetaClass).
								Text("META"),
							).
							OnClick(func(ctx app.Context, ev app.Event) {
								e.content = &e.buttons
								e.Update()
							}),
						mdc.NavListItem().
							Text("Lists").
							Meta(
								app.I().
								Class("material-icons", mdc.ListItemMetaClass).
								Text("bookmark"),
							).
							OnClick(func(ctx app.Context, ev app.Event) {
								e.content = &e.lists
								e.Update()
							}),
					),
				),
			app.Div().
				Class("mdc-drawer-app-content").
				Body(
					e.content,
				),
	)
}

// Styles needed to lay out the content and navigation drawer, see:
// https://github.com/material-components/material-components-web/tree/master/packages/mdc-drawer
const rawStyles = `
<style>
#container {
  display: flex;
  height: 100vh;
}

.mdc-drawer-app-content {
  flex: auto;
  overflow: auto;
  position: relative;
}

.main-content {
  overflow: auto;
  height: 100%;
}

.mdc-top-app-bar {
  position: absolute;
  z-index: 7;
}
</style>
`

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
		RawHeaders: []string{
			rawStyles,
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
