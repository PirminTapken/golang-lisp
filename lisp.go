package lisp

import (
	"container/list"
	"strings"
	"text/scanner"
)

func Tokenize(s string) *list.List {
	r := strings.NewReader(s)

	var (
		sc scanner.Scanner
		ts string
		l  *list.List
	)

	stack := list.New() // Deepness Queue
	l = list.New()
	stack.PushBack(l)
	sc.Init(r)
	for tok := sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
		ts = sc.TokenText()
		switch ts {
		case "(":
			l.PushBack(list.New())
			l = l.Back().Value.(*list.List)
			stack.PushBack(l)
		case ")":
			stack.Remove(stack.Back())
			l = stack.Back().Value.(*list.List)
		default:
			l.PushBack(ts)
		}
	}
	return stack.Front().Value.(*list.List).Front().Value.(*list.List)
}
