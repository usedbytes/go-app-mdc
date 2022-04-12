package mdc

// https://material.io/components/buttons/web

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var idCount int

type button struct {
	app.Compo

	id string

	label string
	leadingIcon string
	trailingIcon string

	raised bool
	outlined bool // Ignored if raised

	onClick app.EventHandler
}

type IButton interface {
	app.UI

	Label(string) IButton
	Raised(bool) IButton
	Outlined(bool) IButton
	LeadingIcon(string) IButton
	TrailingIcon(string) IButton
	OnClick(app.EventHandler) IButton
}

func (b *button) Label(label string) IButton {
	b.label = label
	return b
}

func (b *button) Raised(raised bool) IButton {
	b.raised = raised
	return b
}

func (b *button) Outlined(outlined bool) IButton {
	b.outlined = outlined
	return b
}

func (b *button) LeadingIcon(icon string) IButton {
	b.leadingIcon = icon
	return b
}

func (b *button) TrailingIcon(icon string) IButton {
	b.trailingIcon = icon
	return b
}

func (b *button) OnClick(handler app.EventHandler) IButton {
	b.onClick = handler
	return b
}

func (b *button) OnMount(ctx app.Context) {
	b.id = fmt.Sprintf("button-%d", idCount)
	idCount++
}

func (b *button) Render() app.UI {
	button := app.Button().
		Class("mdc-button").
		ID(b.id)

	if (b.raised) {
		button = button.Class("mdc-button--raised")
	} else if (b.outlined) {
		button = button.Class("mdc-button--outlined")
	}

	var leadingIcon app.UI
	if b.leadingIcon != "" {
		button = button.Class("mdc-button--icon-trailing")
		leadingIcon = app.I().
			Class("material-icons", "mdc-button__icon").
			Aria("hidden", "true").
			Text(b.leadingIcon)
	}

	var trailingIcon app.UI
	if b.trailingIcon != "" {
		button = button.Class("mdc-button--icon-leading")
		trailingIcon = app.I().
			Class("material-icons", "mdc-button__icon").
			Aria("hidden", "true").
			Text(b.trailingIcon)
	}

	if b.onClick != nil {
		button = button.OnClick(b.onClick)
	}

	return button.Body(
		app.Span().Class("mdc-button__ripple"),
		app.If(leadingIcon != nil, leadingIcon),
		app.Span().Class("mdc-button__label").Text(b.label),
		app.If(trailingIcon != nil, trailingIcon),
	)
}

func Button() IButton {
	return &button{}
}
