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

func advancePcToIfEnd(program [][]string, env *Env) {
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
		}
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

func backPcToLoopHead(program [][]string, env *Env) {
	forStack := 0
	env.Pc--
	for env.Pc > 0 {
		if len(program[env.Pc]) > 0 {
			currentCmd := program[env.Pc][0]
			if currentCmd == "for" {
				if forStack == 0 {
					env.Pc--
					return
				}
				forStack--
			} else if currentCmd == "nxt" {
				forStack++
			}
		}
		env.Pc--
	}
}

func advancetoLoopEnd(program [][]string, env *Env) {
	env.Pc++ // first 'for'
	forStack := 0
	for env.Pc <= len(program) {
		if len(program[env.Pc]) > 0 {
			currentCmd := program[env.Pc][0]
			if currentCmd == "for" {
				forStack++
			} else if currentCmd == "nxt" {
				if forStack == 0 {
					return
				}
				forStack--
			}
		}
		env.Pc++
	}
}

// Evaluate ..
func Evaluate(program [][]string, env *Env) *Env {
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
				input := env.In(env)
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
			} else if cmd == "add" || cmd == "sub" || cmd == "mul" || cmd == "div" || cmd == "mod" {
				val1 := env.Express(ts[2])
				val2 := env.Express(ts[3])
				var res *Value
				if cmd == "add" {
					if val1.Type == ValueTypeInt && val2.Type == ValueTypeInt {
						// int + int => int
						res = NewValue(val1.GetValue().(int) + val2.GetValue().(int))
					} else if val1.Type == ValueTypeStr && val2.Type == ValueTypeStr {
						// str + str => int
						res = NewValue(val1.GetValue().(string) + val2.GetValue().(string))
					} else if val1.Type == ValueTypeStr && val2.Type == ValueTypeInt {
						// str + int => str
						res = NewValue(val1.GetValue().(string) + strconv.Itoa(val2.GetValue().(int)))
					} else if val1.Type == ValueTypeInt && val2.Type == ValueTypeStr {
						// int + str => str
						res = NewValue(strconv.Itoa(val1.GetValue().(int)) + val2.GetValue().(string))
					} else if val1.Type == ValueTypeNil && val2.Type == ValueTypeInt {
						// nil + int => str
						res = NewValue(string(val2.GetValue().(int)))
					} else {
						panic(fmt.Sprintf("add unsupported data type: %s, %s\n", val1.Type, val2.Type))
					}
				} else if cmd == "sub" {
					if val1.Type == ValueTypeInt && val2.Type == ValueTypeInt {
						// int - int => int
						res = NewValue(val1.GetValue().(int) - val2.GetValue().(int))
					} else if val1.Type == ValueTypeStr && val2.Type == ValueTypeNil {
						// str - nil => int
						res = NewValue(int(val1.GetValue().(string)[0]))
					} else {
						panic(fmt.Sprintf("sub unsupported data type: %s, %s\n", val1.Type, val2.Type))
					}
				} else if cmd == "mul" {
					if val1.Type == ValueTypeInt && val2.Type == ValueTypeInt {
						// int * int => int
						res = NewValue(val1.GetValue().(int) * val2.GetValue().(int))
					} else if val1.Type == ValueTypeStr && val2.Type == ValueTypeInt {
						// str * int => str
						res = NewValue(strings.Repeat(val1.GetValue().(string), val2.GetValue().(int)))
					} else {
						panic(fmt.Sprintf("mul unsupported data type: %s, %s\n", val1.Type, val2.Type))
					}
				} else if cmd == "div" {
					res = NewValue(val1.GetValue().(int) / val2.GetValue().(int))
				} else if cmd == "mod" {
					res = NewValue(val1.GetValue().(int) % val2.GetValue().(int))
				}
				env.AssignVar(ts[1], res)
				// LIST
			} else if cmd == "psh" {
				listVal := env.Express(ts[1])
				for _, v := range ts[2:] {
					switch listVal.Type {
					case ValueTypeList:
						list := listVal.GetValue().(*List)
						list.Push(env.Express(v))
					case ValueTypeStr:
						str := env.Express(v).GetValue().(string)
						listVal.Val += str
					}
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
							val = NewValue(str[len(str)-1:])
						}
					} else {
						val = NewValue("")
					}

				default:
					fmt.Println(env.GetFrame().Vars)
					panic(fmt.Sprintf("%s invalid data type: %s %s", cmd, ts, listVal.Type))
				}
				env.AssignVar(ts[2], val)
				// MAP
			} else if cmd == "get" || cmd == "put" || cmd == "del" {
				ds := env.Express(ts[1])
				keyValue := env.Express(ts[2])
				if ds.Type == ValueTypeMap {
					// map
					m := ds.GetValue().(*Map)
					key := ""
					if keyValue.Type == ValueTypeStr {
						key = keyValue.GetValue().(string)
					} else if keyValue.Type == ValueTypeInt {
						key = strconv.Itoa(keyValue.GetValue().(int))
					} else if keyValue.Type == ValueTypeNil {
						// key is $nil
						key = ""
					} else {
						panic(fmt.Sprintf("%s: Invalid key data type: %s", cmd, keyValue.Type))
					}

					if cmd == "get" {
						env.AssignVar(ts[3], m.Get(key))
					} else if cmd == "put" {
						m.Put(key, env.Express(ts[3]))
					} else if cmd == "del" {
						m.Delete(key)
					}
				} else if ds.Type == ValueTypeList {
					// list
					l := ds.GetValue().(*List)
					idx := keyValue.GetValue().(int)
					if cmd == "get" {
						env.AssignVar(ts[3], l.GetByIndex(idx))
					} else if cmd == "put" {
						l.SetByIndex(idx, env.Express(ts[3]))
					}
				} else if ds.Type == ValueTypeStr {
					// string
					s := ds.GetValue().(string)
					idx := keyValue.GetValue().(int)
					if cmd == "get" {
						val := NewValue("")
						if idx < len(s) {
							val = NewValue(string(s[idx]))
						}
						env.AssignVar(ts[3], val)
					} else if cmd == "put" {
						c := env.Express(ts[3]).GetValue().(string)
						newVal := NewValue(string(s[:idx] + c + s[idx+1:]))
						// replace string value (removing $)
						env.AssignVar(ts[1][1:], newVal)
					}
				}
			} else if cmd == "key" {
				m := env.Express(ts[1]).GetValue().(*Map)
				env.AssignVar(ts[2], NewValue(m.GetKeys()))
			} else if cmd == "len" {
				ds := env.Express(ts[1])
				if ds.Type == ValueTypeList {
					l := ds.GetValue().(*List)
					env.AssignVar(ts[2], l.Len())
				} else if ds.Type == ValueTypeStr {
					env.AssignVar(ts[2], NewValue(len(ds.GetValue().(string))))
				} else if ds.Type == ValueTypeMap {
					env.AssignVar(ts[2], ds.MapPtr.GetKeys().Len())
				}
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
				timeType := env.Express(ts[2]).String()
				if timeType[0] == '\'' {
					timeType = timeType[1 : len(timeType)-1]
				}
				res := 0
				now := time.Now()
				switch timeType {
				case "now":
					res = int(now.UnixNano() / int64(time.Millisecond))
				case "year":
					res = now.Year()
				case "month":
					res = int(now.Month()) - 1
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
				advancePcToIfEnd(program, env)
			} else if cmd == "for" {
				varName := ts[1]
				rangeVal := env.Express(ts[2])
				val, loopExists := env.loops[varName]
				if !loopExists || val.pc != env.Pc {
					var rangeList []*Value
					if rangeVal.Type == ValueTypeInt {
						rangeInt, _ := strconv.Atoi(rangeVal.Val)
						rangeList = make([]*Value, rangeInt)
						for i := range rangeList {
							rangeList[i] = NewValue(i)
						}
					} else if rangeVal.Type == ValueTypeList {
						rangeList = rangeVal.ListPtr.ToValueArray()
					} else if rangeVal.Type == ValueTypeStr {
						charArr := strings.Split(rangeVal.Val, "")
						for _, char := range charArr {
							rangeList = append(rangeList, NewValue(char))
						}
					} else if rangeVal.Type == ValueTypeMap {
						rangeList = rangeVal.MapPtr.GetKeys().ToValueArray()
					}

					env.loops[varName] = &loopDetail{
						items: rangeList,
						pc:    env.Pc, // to prevent the same var names
						index: 0,
					}
				}

				loopState := env.loops[varName]
				if loopState.index >= len(loopState.items) {
					delete(env.loops, varName)
					advancetoLoopEnd(program, env)
				} else {
					env.AssignVar(varName, loopState.items[loopState.index])
					loopState.index++
				}
			} else if cmd == "nxt" {
				backPcToLoopHead(program, env)
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
			} else if cmd == "prs" {
				jsonStr := env.Express((ts[2]))
				data := Deserialize(jsonStr.GetValue().(string))
				env.AssignVar(ts[1], data)
			} else if cmd == "lod" {
				fileName := env.Express(ts[1]).GetValue().(string)
				fileData, err := ioutil.ReadFile(fileName)
				if err == nil {
					data := string(fileData)
					// replace newline characters for windows
					data = strings.Replace(data, "\r\n", "\n", -1)
					env.AssignVar(ts[2], Deserialize(data))
				} else {
					//fmt.Println(err)
					env.AssignVar(ts[2], NewValue(nil))
				}
			} else if cmd == "sav" {
				fileName := env.Express(ts[1]).GetValue().(string)
				dataContent := env.Express(ts[2])
				data := Serialize(dataContent)
				err := ioutil.WriteFile(fileName, data, 0644)
				if err != nil {
					fmt.Println(err)
				}
			} else if handler, ok := env.Extended[cmd]; ok {
				// extended commands
				args := []*Value{}
				for _, arg := range ts[1:] {
					args = append(args, env.Express(arg))
				}
				handler(env, args)
			} else if cmd == "test_init" {
				env.AssignVar("test_pass", NewValue(0))
				env.AssignVar("test_fail", NewValue(0))
			} else if cmd == "test_assert" {
				val := env.Express(ts[1])
				expect := env.Express(ts[2])
				if !val.Equals(expect) {
					env.AssignVar("test_fail", NewValue(env.Express("$test_fail").GetValue().(int)+1))
					fmt.Printf("Test failed. Line:%d, expected:%s, got:%s\n", env.Pc+1, expect, val)
				} else {
					env.AssignVar("test_pass", NewValue(env.Express("$test_pass").GetValue().(int)+1))
				}
			}
		}
		env.AdvancePc()
	}
	return env
}

func evaluateFuncCall(program [][]string, env *Env, funcName string, args []string) {
	callLine := []string{"cal", funcName}
	callLine = append(callLine, args...)
	extendedProgram := append(program, callLine)
	lastLine := []string{"let", "_", "0"}
	extendedProgram = append(extendedProgram, lastLine)
	backupPc := env.Pc
	env.Pc = len(extendedProgram) - 2
	Evaluate(extendedProgram, env)
	env.Pc++
	for env.Pc < len(extendedProgram) {
		fmt.Println(env.Pc)
		Evaluate(extendedProgram, env)
		env.Pc++
	}
	env.Pc = backupPc
}
