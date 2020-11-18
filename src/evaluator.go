package runtime

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func advancePcUntil(program [][]string, env *Env, cmd string) {
	for env.Pc <= len(program) && (len(program[env.Pc]) < 1 || program[env.Pc][0] != cmd) {
		env.AdvancePc()
	}
}

func advancePcToIfFalse(program [][]string, env *Env) {
	env.AdvancePc() // first ife or ifg
	nestedIfCount := 0
	for env.Pc <= len(program) {
		if len(program[env.Pc]) < 1 {
			env.AdvancePc()
			continue
		}
		currentCmd := program[env.Pc][0]
		if currentCmd == "ife" || currentCmd == "ifg" {
			nestedIfCount++
		} else if currentCmd == "fin" {
			if nestedIfCount == 0 {
				return
			}
			nestedIfCount--
		} else if currentCmd == "els" {
			if nestedIfCount == 0 {
				return
			}
		}
		env.AdvancePc()
	}
}

// Evaluate ..
func Evaluate(program [][]string, env *Env) {
	//pc := 0
	rand.Seed(time.Now().UnixNano())
	for {
		if env.Pc >= len(program) {
			break
		}
		ts := program[env.Pc]
		if len(ts) > 0 {
			cmd := ts[0]
			if cmd == "prt" {
				ending := "\n"
				if len(ts) > 2 {
					ending = env.Express(ts[2]).GetValue().(string)
				}
				env.Out(env.Express(ts[1]).GetValue(), ending)
			} else if cmd == "slp" {
				time.Sleep(time.Duration(env.Express(ts[1]).GetValue().(int)) * time.Millisecond)
			} else if cmd == "let" {
				val := env.Express(ts[2])
				env.AssignVar(ts[1], val)
			} else if cmd == "inp" {
				input := env.In()
				env.AssignVar(ts[1], NewValue(input))
			} else if cmd == "int" {
				val := env.Express(ts[2])
				switch val.Type {
				case ValueTypeInt:
					intVal := val.GetValue().(int)
					env.AssignVar(ts[1], NewValue(intVal))
				case ValueTypeStr:
					intVal, err := strconv.Atoi(val.GetValue().(string))
					if err != nil {
						env.AssignVar(ts[1], NewValue(nil))
					} else {
						env.AssignVar(ts[1], NewValue(intVal))
					}
				}
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
				var res *Value
				if cmd == "add" {
					if val1.Type == ValueTypeInt && val2.Type == ValueTypeInt {
						// int + int
						res = NewValue(val1.GetValue().(int) + val2.GetValue().(int))
					} else if val1.Type == ValueTypeStr && val2.Type == ValueTypeStr {
						// str + str
						res = NewValue(val1.GetValue().(string) + val2.GetValue().(string))
					} else if val1.Type == ValueTypeStr && val2.Type == ValueTypeInt {
						// str + int
						res = NewValue(val1.GetValue().(string) + strconv.Itoa(val2.GetValue().(int)))
					} else if val1.Type == ValueTypeInt && val2.Type == ValueTypeStr {
						// int + str
						res = NewValue(strconv.Itoa(val1.GetValue().(int)) + val2.GetValue().(string))
					} else {
						fmt.Println(ts)
						panic(fmt.Sprintf("add unsupported data type: %s, %s\n", val1.Type, val2.Type))
					}
				} else if cmd == "sub" {
					res = NewValue(val1.GetValue().(int) - val2.GetValue().(int))
				} else if cmd == "mul" {
					res = NewValue(val1.GetValue().(int) * val2.GetValue().(int))
				} else if cmd == "div" {
					res = NewValue(val1.GetValue().(int) / val2.GetValue().(int))
				} else if cmd == "mod" {
					res = NewValue(val1.GetValue().(int) % val2.GetValue().(int))
				}
				env.AssignVar(ts[1], res)
				// LIST
			} else if cmd == "psh" {
				listVal := env.Express(ts[1])
				switch listVal.Type {
				case ValueTypeList:
					list := listVal.GetValue().(*List)
					list.Push(env.Express(ts[2]))
				case ValueTypeStr:
					str := env.Express(ts[2]).GetValue().(string)
					listVal.Val += str
				}
			} else if cmd == "pol" || cmd == "pop" {
				listVal := env.Express(ts[1])
				var val *Value
				switch listVal.Type {
				case ValueTypeList:
					list := listVal.GetValue().(*List)
					if cmd == "pol" {
						val = list.Poll()
					} else {
						val = list.Pop()
					}
				case ValueTypeStr:
					str := listVal.GetValue().(string)
					if len(str) > 0 {
						if cmd == "pol" {
							listVal.Val = str[1:]
							val = NewValue(str[0:1])
						} else {
							listVal.Val = str[0 : len(str)-1]
							val = NewValue(str[len(str)-1 : len(str)])
						}
					} else {
						val = NewValue("")
					}

				default:
					fmt.Println(env.GetFrame().Vars)
					panic(fmt.Sprintf("pol invalid data type: %s %s", ts, listVal.Type))
				}
				env.AssignVar(ts[2], val)
				// MAP
			} else if cmd == "get" || cmd == "put" || cmd == "del" {
				m := env.Express(ts[1]).GetValue().(*Map)
				keyValue := env.Express(ts[2])
				key := ""
				if keyValue.Type == ValueTypeStr {
					key = keyValue.GetValue().(string)
				} else if keyValue.Type == ValueTypeInt {
					key = strconv.Itoa(keyValue.GetValue().(int))
				} else {
					panic(fmt.Sprintf("%s: Invalid key data type", cmd))
				}

				if cmd == "get" {
					val := m.Get(key)
					env.AssignVar(ts[3], val)
				} else if cmd == "put" {
					val := env.Express(ts[3])
					m.Put(key, val)
				} else if cmd == "del" {
					m.Delete(key)
				}
			} else if cmd == "key" {
				m := env.Express(ts[1]).GetValue().(*Map)
				env.AssignVar(ts[2], NewValue(m.GetKeys()))
				// JUMP
			} else if cmd == "jmp" {
				env.GotoLabelByName(ts[1])
			} else if cmd == "jne" || cmd == "jeq" || cmd == "jlt" || cmd == "jgt" {
				val1 := env.Express(ts[1])
				val2 := env.Express(ts[2])
				if (cmd == "jne" && !val1.Equals(val2)) ||
					(cmd == "jeq" && val1.Equals(val2)) ||
					(cmd == "jlt" && val2.IsGreaterThan(val1)) ||
					(cmd == "jgt" && val1.IsGreaterThan(val2)) {
					env.GotoLabelByName(ts[3])
				}
			} else if cmd == "rnd" {
				val1 := env.Express(ts[2]).GetValue().(int)
				val2 := env.Express(ts[3]).GetValue().(int)
				randInt := rand.Intn(val2 - val1)
				env.AssignVar(ts[1], NewValue(val1+randInt))
			} else if cmd == "tim" {
				timeType := ts[2]
				res := 0
				now := time.Now()
				switch timeType {
				case "now":
					res = int(now.UnixNano() / int64(time.Millisecond))
				case "year":
					res = now.Year()
				case "month":
					res = int(now.Month())
				case "date":
					res = now.Day()
				case "day":
					res = int(now.Weekday())
				case "hour":
					res = now.Hour()
				case "minute":
					res = now.Minute()
				case "second":
					res = now.Second()
				case "milli":
					res = int(now.UnixNano()/int64(time.Millisecond)) % 1000
				}
				env.AssignVar(ts[1], NewValue(res))
				// IF_ELSE
			} else if cmd == "ife" || cmd == "ifg" {
				val1 := env.Express(ts[1])
				val2 := env.Express(ts[2])
				if (!val1.Equals(val2) && cmd == "ife") ||
					(!val1.IsGreaterThan(val2) && cmd == "ifg") {
					advancePcToIfFalse(program, env)
				}
			} else if cmd == "els" {
				advancePcUntil(program, env, "fin")
				// FUNC
			} else if cmd == "def" {
				advancePcUntil(program, env, "end")
			} else if cmd == "cal" {
				funcName := ts[1]
				var args []*Value
				for _, v := range ts[2:] {
					argVal := env.Express(v)
					if argVal.Type == ValueTypeList || argVal.Type == ValueTypeMap {
						// pass by reference
						args = append(args, argVal)
					} else {
						// pass by value
						args = append(args, argVal.MakeCopy())
					}
				}
				env.PushFrame(funcName, args)
				env.Pc = env.Funcs[funcName].Pc
			} else if cmd == "end" {
				frame := env.PopFrame()
				env.Pc = frame.Pc
			} else if cmd == "ret" {
				if len(ts) > 1 {
					retValue := env.Express(ts[1])
					frame := env.PopFrame()
					env.AssignReturnedVal(retValue)
					env.Pc = frame.Pc
				} else {
					frame := env.PopFrame()
					env.Pc = frame.Pc
				}
			} else if cmd == "lod" {
				fileName := env.Express(ts[1]).GetValue().(string)
				fileData, err := ioutil.ReadFile(fileName)
				if err != nil {
					fmt.Println(err)
				}
				data := string(fileData)
				// replace newline characters for windows
				data = strings.Replace(data, "\r\n", "\n", -1)
				fmt.Println(data)
				env.AssignVar(ts[2], NewValue(data))
			} else if cmd == "sav" {
				fileName := env.Express(ts[1]).GetValue().(string)
				dataContent := env.Express(ts[2])
				var data []byte
				if dataContent.Type == ValueTypeStr {
					data = []byte(dataContent.GetValue().(string))
				} else if dataContent.Type == ValueTypeList {
					list := dataContent.GetValue().(*List)
					for {
						line := list.Poll()
						if line.GetValue() == nil {
							break
						}
						if len(data) > 0 {
							data = append(data, '\n')
						}
						data = append(data, []byte(line.GetValue().(string))...)
					}
				}

				err := ioutil.WriteFile(fileName, data, 0644)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		env.AdvancePc()
	}
}
