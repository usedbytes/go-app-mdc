package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/usedbytes/go-app-mdc/mdc"
)

type listExample struct {
	app.Compo

	lists []app.UI
}

func ListExample() *listExample {
	return &listExample{
		lists: []app.UI{
			mdc.List().Items(
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
			mdc.ListTwoLine().Items(
				mdc.ListItemTwoLine().
					Text("Two Line Item 1").
					SecondaryText("Secondary text").
					Graphic(
						app.I().
						Class("material-icons", mdc.ListItemGraphicClass).
						Text("thumb_up"),
					),
				mdc.ListItemTwoLine().
					Text("Two Line Item 2").
					SecondaryText("Secondary text").
					Graphic(
						app.I().
						Class("material-icons", mdc.ListItemGraphicClass).
						Text("thumb_down"),
					),
			),
		},
	}
}

func (e *listExample) Render() app.UI {
	return app.Main().
		Class(mdc.AppBarMainClass).
		Class("main-content").
		Body(
			app.Range(e.lists).Slice(func (i int) app.UI {
				return e.lists[i]
			}),
		)
}
