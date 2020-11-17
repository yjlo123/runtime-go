// Package beebui emulates a BBC Micro Computer
package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"strings"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/driver/desktop"
	runtime "github.com/yjlo123/runtime-go/src"
)

const (
	screenInsetX = 92
	screenInsetY = 72

	screenLines = 24
	screenCols  = 38

	cursorChar = "_"
	maxHistory = 10
)

var (
	screenSize = fyne.Size{Width: 800, Height: 600}
	lineDelay  = time.Second * 0 //time.Second / 10

	// Icon ..
	Icon = icon
)

type console struct {
	content       []fyne.CanvasObject
	input         chan string // user input will be sent to this channel after pressing Enter
	inputLen      int         // user input length
	headLen       int         // number of chars before the user input
	history       []string
	historyCursor int

	overlay *canvas.Image
	current int // current line number
}

func (con *console) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return screenSize
}

func (con *console) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	con.overlay.Resize(size)

	y := screenInsetY
	for i := 0; i < screenLines; i++ {
		con.content[i].Move(fyne.NewPos(screenInsetX, y))
		con.content[i].Resize(fyne.NewSize(size.Width-screenInsetX*2, 18))
		y += 19
	}
}

func (con *console) loadUI() fyne.CanvasObject {
	con.content = make([]fyne.CanvasObject, screenLines)

	for i := 0; i < screenLines; i++ {
		con.content[i] = canvas.NewText("", color.RGBA{0xbb, 0xbb, 0xbb, 0xff})
		con.content[i].(*canvas.Text).TextSize = 15
		con.content[i].(*canvas.Text).TextStyle.Monospace = true
	}

	con.input = make(chan string)

	con.overlay = canvas.NewImageFromResource(monitor)
	return fyne.NewContainerWithLayout(con, append(con.content, con.overlay)...)
}

func (con *console) appendLine(line string) {
	con.append(line)
	con.newLine()
	con.inputLen = 0
	con.headLen = 0
}

func (con *console) newLine() {
	time.Sleep(lineDelay)
	text := con.content[con.current].(*canvas.Text)

	if len(text.Text) > 0 && text.Text[con.headLen+con.inputLen:] == cursorChar {
		text.Text = text.Text[:len(text.Text)-1]
		canvas.Refresh(text)
	}

	if con.current == screenLines-1 {
		con.scroll()
	}
	con.current++
}

func (con *console) append(line string) {
	text := con.content[con.current].(*canvas.Text)
	if len(text.Text) > 0 && text.Text[con.headLen+con.inputLen-len(line):] == cursorChar {
		text.Text = text.Text[:len(text.Text)-1]
	}
	text.Text = text.Text + line
	canvas.Refresh(text)
}

func (con *console) blink() {
	for {
		time.Sleep(time.Second / 2)
		line := con.content[con.current].(*canvas.Text)

		if con.headLen+con.inputLen == len(line.Text) {
			line.Text = line.Text + cursorChar
		} else {
			line.Text = line.Text[:len(line.Text)-1]
		}
		canvas.Refresh(con.content[con.current])
	}
}

func (con *console) scroll() {
	for i := 0; i < len(con.content)-1; i++ {
		text1 := con.content[i].(*canvas.Text)
		text2 := con.content[i+1].(*canvas.Text)
		text1.Text = text2.Text

		canvas.Refresh(text1)
	}

	text := con.content[len(con.content)-1].(*canvas.Text)
	text.Text = ""
	canvas.Refresh(text)

	con.current--
}

func (con *console) onRune(r rune) {
	if r > 128 {
		return
	}
	if con.headLen+con.inputLen < screenCols-1 {
		con.inputLen++
		con.append(string(r))
	}
}

