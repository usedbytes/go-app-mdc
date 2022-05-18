package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"

	"github.com/usedbytes/go-app-mdc/mdc"
)

type cardExample struct {
	app.Compo

	cards []app.UI
}

func CardExample() *cardExample {
	ex := new(cardExample)

	ex.cards = []app.UI{
		mdc.Card().
			Contents(
				app.Div().
					Style("padding", "16px").
					Body(
						app.H1().
							Class("mdc-typography--headline5").
							Text("Card 1"),
						app.H2().
							Class("mdc-typography--headline6").
							Text("Needs some more helpers"),
						app.Div().
							Class("mdc-typography--body1").
							Text("This is the body, I need to write some long text to see how it looks. This element is getting really heavily indented, it's quite messy. Still, it would have been much easier to just copy-paste lorem ipsum or something."),
					),
			),
		mdc.Card().
			Class("img-card").
			Outlined(true).
			Contents(
				app.Div().
					Style("padding", "16px").
					Body(
						app.H1().
							Class("mdc-typography--headline5").
							Text("Random Image"),
					),
				mdc.CardMedia().
					Class(mdc.CardMedia16by9Class).
					Image("https://picsum.photos/480/270?random=1"),
				app.Div().
					Style("padding", "16px").
					Body(
						app.Div().
							Class("mdc-typography--body1").
							Text("This is a random image from https://picsum.photos. You can put text on top of the image, but it's hard to pick the right color when the image is random!"),
					),
			),
		mdc.Card().
			Class("img-card").
			Outlined(true).
			Contents(
				mdc.CardMedia().
					Class(mdc.CardMedia16by9Class).
					Image("https://picsum.photos/480/270?random=2"),
			).
			Buttons(
				mdc.Button().
					Class(mdc.CardActionButtonClass).
					Class(mdc.CardActionClass).
					Raised(true).
					Label("Card Button"),
			),
		mdc.Card().
			Class("img-card").
			Outlined(true).
			Contents(
				mdc.CardMedia().
					Class(mdc.CardMedia16by9Class).
					Image("https://picsum.photos/480/270?random=3"),
			).
			Icons(
				mdc.IconButton().
					Class(mdc.CardActionIconClass).
					Class(mdc.CardActionClass).
					Icon("info"),
				mdc.IconButton().
					Class(mdc.CardActionIconClass).
					Class(mdc.CardActionClass).
					Icon("help"),
				mdc.IconButton().
					Class(mdc.CardActionIconClass).
					Class(mdc.CardActionClass).
					Icon("warning"),
				mdc.IconButton().
					Class(mdc.CardActionIconClass).
					Class(mdc.CardActionClass).
					Icon("error"),
			),
		mdc.Card().
			Class("img-card").
			Outlined(true).
			Contents(
				mdc.CardMedia().
					Class(mdc.CardMedia16by9Class).
					Image("https://picsum.photos/480/270?random=4"),
			).
			Buttons(
				mdc.Button().
					Class(mdc.CardActionButtonClass).
					Class(mdc.CardActionClass).
					Raised(true).
					Label("Card Button"),
			).
			Icons(
				mdc.IconButton().
					Class(mdc.CardActionIconClass).
					Class(mdc.CardActionClass).
					Icon("favorite"),
			),
		mdc.Card().
			Class("img-card").
			Outlined(true).
			Contents(
				mdc.CardMedia().
					Class(mdc.CardMedia16by9Class).
					Image("https://picsum.photos/480/270?random=5"),
			).
			ActionFullBleed(true).
			Buttons(
				mdc.Button().
					Class(mdc.CardActionButtonClass).
					Class(mdc.CardActionClass).
					Label("Full Bleed Button"),
			),
	}

	return ex
}

func (e *cardExample) Render() app.UI {
	return app.Main().
		Class(mdc.AppBarMainClass).
		Class("main-content").
		Style("display", "flex").
		Style("flex-wrap", "wrap").
		Body(
			app.Range(e.cards).Slice(func(i int) app.UI {
				return app.Div().
					Style("padding", "5px").
					Body(
						e.cards[i],
					)
			}),
		)
}
