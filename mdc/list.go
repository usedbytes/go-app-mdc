package mdc

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// https://material.io/components/app-bars-top/web#using-the-top-app-bar
// Note that this is the "deprecated" list, so that it can be used in
// other components like the navigation drawer.
// https://github.com/material-components/material-components-web/tree/master/packages/mdc-list

// Note that mixing item sizes in the same list doesn't work properly:
// https://github.com/material-components/material-components-web/issues/4209

const (
	ListItemGraphicClass = "mdc-deprecated-list-item__graphic"
	ListItemMetaClass = "mdc-deprecated-list-item__meta"
)

type list struct {
	app.Compo

	id string
	mdcComponent app.Value
	linesClass string

	Iitems []app.UI
}

type IList interface {
	app.UI

	ID(string) IList
	Items(...app.UI) IList
}

func (l *list) ID(id string) IList {
	l.id = id
	return l
}

func (l *list) Items(i ...app.UI) IList {
	l.Iitems = i
	return l
}

func (l *list) OnMount(ctx app.Context) {
	if l.id == "" {
		l.id = fmt.Sprintf("list-%d", allocID())
	}

	l.mdcComponent = app.Window().
		Get("mdc").
		Get("list").
		Get("MDCList").
		Call("attachTo", ctx.JSSrc())

	app.Log("mounted", l.id, l.mdcComponent)
}

func (l *list) OnDismount(ctx app.Context) {
	l.mdcComponent.Call("destroy")
}

func (l *list) Render() app.UI {
	list := app.Ul().
		Class("mdc-deprecated-list", l.linesClass).
		ID(l.id)

	return list.Body(
		l.Iitems...
	)
}

func List() IList {
	return &list{}
}

func ListTwoLine() IList {
	return &list{
		linesClass: "mdc-deprecated-list--two-line",
	}
}

type listItem struct {
	app.Compo

	id string
	mdcComponent app.Value

	Itext string
	Igraphic app.UI
	Imeta app.UI
}

type IListItem interface {
	app.UI

	ID(string) IListItem
	Text(string) IListItem
	Graphic(app.UI) IListItem
	Meta(app.UI) IListItem
}

func (l *listItem) ID(id string) IListItem {
	l.id = id
	return l
}

func (l *listItem) Text(t string) IListItem {
	l.Itext = t
	return l
}

func (l *listItem) Graphic(g app.UI) IListItem {
	l.Igraphic = g
	return l
}

func (l *listItem) Meta(m app.UI) IListItem {
	l.Imeta = m
	return l
}

func (l *listItem) OnMount(ctx app.Context) {
	if l.id == "" {
		l.id = fmt.Sprintf("list-item-%d", allocID())
	}

	l.mdcComponent = app.Window().
		Get("mdc").
		Get("ripple").
		Get("MDCRipple").
		Call("attachTo", ctx.JSSrc())

	app.Log("mounted", l.id, l.mdcComponent)
}

func (l *listItem) OnDismount(ctx app.Context) {
	l.mdcComponent.Call("destroy")
}

func (l *listItem) Render() app.UI {
	item := app.Li().
		Class("mdc-deprecated-list-item").
		ID(l.id)

	return item.Body(
		app.Span().
			Class("mdc-deprecated-list-item__ripple"),
		l.Igraphic,
		app.Span().
			Class("mdc-deprecated-list-item__text").
			Text(l.Itext),
		l.Imeta,
	)
}

func ListItem() IListItem {
	return &listItem{}
}

type listItemTwoLine struct {
	listItem

	IsecondaryText string
}

type IListItemTwoLine interface {
	app.UI

	ID(string) IListItemTwoLine
	Text(string) IListItemTwoLine
	Graphic(app.UI) IListItemTwoLine

	SecondaryText(string) IListItemTwoLine
}

func (l *listItemTwoLine) SecondaryText(t string) IListItemTwoLine {
	l.IsecondaryText = t
	return l
}

func (l *listItemTwoLine) ID(id string) IListItemTwoLine {
	l.id = id
	return l
}

func (l *listItemTwoLine) Text(t string) IListItemTwoLine {
	l.Itext = t
	return l
}

func (l *listItemTwoLine) Graphic(g app.UI) IListItemTwoLine {
	l.Igraphic = g
	return l
}

func (l *listItemTwoLine) Meta(m app.UI) IListItemTwoLine {
	l.Imeta = m
	return l
}


func (l *listItemTwoLine) Render() app.UI {
	item := app.Li().
		Class("mdc-deprecated-list-item").
		ID(l.id)

	primary := app.Span().
			Class("mdc-deprecated-list-item__primary-text").
			Text(l.Itext)

	secondary := app.Span().
			Class("mdc-deprecated-list-item__secondary-text").
			Text(l.IsecondaryText)

	return item.Body(
		app.Span().
			Class("mdc-deprecated-list-item__ripple"),
		l.Igraphic,
		app.Span().
			Class("mdc-deprecated-list-item__text").
			Body(
				primary,
				secondary,
			),
		l.Imeta,
	)
}

func ListItemTwoLine() IListItemTwoLine {
	return &listItemTwoLine{}
}
