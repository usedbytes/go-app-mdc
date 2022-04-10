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
	Icon string

	Raised bool
	Outlined bool // Ignored if Raised
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

	ripple := app.Span().
		Class("mdc-button__ripple")

	label := app.Span().
		Class("mdc-button__label").
		Text(b.Label)

	icon := app.I().
		Class("material-icons", "mdc-button__icon").
		Aria("hidden", "true").
		Text(b.Icon)

	return button.Body(
		ripple,
		app.If(b.Icon != "", icon),
		label,
	)
}
