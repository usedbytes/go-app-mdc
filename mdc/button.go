package mdc

// https://material.io/components/buttons/web

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type button struct {
	app.Compo

	id string
	mdcComponent app.Value

	Ilabel string
	IleadingIcon string
	ItrailingIcon string

	Iraised bool
	Ioutlined bool // Ignored if raised

	IonClick app.EventHandler
}

type IButton interface {
	app.UI

	ID(string) IButton
	Label(string) IButton
	Raised(bool) IButton
	Outlined(bool) IButton
	LeadingIcon(string) IButton
	TrailingIcon(string) IButton
	OnClick(app.EventHandler) IButton
}

func (b *button) ID(id string) IButton {
	b.id = id
	return b
}

func (b *button) Label(label string) IButton {
	b.Ilabel = label
	return b
}

func (b *button) Raised(raised bool) IButton {
	b.Iraised = raised
	return b
}

func (b *button) Outlined(outlined bool) IButton {
	b.Ioutlined = outlined
	return b
}

func (b *button) LeadingIcon(icon string) IButton {
	b.IleadingIcon = icon
	return b
}

func (b *button) TrailingIcon(icon string) IButton {
	b.ItrailingIcon = icon
	return b
}

func (b *button) OnClick(handler app.EventHandler) IButton {
	b.IonClick = handler
	return b
}

func (b *button) OnMount(ctx app.Context) {
	if b.id == "" {
		b.id = fmt.Sprintf("button-%d", allocID())
	}

	b.mdcComponent = app.Window().
		Get("mdc").
		Get("ripple").
		Get("MDCRipple").
		Call("attachTo", ctx.JSSrc())

	app.Log("mounted", b.id, b.mdcComponent)
}

func (b *button) OnDismount(ctx app.Context) {
	b.mdcComponent.Call("destroy")
}

func (b *button) Render() app.UI {
	button := app.Button().
		Class("mdc-button").
		ID(b.id)

	if (b.Iraised) {
		button = button.Class("mdc-button--raised")
	} else if (b.Ioutlined) {
		button = button.Class("mdc-button--outlined")
	}

	var leadingIcon app.UI
	if b.IleadingIcon != "" {
		button = button.Class("mdc-button--icon-trailing")
		leadingIcon = app.I().
			Class("material-icons", "mdc-button__icon").
			Aria("hidden", "true").
			Text(b.IleadingIcon)
	}

	var trailingIcon app.UI
	if b.ItrailingIcon != "" {
		button = button.Class("mdc-button--icon-leading")
		trailingIcon = app.I().
			Class("material-icons", "mdc-button__icon").
			Aria("hidden", "true").
			Text(b.ItrailingIcon)
	}

	if b.IonClick != nil {
		button = button.OnClick(b.IonClick)
	}

	return button.Body(
		app.Span().Class("mdc-button__ripple"),
		app.If(leadingIcon != nil, leadingIcon),
		app.Span().Class("mdc-button__label").Text(b.Ilabel),
		app.If(trailingIcon != nil, trailingIcon),
	)
}

func Button() IButton {
	return &button{}
}
