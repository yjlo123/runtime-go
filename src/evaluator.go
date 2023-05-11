package runtime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"reflect"
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

// ParseJSON ..
func ParseJSON(str string) *Value {
	if len(str) == 0 {
		return NewValue(nil)
	}

	if str[0] == '[' {
		list := &List{}
		content := str[1 : len(str)-1]
		vals := strings.Split(content, ",")
		for _, v := range vals {
			i, _ := strconv.Atoi(v)
			list.Push(NewValue(i))
		}
		return NewValue(list)
	}

	if str[0] == '{' {
		var result map[string]interface{}
		json.Unmarshal([]byte(str), &result)
		m := &Map{}
		for k, v := range result {
			m.Put(k, parseJSONRec(v))
		}
		return NewValue(m)
	}

	return NewValue(nil)
}

func parseJSONRec(data interface{}) *Value {
	kind := reflect.ValueOf(data).Kind()
	if kind == reflect.Map {
		mm := &Map{}
		for k, v := range data.(map[string]interface{}) {
			mm.Put(k, parseJSONRec(v))
		}
		return NewValue(mm)
	} else if kind == reflect.Slice {
		lst := &List{}
		for _, v := range data.([]interface{}) {
			lst.Push(parseJSONRec(v))
		}
		return NewValue(lst)
	} else if kind == reflect.String {
		return NewValue(data.(string))
	}
	return NewValue(nil)
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
					} else if val1.Type == ValueTypeNil && val2.Type == ValueTypeStr {
						// nil + int => str
						res = NewValue(string(val1.GetValue().(int)))
					} else {
						panic(fmt.Sprintf("add unsupported data type: %s, %s\n", val1.Type, val2.Type))
					}
				} else if cmd == "sub" {
					if val1.Type == ValueTypeInt && val2.Type == ValueTypeInt {
						// int - int => int
						res = NewValue(val1.GetValue().(int) - val2.GetValue().(int))
					} else if val1.Type == ValueTypeStr && val2.Type == ValueTypeNil {
						// str - nil => int
						res = NewValue(int(val2.GetValue().(string)[0]))
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
				env.AssignVar(ts[1], ParseJSON(jsonStr.GetValue().(string)))
			} else if cmd == "lod" {
				fileName := env.Express(ts[1]).GetValue().(string)
				fileData, err := ioutil.ReadFile(fileName)
				if err == nil {
					data := string(fileData)
					// replace newline characters for windows
					data = strings.Replace(data, "\r\n", "\n", -1)
					lines := strings.Split(data, "\n")
					env.AssignVar(ts[2], deserialize(lines))
				} else {
					//fmt.Println(err)
					env.AssignVar(ts[2], NewValue(nil))
				}
			} else if cmd == "sav" {
				fileName := env.Express(ts[1]).GetValue().(string)
				dataContent := env.Express(ts[2])
				data := serialize(dataContent)
				err := ioutil.WriteFile(fileName, data, 0644)
				if err != nil {
					fmt.Println(err)
				}
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

func serialize(content *Value) []byte {
	var data []byte
	if content.Type == ValueTypeStr {
		data = append(data, '\'')
		strVal := strings.Replace(content.Val, "\n", "\\n", -1)
		data = append(data, []byte(strVal)...)
		data = append(data, '\'')
	} else if content.Type == ValueTypeInt {
		data = []byte(content.Val)
	} else if content.Type == ValueTypeNil {
		data = []byte("$nil")
	} else if content.Type == ValueTypeList {
		list := content.GetValue().(*List)
		listLen := list.Len().GetValue().(int)
		data = append(data, []byte("[\n")...)
		for i := 0; i < listLen; i++ {
			data = append(data, serialize(list.GetByIndex(i))...)
		}
		data = append(data, ']')
	} else if content.Type == ValueTypeMap {
		m := content.GetValue().(*Map)
		keys := m.GetKeys()
		keysLen := keys.Len().GetValue().(int)
		data = append(data, []byte("{\n")...)
		for i := 0; i < keysLen; i++ {
			key := keys.GetByIndex(i).Val
			data = append(data, []byte(key)...)
			data = append(data, '\n')
			data = append(data, serialize(m.Get(key))...)
		}
		data = append(data, '}')
	}
	data = append(data, '\n')
	return []byte(data)
}

func deserialize(lines []string) *Value {
	if len(lines) == 0 {
		return NewValue(nil)
	} else if len(lines) == 1 {
		tempEnv := &Env{}
		line := strings.Replace(lines[0], "\\n", "\n", -1)
		return tempEnv.Express(line)
	} else {
		if lines[0] == "[" {
			list := &List{}
			for i := 1; i < len(lines); i++ {
				if lines[i] == "]" {
					break
				} else if lines[i] == "[" {
					start := i
					i++
					count := 1
					for i < len(lines) {
						if lines[i] == "[" {
							count++
						} else if lines[i] == "]" {
							count--
						}
						if count == 0 {
							list.Push(deserialize(lines[start : i+1]))
							break
						}
						i++
					}
				} else if lines[i] == "{" {
					start := i
					i++
					count := 1
					for i < len(lines) {
						if lines[i] == "{" {
							count++
						} else if lines[i] == "}" {
							count--
						}
						if count == 0 {
							list.Push(deserialize(lines[start : i+1]))
							break
						}
						i++
					}
				} else {
					list.Push(deserialize(lines[i : i+1]))
				}
			}
			return NewValue(list)
		} else if lines[0] == "{" {
			m := &Map{}
			for i := 1; i < len(lines); i++ {
				if lines[i] == "}" {
					break
				}
				key := lines[i]
				var val *Value
				i++
				if lines[i] == "}" {
					break
				} else if lines[i] == "{" {
					start := i
					i++
					count := 1
					for i < len(lines) {
						if lines[i] == "{" {
							count++
						} else if lines[i] == "}" {
							count--
						}
						if count == 0 {
							val = deserialize(lines[start : i+1])
							break
						}
						i++
					}
				} else if lines[i] == "[" {
					start := i
					i++
					count := 1
					for i < len(lines) {
						if lines[i] == "[" {
							count++
						} else if lines[i] == "]" {
							count--
						}
						if count == 0 {
							val = deserialize(lines[start : i+1])
							break
						}
						i++
					}
				} else {
					val = deserialize(lines[i : i+1])
				}
				m.Put(key, val)
			}
			return NewValue(m)
		}
	}
	return NewValue(nil)
}
