package main

import (
	"bufio"
	"os"
	"strings"

	runtime "github.com/yjlo123/runtime-go/src"
)

func readLine() string {
	consoleReader := bufio.NewReader(os.Stdin)
	input, _ := consoleReader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "\n", -1)
	return input[0 : len(input)-1]
}

func main() {
	program := runtime.Tokenize("")
	env := runtime.Parse(program)
	//fmt.Println(env.Vars)
	line := readLine()
	for line != "quit" {
		program := runtime.Tokenize(line)
		env.Pc = 0
		env = runtime.Evaluate(program, env)
		//fmt.Println(env.Vars)
		line = readLine()
	}

}
