package runtime

import (
	"fmt"
	"reflect"

	"github.com/pkg/term"
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

func readChar() (ascii int, keyCode int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".

		// Since there are no ASCII codes for arrow keys, we use
		// Javascript key codes.
		if bytes[2] == 65 {
			// Up
			keyCode = 38
		} else if bytes[2] == 66 {
			// Down
			keyCode = 40
		} else if bytes[2] == 67 {
			// Right
			keyCode = 39
		} else if bytes[2] == 68 {
			// Left
			keyCode = 37
		}
	} else if numRead == 1 {
		ascii = int(bytes[0])
	} else {
		// Two characters read??
	}
	t.Restore()
	t.Close()
	return
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
		In: func() string {
			input := ""
			for {
				ascii, keyCode, err := readChar()
				char := rune(ascii)
				if err != nil {
					fmt.Println(err)
					break
				}

				if char == '\r' {
					fmt.Print("\n\r")
					break
				}
				if char == '\t' {
					continue
				}
				if char == '\u007f' {
					if input == "" {
						continue
					}
					input = input[:len(input)-1]
					fmt.Printf("\b \b")
					continue
				}

				if keyCode == 37 {
					fmt.Print("<")
				} else if keyCode == 38 {
					fmt.Print("^")
				} else if keyCode == 39 {
					fmt.Print(">")
				} else if keyCode == 40 {
					fmt.Print("v")
				}

				input += string(char)
				fmt.Printf("%c", char)
			}
			return input
		},
		loops: make(map[string]*loopDetail),
	}
}
