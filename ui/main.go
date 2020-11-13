package main

import (
	"fyne.io/fyne/app"
)

func main() {
	//a := app.New()
	//w := a.NewWindow("Runtime Script")

	app := app.New()
	app.SetIcon(Icon)

	Show(app)
	app.Run()

	//w.SetFixedSize(true)
	//w.Resize(fyne.Size{Width: 800, Height: 600})

	//hello := widget.NewLabel("Hello Fyne!")
	// w.SetContent(widget.NewVBox(
	// 	hello,
	// 	widget.NewButton("Hi!", func() {
	// 		hello.SetText("Welcome :)")
	// 	}),
	// 	widget.NewButton("abc", func() {
	// 		hello.SetText("Hello World")
	// 	}),
	// ))

}
