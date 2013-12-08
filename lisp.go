package lisp

import (
	"container/list"
)

func Tokenize(s string) *list.List {
	l := new(list.List)
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")
	return l
}
