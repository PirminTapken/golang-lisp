package lisp

import (
	"container/list"
	"github.com/PirminTapken/golang-lisp"
)

// Interpreter is a interpreter for lisp
// LoadBasic functions should be set to true if
// you want the basic functions being loaded
type Interpreter struct {
	functions          map[string]func(*list.List) *list.List
	LoadBasicFunctions bool
}

func New() *Interpreter {
	i := new(Interpreter)
	i.LoadBasicFunctions = true
	return i
}

func (i *Interpreter) Register(name string, f func(*list.List) *list.List) {
	i.functions[name] = f
}

func (i *Interpreter) Unregister(name string) {
	delete(i.functions, name)
}

func (i *Interpreter) loadBasicOperators() {
	operatorList := map[string]func(*list.List) *list.List{
		"quote": opQuote,
	}
	for k, v := range operatorList {
		i.Register(k, v)
	}
}

func (i *Interpreter) Eval(t *list.List) {
	if i.LoadBasicFunctions {
	}
}
