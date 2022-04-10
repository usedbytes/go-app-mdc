package mdc

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func MaterialIcon(icon string, class ...string) app.HTMLI {
	i := app.I().
		Class("material-icons").
		Aria("hidden", "true")

	for _, c := range class {
		i = i.Class(c)
	}

	return i.Text(icon)
}
