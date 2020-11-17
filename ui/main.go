package main

import (
	"fyne.io/fyne/app"
)

func main() {
	app := app.New()
	app.SetIcon(Icon)

	Show(app)
	app.Run()
}
