package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

	program := Tokenize(src)
	env := Parse(program)
	Evaluate(program, env)
}
