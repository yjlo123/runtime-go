package main

import (
	"fmt"
	"strconv"
	"time"
)

func expression(env *Env, expr string) *Value {
	if len(expr) == 0 {
		return nil
	}
	if expr[0] == '$' {
		// reference
		return env.GetVarVal(expr[1:len(expr)])
	} else if num, err := strconv.Atoi(expr); err == nil {
		// integer
		return NewValue(num)
	} else if expr[0] == '\'' && expr[len(expr)-1] == '\'' {
		// string
		if len(expr) > 1 {
			expr = expr[1 : len(expr)-1]
		}
		return NewValue(expr)
	} else if expr[0] == '[' {
		return NewValue(&List{})
	}

	// string ?
	return NewValue(expr)
}

// Evaluate ..
func Evaluate(program [][]string, env *Env) {
	pc := 0
	for {
		if pc >= len(program) {
			break
		}
		ts := program[pc]
		if len(ts) > 0 {
			cmd := ts[0]
			if cmd == "prt" {
				fmt.Println(expression(env, ts[1]).GetValue())
			} else if cmd == "slp" {
				time.Sleep(time.Duration(expression(env, ts[1]).GetValue().(int)) * time.Millisecond)
			} else if cmd == "let" {
				val := expression(env, ts[2])
				env.AssignVar(ts[1], val)
			} else if cmd == "inp" {
				var input string
				fmt.Scanln(&input)
				env.AssignVar(ts[1], NewValue(input))
			} else if cmd == "int" {
				val := expression(env, ts[2])
				intVal, err := strconv.Atoi(val.GetValue().(string))
				if err != nil {
					env.AssignVar(ts[1], nil)
				}
				env.AssignVar(ts[1], NewValue(intVal))
			} else if cmd == "add" {
				val1 := expression(env, ts[2])
				val2 := expression(env, ts[3])
				res := val1.GetValue().(int) + val2.GetValue().(int)
				env.AssignVar(ts[1], NewValue(res))
			} else if cmd == "sub" {
				val1 := expression(env, ts[2])
				val2 := expression(env, ts[3])
				res := val1.GetValue().(int) - val2.GetValue().(int)
				env.AssignVar(ts[1], NewValue(res))
			} else if cmd == "div" {
				val1 := expression(env, ts[2])
				val2 := expression(env, ts[3])
				res := val1.GetValue().(int) / val2.GetValue().(int)
				env.AssignVar(ts[1], NewValue(res))
			} else if cmd == "psh" {
				list := expression(env, ts[1]).GetValue().(*List)
				list.Push(expression(env, ts[2]))
			} else if cmd == "pop" {
				list := expression(env, ts[1])
				val := list.GetValue().(*List).Pop()
				env.AssignVar(ts[2], val)
			} else if cmd == "pol" {
				list := expression(env, ts[1])
				val := list.GetValue().(*List).Poll()
				env.AssignVar(ts[2], val)
			} else if cmd == "jne" {
				val1 := expression(env, ts[1])
				val2 := expression(env, ts[2])
				if !val1.Equals(val2) {
					// TODO check if label exists
					pc = env.Labels[ts[3]]
				}
			}
		}
		pc++
	}
}
