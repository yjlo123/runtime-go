package runtime

import (
	"fmt"
	"reflect"

	gt "github.com/leandroveronezi/go-terminal"
	"github.com/mattn/go-tty"
)

// Tokenize ..
func Tokenize(src string) [][]string {
	var tokens [][]string

	var line []string
	token := ""
	i := 0
	for i < len(src) {
		c := src[i]
		if c == '\n' || c == ' ' {
			// end of a token / line
			if len(token) > 0 {
				line = append(line, token)
			}
			token = ""
			if c == '\n' {
				tokens = append(tokens, line)
				line = []string{}
			}
		} else if c == '\'' || c == '"' {
			// string
			quote := c
			token += string(src[i])
			i++
			for {
				if i == len(src) || src[i] == '\n' {
					panic(fmt.Sprintf("Unterminated string: `%s`", token))
				}
				if src[i] == '\\' {
					i++
					if src[i] == 'n' {
						token += string('\n')
					} else {
						token += string(src[i])
					}
					i++
					continue
				}

				if src[i] == quote {
					break
				}
				token += string(src[i])
				i++
			}
			token += string(src[i])
		} else if c == '/' {
			for {
				if i == len(src) || src[i] == '\n' {
					i-- // to trigger newline in the next iteration
					break
				}
				i++
			}
		} else {
			token += string(c)
		}
		i++
	}
	if len(token) > 0 {
		line = append(line, token)
	}
	if len(line) > 0 {
		tokens = append(tokens, line)
	}
	//fmt.Println(tokens)
	return tokens
}

// Parse ..
func Parse(program [][]string) *Env {
	labels := make(map[string]int)
	funcs := make(map[string]*funcDetail)

	var currentFuncName string
	var currentFuncDetail *funcDetail
	for i, ts := range program {
		if len(ts) < 1 {
			continue
		}
		cmd := ts[0]
		if cmd[0] == '#' {
			if currentFuncName != "" {
				// func label
				currentFuncDetail.Labels[cmd[1:]] = i
			} else {
				// global label
				labels[cmd[1:]] = i
			}
		}
		if cmd == "def" {
			currentFuncName = ts[1]
			currentFuncDetail = &funcDetail{
				Pc:     i,
				Labels: make(map[string]int),
			}
		} else if cmd == "end" {
			funcs[currentFuncName] = currentFuncDetail
			currentFuncName = ""
			currentFuncDetail = nil
		}
	}
	return &Env{
		Labels: labels,
		Funcs:  funcs,
		Vars:   make(map[string]*Value),
		stack:  []*Frame{},
		Out: func(content interface{}, ending string) {
			if reflect.ValueOf(content).Kind() == reflect.String {
				cont := content.(string)
				if len(cont) > 3 && cont[:4] == "\\033" {
					return
				}
				//content = strings.Replace(content.(string), "\\033", "\x0cOn", -1)
			}

			if ending == "\n" {
				fmt.Println(content)
			} else {
				fmt.Print(content)
				fmt.Print(ending)
			}
		},
		In: func(env *Env) string {
			tty, err := tty.Open()
			if err != nil {
				fmt.Println(err)
			}
			defer tty.Close()

			// go func() {
			// 	for ws := range tty.SIGWINCH() {
			// 		fmt.Println("Resized", ws.W, ws.H)
			// 	}
			// }()

			input := ""
			cursor := 0
			consoleHistoryIndex := len(env.ConsoleHistory)

			clearInput := func() {
				gt.CursorRight(len(input) - cursor)
				cursor = len(input)
				for i := 0; i < cursor; i++ {
					fmt.Print("\b \b")
				}
				input = ""
				cursor = 0
			}

			setInput := func(text string) {
				clearInput()
				fmt.Print(text)
				input = text
				cursor = len(input)
			}

			clean, err := tty.Raw()
			if err != nil {
				fmt.Println(err)
			}
			defer clean()

			for {
				r, err := tty.ReadRune()
				if err != nil {
					fmt.Println(err)
				}

				if r == 0 {
					continue
				}

				if r == 27 && tty.Buffered() {
					// Arrow keys
					r, err = tty.ReadRune()
					if err == nil && r == 0x5b {
						r, err = tty.ReadRune()
						if err != nil {
							panic(err)
						}
						switch r {
						case 'A':
							// Arrow Up
							if consoleHistoryIndex > 0 {
								consoleHistoryIndex--
								record := env.ConsoleHistory[consoleHistoryIndex]
								setInput(record)
							}
							continue
						case 'B':
							// Arrow Down
							if consoleHistoryIndex < len(env.ConsoleHistory) {
								consoleHistoryIndex++
								if consoleHistoryIndex == len(env.ConsoleHistory) {
									clearInput()
								} else {
									setInput(env.ConsoleHistory[consoleHistoryIndex])
								}
							}
							continue
						case 'C':
							// Arrow Right
							if cursor < len(input) {
								gt.CursorRight(1)
								cursor++
							}
							continue
						case 'D':
							// Arrow Left
							if cursor > 0 {
								gt.CursorLeft(1)
								cursor--
							}
							continue
						}
					}
				}

				if r == 1 {
					// ^A
					gt.CursorLeft(cursor)
					cursor = 0
					continue
				} else if r == 3 {
					// ^C
					fmt.Print("^C")
					continue
				} else if r == 5 {
					// ^E
					gt.CursorRight(len(input) - cursor)
					cursor = len(input)
					continue
				} else if r == 8 || r == 127 {
					// Backspace
					// Win: 8, Linux: 127
					if input == "" || cursor == 0 {
						continue
					}
					left := input[:cursor-1]
					right := input[cursor:]
					input = left + right
					fmt.Print("\b \b" + right + " ")
					gt.CursorLeft(len(right) + 1)
					cursor--
					continue
				} else if r == 9 {
					// Tab
					continue
				} else if r == 13 {
					// Enter
					fmt.Print("\n\r")
					break
				}

				left := input[:cursor]
				right := input[cursor:]
				input = left + string(r) + right
				fmt.Printf("%c%s", r, right)
				cursor++
				gt.CursorLeft(len(right))
			}
			// record input in env history
			if len(input) > 0 {
				env.ConsoleHistory = append(env.ConsoleHistory, input)
				if len(env.ConsoleHistory) > 20 {
					env.ConsoleHistory = env.ConsoleHistory[1:]
				}
			}

			return input
		},
		loops: make(map[string]*loopDetail),
	}
}
