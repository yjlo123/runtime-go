package main

import "fmt"

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
					token += string(src[i])
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
	}
}
