package mdc

// https://material.io/components/buttons/web

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var idCount int

type Button struct {
	app.Compo

	id string

	Label string
	Icon interface{}

	Raised bool
	Outlined bool // Ignored if Raised
	TrailingIcon bool

	OnClick app.EventHandler
}

func (b *Button) OnMount(ctx app.Context) {
	b.id = fmt.Sprintf("button-%d", idCount)
	idCount++
}

func (b *Button) Render() app.UI {
	button := app.Button().
		Class("mdc-button").
		ID(b.id)

	if (b.Raised) {
		button = button.Class("mdc-button--raised")
	} else if (b.Outlined) {
		button = button.Class("mdc-button--outlined")
	}

	var icon app.UI
	if b.Icon != nil {
		if (b.TrailingIcon) {
			button = button.Class("mdc-button--icon-trailing")
		} else {
			button = button.Class("mdc-button--icon-leading")
		}

		switch t := b.Icon.(type) {
		case string:
			icon = MaterialIcon(t, "mdc-button__icon")
		default:
			app.Log("Unsupported type for button icon:", b.id)
		}
	}

	if b.OnClick != nil {
		button = button.OnClick(b.OnClick)
	}

	return button.Body(
		app.Span().Class("mdc-button__ripple"),
		app.If(icon != nil && !b.TrailingIcon, icon),
		app.Span().Class("mdc-button__label").Text(b.Label),
		app.If(icon != nil && b.TrailingIcon, icon),
	)
}