func (con *console) onKey(ev *fyne.KeyEvent) {
	line := con.content[con.current].(*canvas.Text)
	text := line.Text
	switch ev.Name {
	case fyne.KeyReturn:
		userInput := text[con.headLen : con.headLen+con.inputLen]
		con.appendLine("")
		fmt.Println(userInput)
		con.input <- userInput
		if len(userInput) > 0 {
			// push history
			con.history = append(con.history, userInput)
			if len(con.history) > maxHistory {
				con.history = con.history[:maxHistory]
			}
		}
		con.historyCursor = 0
	case fyne.KeyBackspace:
		if text[con.headLen+con.inputLen:] == cursorChar {
			text = text[:len(text)-1]
		}
		if len(text) > 0 && con.inputLen > 0 {
			line.Text = text[:len(text)-1]
			con.inputLen--
			canvas.Refresh(line)
		}
	case fyne.KeyEscape:
		if text[con.headLen+con.inputLen:] == cursorChar {
			text = text[:len(text)-1]
		}
		line.Text = text[:len(text)-con.inputLen]
		con.inputLen = 0
		canvas.Refresh(line)
	case fyne.KeyUp:
		if len(con.history) > 0 {
			historyIdx := len(con.history) - 1 - con.historyCursor
			if historyIdx < 0 {
				return
			}
			showHistory(con, historyIdx)
			con.historyCursor = con.historyCursor + 1
		}
		fmt.Println(con.historyCursor)
	case fyne.KeyDown:
		if con.historyCursor > 0 {
			historyIdx := len(con.history) + 1 - con.historyCursor
			if historyIdx == len(con.history) {
				return
			}
			showHistory(con, historyIdx)
			con.historyCursor = con.historyCursor - 1
		}
		fmt.Println(con.historyCursor)
	}
}

func showHistory(con *console, historyIdx int) {
	line := con.content[con.current].(*canvas.Text)
	text := line.Text
	if text[con.headLen+con.inputLen:] == cursorChar {
		text = text[:len(text)-1]
	}
	line.Text = text[:len(text)-con.inputLen]
	lastInput := con.history[historyIdx]
	con.inputLen = len(lastInput)
	con.append(lastInput)
	canvas.Refresh(line)
}

func runProgram(con *console, app fyne.App) {
	dat, err := ioutil.ReadFile("../examples/rundis.runtime")
	if err != nil {
		fmt.Println(err)
	}
	src := string(dat)

	// replace newline characters for windows
	src = strings.Replace(src, "\r\n", "\n", -1)

	program := runtime.Tokenize(src)
	env := runtime.Parse(program)
	env.Out = func(content interface{}, ending string) {
		contentStr := fmt.Sprintf("%v", content)
		if ending == "\n" {
			for con.headLen+len(contentStr) > screenCols {
				con.headLen = 0
				con.appendLine(contentStr[:screenCols])
				contentStr = contentStr[screenCols:]
			}
			con.headLen += len(contentStr)
			con.appendLine(contentStr)

		} else {
			contentStr += ending

			for con.headLen+len(contentStr) > screenCols {
				con.headLen = 0
				con.appendLine(contentStr[:screenCols])
				contentStr = contentStr[screenCols:]
			}
			con.headLen += len(contentStr)
			con.append(contentStr)
		}
	}
	env.In = func() string {
		return <-con.input
	}
	runtime.Evaluate(program, env)
	time.Sleep(time.Second)
	app.Quit()
}

// Show starts a new console computer simulator
func Show(app fyne.App) {
	con := console{}
	app.Settings().SetTheme(&beebTheme{})

	window := app.NewWindow("Runtime Script")
	window.SetContent(con.loadUI())
	window.SetPadded(false)
	window.SetFixedSize(true)
	window.Resize(screenSize)

	window.Canvas().SetOnTypedRune(con.onRune)
	window.Canvas().SetOnTypedKey(con.onKey)
	window.Canvas().AddShortcut(&desktop.CustomShortcut{
		Modifier: desktop.ControlModifier,
		KeyName:  fyne.KeyD,
	}, func(fyne.Shortcut) {
		go func() {
			app.Quit()
		}()
	})

	go con.blink()
	window.Show()
	go runProgram(&con, app)
}
