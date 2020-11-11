package main

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
					panic("Unterminated string")
				}
				if src[i] == quote {
					break
				}
				token += string(src[i])
				i++
			}
			token += string(src[i])
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
	return tokens
}

// Parse ..
func Parse(program [][]string) *Env {
	labels := make(map[string]int)
	for i, ts := range program {
		if len(ts) < 1 {
			continue
		}
		cmd := ts[0]
		if cmd[0] == '#' {
			labels[cmd[1:len(cmd)]] = i
		}
	}
	return &Env{
		Labels: labels,
		Vars:   make(map[string]*Value),
	}
}
