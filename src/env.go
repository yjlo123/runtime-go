package main

import "fmt"

// Env ..
type Env struct {
	Labels map[string]int
	Vars   map[string]*Value // global vars
	Funcs  map[string]int
	stack  []*Frame
}

// AssignVar ..
func (env *Env) AssignVar(varName string, val *Value) {
	env.Vars[varName] = val
}

// GetVarVal ..
func (env *Env) GetVarVal(varName string) *Value {
	if val, ok := env.Vars[varName]; ok {
		return val
	}
	return NewValue(nil)
}

// PushFrame ..
func (env *Env) PushFrame(frame *Frame) {
	env.stack = append(env.stack, frame)
}

// PopFrame ..
func (env *Env) PopFrame() *Frame {
	frame := env.stack[len(env.stack)-1]
	env.stack = env.stack[:len(env.stack)-1]
	return frame
}

//Frame ..
type Frame struct {
	FuncName string
	Pc       int
	Vars     map[string]*Value //local vars
}

func (f *Frame) String() string {
	return fmt.Sprintf("[%s:%d:%s]", f.FuncName, f.Pc, f.Vars)
}

// NewFrame ..
func NewFrame(funcName string, pc int, args []*Value) *Frame {
	return &Frame{
		FuncName: funcName,
		Pc:       pc,
		// TODO args, copy by value? ref?
	}
}
