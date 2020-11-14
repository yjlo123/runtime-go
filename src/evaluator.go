package main

import (
	"fmt"
	"strconv"
	"time"
)

// Evaluate ..
func Evaluate(program [][]string, env *Env) {
	//pc := 0
	for {
		if env.Pc >= len(program) {
			break
		}
		ts := program[env.Pc]
		if len(ts) > 0 {
			cmd := ts[0]
			if cmd == "prt" {
				fmt.Println(env.Express(ts[1]).GetValue())
			} else if cmd == "slp" {
				time.Sleep(time.Duration(env.Express(ts[1]).GetValue().(int)) * time.Millisecond)
			} else if cmd == "let" {
				val := env.Express(ts[2])
				env.AssignVar(ts[1], val)
			} else if cmd == "inp" {
				var input string
				fmt.Scanln(&input)
				env.AssignVar(ts[1], NewValue(input))
			} else if cmd == "int" {
				val := env.Express(ts[2])
				intVal, err := strconv.Atoi(val.GetValue().(string))
				if err != nil {
					env.AssignVar(ts[1], nil)
				}
				env.AssignVar(ts[1], NewValue(intVal))
			} else if cmd == "str" {
				val := env.Express(ts[2])
				strVal := strconv.Itoa(val.GetValue().(int))
				env.AssignVar(ts[1], NewValue(strVal))
			} else if cmd == "typ" {
				val := env.Express(ts[2])
				env.AssignVar(ts[1], NewValue(val.Type))
			} else if cmd == "add" || cmd == "sub" || cmd == "mul" || cmd == "div" {
				val1 := env.Express(ts[2])
				val2 := env.Express(ts[3])
				var res int
				if cmd == "add" {
					res = val1.GetValue().(int) + val2.GetValue().(int)
				} else if cmd == "sub" {
					res = val1.GetValue().(int) - val2.GetValue().(int)
				} else if cmd == "mul" {
					res = val1.GetValue().(int) * val2.GetValue().(int)
				} else if cmd == "div" {
					res = val1.GetValue().(int) / val2.GetValue().(int)
				} else if cmd == "mod" {
					res = val1.GetValue().(int) % val2.GetValue().(int)
				}
				env.AssignVar(ts[1], NewValue(res))
				// LIST
			} else if cmd == "psh" {
				list := env.Express(ts[1]).GetValue().(*List)
				list.Push(env.Express(ts[2]))
			} else if cmd == "pop" {
				list := env.Express(ts[1]).GetValue().(*List)
				val := list.Pop()
				env.AssignVar(ts[2], val)
			} else if cmd == "pol" {
				list := env.Express(ts[1]).GetValue().(*List)
				val := list.Poll()
				env.AssignVar(ts[2], val)
				// MAP
			} else if cmd == "get" {
				m := env.Express(ts[1]).GetValue().(*Map)
				key := env.Express(ts[2]).GetValue().(string)
				val := m.Get(key)
				env.AssignVar(ts[3], val)
			} else if cmd == "put" {
				m := env.Express(ts[1]).GetValue().(*Map)
				key := env.Express(ts[2]).GetValue().(string)
				val := env.Express(ts[3])
				m.Put(key, val)
			} else if cmd == "del" {
				m := env.Express(ts[1]).GetValue().(*Map)
				key := env.Express(ts[2]).GetValue().(string)
				m.Delete(key)
			} else if cmd == "key" {
				m := env.Express(ts[1]).GetValue().(*Map)
				env.AssignVar(ts[2], NewValue(m.GetKeys()))
				// JUMP
			} else if cmd == "jmp" {
				env.GotoLabelByName(ts[3])
			} else if cmd == "jne" || cmd == "jeq" || cmd == "jlt" || cmd == "jgt" {
				val1 := env.Express(ts[1])
				val2 := env.Express(ts[2])
				if (cmd == "jne" && !val1.Equals(val2)) ||
					(cmd == "jeq" && val1.Equals(val2)) ||
					(cmd == "jlt" && val2.IsGreaterThan(val1)) ||
					(cmd == "jgt" && val1.IsGreaterThan(val2)) {
					env.GotoLabelByName(ts[3])
				}
				// FUNC
			} else if cmd == "def" {
				for env.Pc <= len(program) && program[env.Pc][0] != "end" {
					env.AdvancePc()
				}
			} else if cmd == "cal" {
				funcName := ts[1]
				var args []*Value
				for _, v := range ts[2:] {
					args = append(args, env.Express(v))
				}
				env.PushFrame(funcName, args)
				env.Pc = env.Funcs[funcName]
			} else if cmd == "end" {
				frame := env.PopFrame()
				env.Pc = frame.Pc
			} else if cmd == "ret" {
				retValue := env.Express(ts[1])
				frame := env.PopFrame()
				env.AssignReturnedVal(retValue)
				env.Pc = frame.Pc
			}
		}
		env.AdvancePc()
	}
}
