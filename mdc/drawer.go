package mdc

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// https://material.io/components/navigation-drawer/web
// https://github.com/material-components/material-components-web/tree/master/packages/mdc-drawer

type drawer struct {
	app.Compo

	id string
	mdcComponent app.Value

	Iclass []string
	Ititle string
	Isubtitle string
	Icontent []app.UI
}

type IDrawer interface {
	app.UI

	ID(string) IDrawer
	Class(...string) IDrawer
	Title(string) IDrawer
	Subtitle(string) IDrawer
	Content(...app.UI) IDrawer
}

func (d *drawer) ID(id string) IDrawer {
	d.id = id
	return d
}

func (d *drawer) Class(c ...string) IDrawer {
	d.Iclass = c
	return d
}

func (d *drawer) Title(t string) IDrawer {
	d.Ititle = t
	return d
}

func (d *drawer) Subtitle(t string) IDrawer {
	d.Isubtitle = t
	return d
}

func (d *drawer) Content(i ...app.UI) IDrawer {
	d.Icontent = i
	return d
}

func (d *drawer) OnMount(ctx app.Context) {
	if d.id == "" {
		d.id = fmt.Sprintf("drawer-%d", allocID())
	}

	/* Only for dismissable drawers
	d.mdcComponent = app.Window().
		Get("mdc").
		Get("drawer").
		Get("MDCDrawer").
		Call("attachTo", ctx.JSSrc())
	*/

	app.Log("mounted", d.id, d.mdcComponent)
}

func (d *drawer) OnDismount(ctx app.Context) {
	//d.mdcComponent.Call("destroy")
}

func (d *drawer) Render() app.UI {
	aside := app.Aside().
		Class("mdc-drawer").
		Class(d.Iclass...).
		ID(d.id)

	var header app.UI
	if d.Ititle != "" || d.Isubtitle != "" {
		header = app.Div().
			Class("mdc-drawer__header").
			Body(
				app.If(d.Ititle != "", app.H3().
					Class("mdc-drawer__title").
					Text(d.Ititle)),
				app.If(d.Isubtitle != "", app.H6().
					Class("mdc-drawer__subtitle").
					Text(d.Isubtitle)),
			)
	}

	content := app.Div().
		Class("mdc-drawer__content").
		Body(
			d.Icontent...
		)

	return aside.Body(
		header,
		content,
	)
}

func Drawer() IDrawer {
	return &drawer{}
}
