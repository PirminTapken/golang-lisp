package lisp

import (
	"github.com/PirminFix/golang-lisp/interpreter"
)

func Eval(s string) {
	tokens := Tokenize(s)
	i := interpreter.New()
	i.Eval(tokens)
}
