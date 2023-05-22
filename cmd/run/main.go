package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"

	"github.com/mattn/go-tty"
	rts "github.com/yjlo123/runtime-go/src"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Printf("Usage: runtime <file_path>\n")
		return
	}

	dat, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
	}
	src := string(dat)

	// replace newline characters for windows
	src = strings.Replace(src, "\r\n", "\n", -1)

	program := rts.Tokenize(src)

	tty, err := tty.Open()
	if err != nil {
		fmt.Println(err)
	}
	defer tty.Close()

	env := rts.Parse(program, tty)

	go func() {
		for ws := range tty.SIGWINCH() {
			env.Vars["term_w"] = rts.NewValue(ws.W)
			env.Vars["term_h"] = rts.NewValue(ws.H)
		}
	}()

	w, h, _ := tty.Size()

	// For Moon OS
	env.Vars["term_w"] = rts.NewValue(w)
	env.Vars["term_h"] = rts.NewValue(h)
	env.Vars["os_host"] = rts.NewValue(runtime.GOOS + " (golang)")

	rts.Evaluate(program, env)
}
