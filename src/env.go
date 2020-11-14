package main

import (
	"fmt"
	"strconv"
)

// Env ..
type Env struct {
	Labels map[string]int
	Vars   map[string]*Value // global vars
	Funcs  map[string]int
	stack  []*Frame
	Pc     int
}

// AdvancePc ..
func (env *Env) AdvancePc() {
	env.Pc++
}

// Express ..
func (env *Env) Express(expr string) *Value {
	if len(expr) == 0 {
		return nil
	}
	if expr[0] == '$' {
		// reference
		return env.GetVarVal(expr[1:])
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
	} else if expr[0] == '{' {
		return NewValue(&Map{})
	}

	// string ?
	return NewValue(expr)
}

// AssignVar ..
func (env *Env) AssignVar(varName string, val *Value) {
	lastFrame := env.GetFrame()
	if varName[0] == '_' && lastFrame != nil {
		// function scoped var
		lastFrame.Vars[varName] = val
	} else {
		// global var
		env.Vars[varName] = val
	}
}

// GetVarVal ..
func (env *Env) GetVarVal(varName string) *Value {
	vars := env.Vars // global
	lastFrame := env.GetFrame()
	if _, err := strconv.Atoi(varName); err == nil {
		// function param
		vars = lastFrame.Vars
	} else if (varName[0] == '_' || varName == "ret") && lastFrame != nil {
		// function scope
		vars = lastFrame.Vars
	}

	if val, ok := vars[varName]; ok {
		return val
	}
	return NewValue(nil)
}

// GotoLabelByName ..
func (env *Env) GotoLabelByName(name string) {
	env.Pc = env.Labels[name]
}

/* =========
   FRAME
   ========= */

//Frame ..
type Frame struct {
	FuncName string
	Pc       int
	Vars     map[string]*Value //local vars
}

func (f *Frame) String() string {
	return fmt.Sprintf("[%s:%d:%s]", f.FuncName, f.Pc, f.Vars)
}

// PushFrame ..
func (env *Env) PushFrame(funcName string, args []*Value) {
	env.stack = append(env.stack, env.NewFrame(funcName, args))
}

// PopFrame ..
func (env *Env) PopFrame() *Frame {
	frame := env.stack[len(env.stack)-1]
	env.stack = env.stack[:len(env.stack)-1]
	return frame
}

// GetFrame ..
func (env *Env) GetFrame() *Frame {
	if len(env.stack) > 0 {
		frame := env.stack[len(env.stack)-1]
		return frame
	}
	return nil
}

// NewFrame ..
func (env *Env) NewFrame(funcName string, args []*Value) *Frame {
	frameVars := make(map[string]*Value)
	for i, val := range args {
		frameVars[strconv.Itoa(i)] = val
	}
	return &Frame{
		FuncName: funcName,
		Pc:       env.Pc,
		Vars:     frameVars,
	}
}

// AssignReturnedVal ..
func (env *Env) AssignReturnedVal(val *Value) {
	lastFrame := env.GetFrame()
	if lastFrame == nil {
		env.Vars["ret"] = val
	} else {
		lastFrame.Vars["ret"] = val
	}
}
