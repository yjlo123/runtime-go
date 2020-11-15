package main

import (
	"io/ioutil"
	"strings"
)

func main() {

	dat, err := ioutil.ReadFile("../examples/runtime_script.runtime")
	if err != nil {
		panic(err)
	}
	src := string(dat)

	// replace newline characters for windows
	src = strings.Replace(src, "\r\n", "\n", -1)

	program := Tokenize(src)

	// for i, ts := range program {
	// 	fmt.Println(i)
	// 	fmt.Printf("%d: [%s]\n", i, strings.Join(ts, ", "))
	// }

	env := Parse(program)
	Evaluate(program, env)
	//fmt.Println(env.Funcs)
}
