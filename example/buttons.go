package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/usedbytes/go-app-mdc/mdc"
)

type buttonExample struct {
	app.Compo
}

func (e *buttonExample) Render() app.UI {
	clickHandler := func(c app.Context, e app.Event) {
		app.Log(c.JSSrc().Get("id"), "clicked")
	}

	return app.Main().
		Class(mdc.AppBarMainClass).
		Class("main-content").
		Body(
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
