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
							Text("Item 1").
							Meta(
								app.Span().
								Class(mdc.ListItemMetaClass).
								Text("META"),
							),
						mdc.NavListItem().
							Text("Item 2").
							Meta(
								app.I().
								Class("material-icons", mdc.ListItemMetaClass).
								Text("bookmark"),
							),
					),
				),
			app.Div().
				Class("mdc-drawer-app-content").
				Body(
					app.Main().
						Class(mdc.AppBarMainClass).
						Class("main-content").
						Body(
							// Note that mixing item sizes in the same
							// list doesn't work properly:
							// https://github.com/material-components/material-components-web/issues/4209
							mdc.ListTwoLine().Items(
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
