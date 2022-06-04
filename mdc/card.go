package mdc

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

// https://material.io/components/app-bars-top/web#using-the-top-app-bar
// https://material.io/components/cards/web#card

const (
	CardPrimaryAction = "mdc-card__primary-action"
	CardActionsClass = "mdc-card__actions"
	CardActionsFullBleedClass = "mdc-card__actions--full-bleed"
	CardActionClass = "mdc-card__action"
	CardActionButtonsClass = "mdc-card__action-buttons"
	CardActionButtonClass = "mdc-card__action--button"
	CardActionIconsClass = "mdc-card__action-icons"
	CardActionIconClass = "mdc-card__action--icon"
	CardMediaClass = "mdc-card__media"
	CardMediaSquareClass = "mdc-card__media--square"
	CardMedia16by9Class = "mdc-card__media--16-9"
	CardMediaContentClass = "mdc-card__media"
)

type card struct {
	app.Compo

	id string

	IContents []app.UI
	IButtons []app.UI
	IIcons []app.UI
	IOutlined bool
	IActionFullBleed bool
	IClasses []string
	IonClick app.EventHandler
}

type ICard interface {
	app.UI

	ID(string) ICard
	Class(...string) ICard
	Outlined(bool) ICard
	ActionFullBleed(bool) ICard
	Contents(...app.UI) ICard
	Buttons(...app.UI) ICard
	Icons(...app.UI) ICard
	OnClick(app.EventHandler) ICard
}

func (c *card) ID(id string) ICard {
	c.id = id
	return c
}

func (c *card) Class(classes ...string) ICard {
	c.IClasses = append(c.IClasses, classes...)
	return c
}

func (c *card) Outlined(outlined bool) ICard {
	c.IOutlined = outlined
	return c
}

func (c *card) ActionFullBleed(bleed bool) ICard {
	c.IActionFullBleed = bleed
	return c
}

func (c *card) Contents(v ...app.UI) ICard {
	c.IContents = v
	return c
}

func (c *card) Buttons(v ...app.UI) ICard {
	c.IButtons = v
	return c
}

func (c *card) Icons(v ...app.UI) ICard {
	c.IIcons = v
	return c
}

func (c *card) OnClick(handler app.EventHandler) ICard {
	c.IonClick = handler
	return c
}


func (c *card) OnMount(ctx app.Context) {
	if c.id == "" {
		c.id = fmt.Sprintf("card-%d", allocID())
	}

	app.Log("mounted", c.id)
}

func (c *card) Render() app.UI {
	div := app.Div().
		ID(c.id).
		Class("mdc-card").
		Class(c.IClasses...)

	if c.IOutlined {
		div = div.Class("mdc-card--outlined")
	}

	if c.IonClick != nil {
		div = div.OnClick(c.IonClick)
	}

	// TODO: The styles are needed to make media/images
	// work. Adding mdc-card__primary-action works too, but
	// that adds additional meaning.
	content := app.Div().
		Style("display", "flex").
		Style("flex-direction", "column").
		Body(c.IContents...)

	// Handle wrappers for buttons/icons
	// This isn't very neat...
	var actions app.HTMLDiv
	var buttons app.HTMLDiv
	var icons app.HTMLDiv

	if c.IButtons != nil {
		if c.IIcons == nil {
			// Buttons only
			buttons = app.Div().
				Class(CardActionsClass)
		} else {
			buttons = app.Div().
				Class(CardActionButtonsClass)
		}

		buttons = buttons.Body(c.IButtons...)
	}

	if c.IIcons != nil {
		if c.IButtons == nil {
			// Icons only
			icons = app.Div().
				Class(CardActionsClass)
		} else {
			icons = app.Div().
				Class(CardActionIconsClass)
		}

		icons = icons.Body(c.IIcons...)
	}

	if buttons != nil && icons != nil {
		actions = app.Div().
			Class(CardActionsClass).
			Body(buttons, icons)
	} else if buttons != nil {
		actions = buttons
	} else if icons != nil {
		actions = icons
	}

	if actions != nil && c.IActionFullBleed {
		actions = actions.Class(CardActionsFullBleedClass)
	}

	return div.Body(
		content,
		actions,
	)
}

func Card() ICard {
	return &card{}
}

type cardMedia struct {
	app.Compo

	id string

	IContents []app.UI
	IClasses []string
	IImage string
}

type ICardMedia interface {
	app.UI

	ID(string) ICardMedia
	Class(...string) ICardMedia
	Image(string) ICardMedia
	Contents(...app.UI) ICardMedia
}

func (c *cardMedia) ID(id string) ICardMedia {
	c.id = id
	return c
}

func (c *cardMedia) Class(classes ...string) ICardMedia {
	c.IClasses = append(c.IClasses, classes...)
	return c
}

func (c *cardMedia) Image(url string) ICardMedia {
	c.IImage = url
	return c
}

func (c *cardMedia) Contents(v ...app.UI) ICardMedia {
	c.IContents = v
	return c
}

func (c *cardMedia) OnMount(ctx app.Context) {
	if c.id == "" {
		c.id = fmt.Sprintf("card-media-%d", allocID())
	}

	app.Log("mounted", c.id)
}

func (c *cardMedia) Render() app.UI {
	div := app.Div().
		ID(c.id).
		Class(CardMediaClass).
		Class(c.IClasses...).
		Style("background-image", "url(" + c.IImage + ")")

	return div.Body(c.IContents...)
}

func CardMedia() ICardMedia {
	return &cardMedia{}
}

