package mdc

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// https://material.io/components/buttons-floating-action-button/web

const (
	FABIconClass = "mdc-fab__icon"
)

type fab struct {
	app.Compo

	id string
	mdcComponent app.Value

	Iicon app.UI
	Ilabel string
	IonClick app.EventHandler
}

type IFAB interface {
	app.UI

	ID(string) IFAB
	Label(string) IFAB
	Icon(app.UI) IFAB
	OnClick(app.EventHandler) IFAB
}

func (f *fab) ID(id string) IFAB {
	f.id = id
	return f
}

func (f *fab) Label(label string) IFAB {
	f.Ilabel = label
	return f
}

func (f *fab) Icon(icon app.UI) IFAB {
	f.Iicon = icon
	return f
}

func (f *fab) OnClick(handler app.EventHandler) IFAB {
	f.IonClick = handler
	return f
}

func (f *fab) OnMount(ctx app.Context) {
	if f.id == "" {
		f.id = fmt.Sprintf("fab-%d", allocID())
	}

	f.mdcComponent = app.Window().
		Get("mdc").
		Get("ripple").
		Get("MDCRipple").
		Call("attachTo", ctx.JSSrc())

	app.Log("mounted", f.id, f.mdcComponent)
}

func (f *fab) OnDismount(ctx app.Context) {
	if f.mdcComponent != nil {
		f.mdcComponent.Call("destroy")
	}
}

func (f *fab) Render() app.UI {
	fab := app.Button().
		Class("mdc-fab").
		ID(f.id)

	if f.IonClick != nil {
		fab = fab.OnClick(f.IonClick)
	}

	if f.Ilabel != "" {
		fab = fab.Class("mdc-fab--extended")
	}

	return fab.Body(
		app.Span().Class("mdc-fab__ripple"),
		app.If(f.Iicon != nil, f.Iicon),
		app.If(
			f.Ilabel != "",
			app.Span().Class("mdc-fab__label").Text(f.Ilabel),
		),
	)
}

func FAB() IFAB {
	return &fab{}
}
