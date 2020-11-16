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
	screenInsetX = 80
	screenInsetY = 62

	screenLines = 25
	screenCols  = 40
)

var (
	screenSize = fyne.Size{Width: 800, Height: 600}
	lineDelay  = time.Second / 10

	// Icon ..
	Icon = icon
)

type console struct {
	content []fyne.CanvasObject
	input   chan string // user input will be sent to this channel after pressing Enter

	overlay *canvas.Image
	current int

	bufInput []byte
	endInput bool
	nextAuto int
}

func (con *console) Read(p []byte) (n int, err error) {
	if con.endInput {
		con.endInput = false
		con.bufInput = nil
		p[0] = '\n'
		return 1, nil
	}
	con.bufInput = p

	if p[0] != 0 {
		return 1, nil
	}

	time.Sleep(lineDelay)
	return 0, nil
}

func (con *console) Write(p []byte) (n int, err error) {
	str := string(p)
	if str[len(str)-1] == '\n' {
		con.appendLine(str[:len(str)-1])
	} else {
		con.append(str)
	}
	return len(p), nil
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
}

func (con *console) newLine() {
	time.Sleep(lineDelay)
	text := con.content[con.current].(*canvas.Text)

	if len(text.Text) > 0 && text.Text[len(text.Text)-1] == '_' {
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
	if len(text.Text) > 0 && text.Text[len(text.Text)-1] == '_' {
		text.Text = text.Text[:len(text.Text)-1] + line + "_"
	} else {
		text.Text = text.Text + line
	}

	canvas.Refresh(text)
}

func (con *console) blink() {
	for {
		time.Sleep(time.Second / 2)
		line := con.content[con.current].(*canvas.Text)

		if line.Text == "" {
			continue
		}
		if line.Text[len(line.Text)-1] == '_' {
			line.Text = line.Text[:len(line.Text)-1]
		} else {
			line.Text = line.Text + "_"
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
	if con.bufInput != nil {
		con.bufInput[0] = byte(r) // TODO could we have typed another?
	}
	con.append(string(r))
}

func (con *console) onKey(ev *fyne.KeyEvent) {
	if con.bufInput != nil {
		if ev.Name == fyne.KeyReturn {
			con.endInput = true
		}
		return
	}
	switch ev.Name {
	case fyne.KeyReturn:
		//prog := ">"
		text := con.content[con.current].(*canvas.Text).Text
		// if len(text) > 1 {
		// 	text = text[1:]
		// 	if len(text) > 0 && text[len(text)-1] == '_' {
		// 		text = text[:len(text)-1]
		// 	}
		// 	prog = strings.TrimSpace(text) + "\n"
		// }
		con.appendLine("")
		//first := prog[0]
		// if first >= '0' && first <= '9' {
		// 	con.program += prog
		// } else {
		// commands that can't be called from within a program
		//cmd := strings.ToUpper(prog[:len(prog)-1])
		fmt.Println(text)
		con.input <- text[2:]

		// if cmd == "AUTO" {
		// 	con.nextAuto = 10
		// } else if cmd == "RUN" {
		// 	con.RUN()
		// } else if cmd == "NEW" {
		// 	con.NEW()
		// } else if cmd == "LIST" {
		// 	con.LIST()
		// } else if cmd == "QUIT" || cmd == "EXIT" {
		// 	con.QUIT(fyne.CurrentApp())
		// } else {
		// 	//con.runProg(prog)
		// }
		//}
		//con.append(">")
		if con.nextAuto > 0 {
			con.append(fmt.Sprintf("%d ", con.nextAuto))
			con.nextAuto += 10
		}
	case fyne.KeyBackspace:
		//line := con.content[con.current].(*canvas.Text)
		//text := line.Text[1:]
		// if len(text) > 0 && text[len(text)-1] == '_' {
		// 	text = text[:len(text)-1]
		// }
		// if len(text) > 0 {
		// 	line.Text = ">" + text[:len(text)-1]
		// 	canvas.Refresh(line)
		// }
	case fyne.KeyEscape:
		text := con.content[con.current].(*canvas.Text)
		if len(text.Text) == 0 || text.Text[0] != '>' {
			break
		}

		con.nextAuto = 0
		text.Text = ">"
		canvas.Refresh(text)
	}
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
		con.append("QUIT")
		con.appendLine("")
		go func() {
			fmt.Println("Quit")
		}()
	})

	fmt.Println("Restart")

	//con.append(">")
	//go con.blink()

	window.Show()

	go runProgram(&con)
}

func runProgram(con *console) {
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
		if ending == "\n" {
			con.appendLine(fmt.Sprintf("%v", content))
		} else {
			con.append(fmt.Sprintf("%v", content))
			con.append(fmt.Sprintf("%v", ending))
		}
	}
	env.In = func() string {
		return <-con.input
	}
	runtime.Evaluate(program, env)
}
