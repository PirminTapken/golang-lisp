package lisp

import (
	"container/list"
	"strings"
	"text/scanner"
)

func Tokenize(s string) *list.List {
	l := new(list.List)

	r := strings.NewReader(s)

	var (
		sc scanner.Scanner
		ts string
	)

	sc.Init(r)
	for tok := sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
		ts = sc.TokenText()
		if ts != "(" && ts != ")" {
			l.PushBack(ts)
		}
	}

	return l
}
