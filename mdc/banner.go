package mdc

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// https://material.io/components/banners/web#banners

type banner struct {
	app.Compo

	id string
	mdcComponent app.Value

	Iclass []string
	Iopen bool
	Ifixed bool
	Itext string
	Igraphic app.UI
	Iactions []app.UI
}

const (
	BannerGraphicClass = "mdc-banner__icon"
)

type IBanner interface {
	app.UI

	ID(string) IBanner
	Class(...string) IBanner
	Fixed(bool) IBanner
	Text(string) IBanner
	Graphic(app.UI) IBanner
	Actions(...app.UI) IBanner

	Open() IBanner
	Close() IBanner
}

func (b *banner) ID(id string) IBanner {
	b.id = id
	return b
}

func (b *banner) Class(class ...string) IBanner {
	b.Iclass = class
	return b
}

func (b *banner) Fixed(fixed bool) IBanner {
	b.Ifixed = fixed
	return b
}

func (b *banner) Text(text string) IBanner {
	b.Itext = text
	return b
}

func (b *banner) Graphic(graphic app.UI) IBanner {
	b.Igraphic = graphic
	return b
}

func (b *banner) Actions(actions ...app.UI) IBanner {
	b.Iactions = actions
	return b
}

func (b *banner) OnMount(ctx app.Context) {
	if b.id == "" {
		b.id = fmt.Sprintf("banner-%d", allocID())
	}

	b.mdcComponent = app.Window().
		Get("mdc").
		Get("banner").
		Get("MDCBanner").
		Call("attachTo", ctx.JSSrc())

	if b.Iopen {
		b.mdcComponent.Call("open")
	}

	app.Log("mounted", b.id, b.mdcComponent)
}

func (b *banner) Open() IBanner {
	b.Iopen = true
	if b.mdcComponent != nil {
		b.mdcComponent.Call("open")
	}

	return b
}

func (b *banner) Close() IBanner {
	b.Iopen = false
	if b.mdcComponent != nil {
		b.mdcComponent.Call("close")
	}

	return b
}

func (b *banner) OnDismount(ctx app.Context) {
	if b.mdcComponent != nil {
		b.mdcComponent.Call("destroy")
	}
}

func (b *banner) Render() app.UI {
	banner := app.Div().
		Class("mdc-banner").
		Class(b.Iclass...).
		Role("banner")

	var graphic app.UI
	if b.Igraphic != nil {
		graphic = app.Div().
			Class("mdc-banner__graphic").
			Role("img").
			Body(b.Igraphic)
	}

	text := app.Div().
		Class("mdc-banner__graphic-text-wrapper").
		Body(
			graphic,
			app.Div().
				Class("mdc-banner__text").
				Text(b.Itext),
		)

	actions := app.Div().
		Class("mdc-banner__actions").
		Body(
			b.Iactions...
		)

	content := app.Div().
		Class("mdc-banner__content").
		Role("alert-dialog").
		Aria("live", "assistive").
		Body(
			text,
			actions,
		)

	if b.Ifixed {
		content = app.Div().
			Class("mdc-banner__fixed").
			Body(content)
	}

	return banner.Body(
		content,
	)
}

func Banner() IBanner {
	return &banner{}
}
