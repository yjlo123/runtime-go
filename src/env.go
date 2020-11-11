package main

// Env ..
type Env struct {
	Labels map[string]int
	Vars   map[string]*Value
	Funcs  map[string]int
}

// AssignVar ..
func (env Env) AssignVar(varName string, val *Value) {
	env.Vars[varName] = val
}

// GetVarVal ..
func (env Env) GetVarVal(varName string) *Value {
	if val, ok := env.Vars[varName]; ok {
		return val
	}
	return nil
}
