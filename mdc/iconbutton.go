package mdc

// https://material.io/develop/web/components/buttons/icon-buttons

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type iconButton struct {
	app.Compo

	id string
	mdcComponent app.Value

	Iicon string

	IonClick app.EventHandler

	Iclass []string
}

type IIconButton interface {
	app.UI

	ID(string) IIconButton
	Icon(string) IIconButton
	OnClick(app.EventHandler) IIconButton

	Class(...string) IIconButton
}

func (b *iconButton) ID(id string) IIconButton {
	b.id = id
	return b
}

func (b *iconButton) Icon(icon string) IIconButton {
	b.Iicon = icon
	return b
}

func (b *iconButton) OnClick(handler app.EventHandler) IIconButton {
	b.IonClick = handler
	return b
}

func (b *iconButton) Class(c ...string) IIconButton {
	b.Iclass = c
	return b
}

func (b *iconButton) OnMount(ctx app.Context) {
	if b.id == "" {
		b.id = fmt.Sprintf("iconbutton-%d", allocID())
	}

	// FIXME: For some reason this causes a spurious additional
	// circle
	// b.mdcComponent = app.Window().
	// 	Get("mdc").
	// 	Get("ripple").
	// 	Get("MDCRipple").
	// 	Call("attachTo", ctx.JSSrc())

	app.Log("mounted", b.id, b.mdcComponent)
}

func (b *iconButton) OnDismount(ctx app.Context) {
	// b.mdcComponent.Call("destroy")
}

func (b *iconButton) Render() app.UI {
	iconButton := app.Button().
		Class("mdc-icon-button", "material-icons").
		Class(b.Iclass...).
		ID(b.id)


	if b.IonClick != nil {
		iconButton = iconButton.OnClick(b.IonClick)
	}

	return iconButton.Body(
		app.Div().Class("mdc-icon-button__ripple"),
		app.Text(b.Iicon),
	)
}

func IconButton() IIconButton {
	return &iconButton{}
}
