package main

import (
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/app"
)

const (
	title   = "Runtime Script"
	srcPath = "src/program.runtime"
)

func main() {
	app := app.New()

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(dir, srcPath)
	// color eg []uint8{0x00, 0xbb, 0xbb, 0xff}
	ShowConsole(app, title, path, nil)
	app.Run()
}
