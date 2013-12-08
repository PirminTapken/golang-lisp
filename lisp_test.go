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
