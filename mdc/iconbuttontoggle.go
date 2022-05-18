package mdc

// https://material.io/develop/web/components/buttons/icon-buttons

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type iconButtonToggle struct {
	app.Compo

	id string
	mdcComponent app.Value

	IiconOff string
	IiconOn string
	Ion bool

	IonClick app.EventHandler

	Iclass []string
}

type IIconButtonToggle interface {
	app.UI

	ID(string) IIconButtonToggle
	IconOn(string) IIconButtonToggle
	IconOff(string) IIconButtonToggle
	On(bool) IIconButtonToggle
	OnClick(app.EventHandler) IIconButtonToggle
	Class(...string) IIconButtonToggle
}

func (b *iconButtonToggle) ID(id string) IIconButtonToggle {
	b.id = id
	return b
}

func (b *iconButtonToggle) IconOff(icon string) IIconButtonToggle {
	b.IiconOff = icon
	return b
}

func (b *iconButtonToggle) IconOn(icon string) IIconButtonToggle {
	b.IiconOn = icon
	return b
}

func (b *iconButtonToggle) On(on bool) IIconButtonToggle {
	b.Ion = on
	return b
}

func (b *iconButtonToggle) OnClick(handler app.EventHandler) IIconButtonToggle {
	b.IonClick = handler
	return b
}

func (b *iconButtonToggle) Class(c ...string) IIconButtonToggle {
	b.Iclass = append(b.Iclass, c...)
	return b
}

func (b *iconButtonToggle) OnMount(ctx app.Context) {
	if b.id == "" {
		b.id = fmt.Sprintf("iconbutton-%d", allocID())
	}

	b.mdcComponent = app.Window().
		Get("mdc").
		Get("iconButton").
		Get("MDCIconButtonToggle").
		Call("attachTo", ctx.JSSrc())

	app.Log("mounted", b.id, b.mdcComponent)
}

func (b *iconButtonToggle) OnDismount(ctx app.Context) {
	b.mdcComponent.Call("destroy")
}

func (b *iconButtonToggle) Render() app.UI {
	iconButtonToggle := app.Button().
		Class("mdc-icon-button").
		Class(b.Iclass...).
		Aria("pressed", b.Ion).
		ID(b.id)

	if b.Ion {
		iconButtonToggle = iconButtonToggle.Class("mdc-icon-button--on")
	}

	if b.IonClick != nil {
		iconButtonToggle = iconButtonToggle.OnClick(b.IonClick)
	}

	iconOn := app.I().
		Class("material-icons").
		Class("mdc-icon-button__icon").
		Class("mdc-icon-button__icon--on").
		Text(b.IiconOn)

	iconOff := app.I().
		Class("material-icons").
		Class("mdc-icon-button__icon").
		Text(b.IiconOff)

	return iconButtonToggle.Body(
		app.Div().Class("mdc-icon-button__ripple"),
		iconOn,
		iconOff,
	)
}

func IconButtonToggle() IIconButtonToggle {
	return &iconButtonToggle{}
}
