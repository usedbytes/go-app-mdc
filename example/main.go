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
		mdc.AppBarTop().
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
		app.Main().
			Class(mdc.AppBarMainClass).
			Body(
				mdc.List().
					Items(
						mdc.ListItem().
							Text("Item 1").
							Meta(
								app.Span().
								Class(mdc.ListItemMetaClass).
								Text("META"),
							),
						mdc.ListItem().
							Text("Item 2").
							Meta(
								app.I().
								Class("material-icons", mdc.ListItemMetaClass).
								Text("bookmark"),
							),
					),
				// Note that mixing item sizes in the same
				// list doesn't work properly:
				// https://github.com/material-components/material-components-web/issues/4209
				mdc.ListTwoLine().
					Items(
						mdc.ListItemTwoLine().
							Text("Two Line Item 1").
							SecondaryText("Secondary text").
							Graphic(app.I().Class("material-icons", mdc.ListItemGraphicClass).Text("thumb_up")),
						mdc.ListItemTwoLine().
							Text("Two Line Item 2").
							SecondaryText("Secondary text").
							Graphic(app.I().Class("material-icons", mdc.ListItemGraphicClass).Text("thumb_down")),
					),
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
