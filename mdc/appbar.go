package mdc

// https://material.io/components/app-bars-top/web#using-the-top-app-bar

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	AppBarMainClass = "mdc-top-app-bar--fixed-adjust"
	AppBarNavigationClass = "mdc-top-app-bar__navigation-item"
	AppBarActionClass = "mdc-top-app-bar__action-item"
)

type appBarTop struct {
	app.Compo

	id string
	mdcComponent app.Value

	InavButton app.UI
	Ititle string
	IactionButtons []app.UI

	IonClick app.EventHandler
}

type IAppBar interface {
	app.UI

	ID(string) IAppBar
	NavIcon(app.UI) IAppBar
	Title(string) IAppBar
	Actions(...app.UI) IAppBar
}

func (a *appBarTop) ID(id string) IAppBar {
	a.id = id
	return a
}

func (a *appBarTop) NavIcon(icon app.UI) IAppBar {
	a.InavButton = icon
	return a
}

func (a *appBarTop) Title(t string) IAppBar {
	a.Ititle = t
	return a
}

func (a *appBarTop) Actions(v ...app.UI) IAppBar {
	a.IactionButtons = app.FilterUIElems(v...)
	return a
}

func (a *appBarTop) OnMount(ctx app.Context) {
	if a.id == "" {
		a.id = fmt.Sprintf("appbar-%d", allocID())
	}

	a.mdcComponent = app.Window().
		Get("mdc").
		Get("topAppBar").
		Get("MDCTopAppBar").
		Call("attachTo", ctx.JSSrc())

	app.Log("mounted", a.id, a.mdcComponent)
}

func (a *appBarTop) OnDismount(ctx app.Context) {
	a.mdcComponent.Call("destroy")
}

func (a *appBarTop) Render() app.UI {
	appBarTop := app.Header().
		Class("mdc-top-app-bar").
		ID(a.id)

	sectionStart := app.Section().
		Class("mdc-top-app-bar__section").
		Class("mdc-top-app-bar__section--align-start").
		Body(
			a.InavButton,
			app.Span().
				Class("mdc-top-app-bar__title").
				Text(a.Ititle),
		)

	sectionEnd := app.Section().
		Class("mdc-top-app-bar__section").
		Class("mdc-top-app-bar__section--align-end").
		Role("toolbar").
		Body(
			a.IactionButtons...,
		)

	row := app.Div().
		Class("mdc-top-app-bar__row").
		Body(
			sectionStart,
			sectionEnd,
		)

	return appBarTop.Body(
		row,
	)
}

func AppBarTop() IAppBar {
	return &appBarTop{}
}
