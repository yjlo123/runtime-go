package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	runtime "github.com/yjlo123/runtime-go/src"
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

	program := runtime.Tokenize(src)
	env := runtime.Parse(program)
	runtime.Evaluate(program, env)
}
