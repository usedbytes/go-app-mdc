package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/usedbytes/go-app-mdc/mdc"
)

type buttonExample struct {
	app.Compo

	eyeToggle bool
	arrowToggle bool
	buttons []app.UI
}

func ButtonExample() *buttonExample {
	ex := new(buttonExample)

	clickHandler := func(c app.Context, e app.Event) {
		app.Log(c.JSSrc().Get("id"), "clicked")
	}

	eyeButton := mdc.IconButtonToggle().
		On(ex.eyeToggle).
		IconOff("visibility_off").
		IconOn("visibility_on")

	eyeButton.OnClick(func(c app.Context, e app.Event) {
		ex.eyeToggle = !ex.eyeToggle
		eyeButton.On(ex.eyeToggle)
		clickHandler(c, e)
	})

	arrowButton := mdc.IconButtonToggle().
		On(ex.arrowToggle).
		IconOff("arrow_back").
		IconOn("arrow_forward")

	arrowButton.OnClick(func(c app.Context, e app.Event) {
		ex.arrowToggle = !ex.arrowToggle
		arrowButton.On(ex.arrowToggle)
		clickHandler(c, e)
	})

	ex.buttons = []app.UI{
		mdc.Button().
			ID("text-button").
			Label("Text Button").
			OnClick(clickHandler),
		mdc.Button().
			Label("Outlined Button").
			Outlined(true).
			OnClick(clickHandler),
		mdc.Button().
			Label("With icon").
			Outlined(true).
			LeadingIcon("favorite").
			OnClick(clickHandler),
		mdc.Button().
			Label("Trailing icon").
			Outlined(true).
			TrailingIcon("settings").
			OnClick(clickHandler),
		mdc.Button().
			Label("Raised").
			Raised(true).
			OnClick(clickHandler),
		mdc.Button().
			Label("Both icons").
			Raised(true).
			LeadingIcon("chevron_left").
			TrailingIcon("chevron_right").
			OnClick(clickHandler),
		app.Div().Body(
			mdc.IconButton().
				Icon("favorite").
				OnClick(clickHandler),
			eyeButton,
			arrowButton,
		),
		mdc.FAB().
			ID("fab").
			Icon(app.I().
				Class("material-icons", mdc.FABIconClass).
				Text("code"),
			),
	}

	return ex
}

func (e *buttonExample) Render() app.UI {
	return app.Main().
		Class(mdc.AppBarMainClass).
		Class("main-content").
		Body(
			app.Range(e.buttons).Slice(func(i int) app.UI {
				return app.Div().
					Style("padding", "5px").
					Body(
						e.buttons[i],
					)
			}),
		)
}
