package lisp

import (
	"container/list"
	"testing"
)

// Test if a and b are equal. call t.Fail if not
func testEqual(a, b *list.List, t *testing.T) {
	if a.Len() != b.Len() {
		t.Fail()
	}

	i := a.Front()
	j := b.Front()

	for i != nil && j != nil {
		if i.Value != j.Value {
			t.Fail()
		}
		i = i.Next()
		j = j.Next()
	}
}

func printTokenList(l *list.List, t *testing.T) {
	for e := l.Front(); e != nil; e = e.Next() {
		t.Log(e)
	}
}

// Test if the tokenizer correctly parses (a b c)
func TestTokenizerWithABC(t *testing.T) {
	example := "(a b c)"
	result := Tokenize(example)
	reference := list.New()
	reference.PushBack("a")
	reference.PushBack("b")
	reference.PushBack("c")

	testEqual(result, reference, t)
}

// Test if tokenizer parses correctly an empty list
func TestTokenizerWithEmptyList(t *testing.T) {
	example := "()"
	result := Tokenize(example)
	reference := list.New()

	testEqual(result, reference, t)
}
