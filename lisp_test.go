package lisp

import (
	"container/list"
	"testing"
)

// Test if string to token works correctly
func TestTokenizer(t *testing.T) {
	example := "(a b c)"
	result := Tokenize(example)
	reference := list.New()
	reference.PushBack("a")
	reference.PushBack("b")
	reference.PushBack("c")

	if result.Len() != reference.Len() {
		t.Fail()
	}
}
