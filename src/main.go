package main

import (
	"io/ioutil"
	"strings"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	dat, err := ioutil.ReadFile("../examples/list.runtime")
	checkErr(err)
	src := string(dat)
	// replace newline character for windows
	src = strings.Replace(src, "\r\n", "\n", -1)

	program := Tokenize(src)

	// for i, ts := range program {
	// 	fmt.Println(i)
	// 	fmt.Printf("%d: [%s]\n", i, strings.Join(ts, ", "))
	// }

	env := Parse(program)
	Evaluate(program, env)
	//fmt.Println(env.Vars)
}
