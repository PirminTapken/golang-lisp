package lisp

import (
	"container/list"
	"reflect"
	"testing"
)

// Test if a and b are equal. call t.Fail if not
func testEqual(a, b *list.List, t *testing.T) {
	if a.Len() != b.Len() {
		t.Fatalf("Length not equal: %d != %d\n", a.Len(), b.Len())
	}

	i := a.Front()
	j := b.Front()

	for i != nil && j != nil {
		if reflect.TypeOf(i.Value) != reflect.TypeOf(j.Value) {
			t.Fatal("Types not equal", reflect.TypeOf(i.Value), reflect.TypeOf(j.Value))
		}
		switch i.Value.(type) {
		case *list.List:
			testEqual(i.Value.(*list.List), j.Value.(*list.List), t)
		default:
			if i.Value != j.Value {
				t.Fail()
			}
		}
		i = i.Next()
		j = j.Next()
	}
}

func testNotEqual(a, b *list.List, t *testing.T) {
	if a.Len() != b.Len() {
		t.Fatalf("Length not equal: %d != %d\n", a.Len(), b.Len())
	}

	i := a.Front()
	j := b.Front()

	for i != nil && j != nil {
		if reflect.TypeOf(i.Value) != reflect.TypeOf(j.Value) {
			t.Fatal("Types not equal", reflect.TypeOf(i.Value), reflect.TypeOf(j.Value))
		}
		switch i.Value.(type) {
		case *list.List:
			testEqual(i.Value.(*list.List), j.Value.(*list.List), t)
		default:
			if i.Value != j.Value {
				t.Fail()
			}
		}
		i = i.Next()
		j = j.Next()
	}
}

func printTokenList(l *list.List, t *testing.T) {
	for e := l.Front(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case *list.List:
			printTokenList(e.Value.(*list.List), t)
		default:
			t.Log(e.Value)
		}
	}
}

// Test if the tokenizer correctly parses (a b c)
func TestTokenizerWithABC(t *testing.T) {
	example := "(a b c)"
	reference := list.New()
	reference.PushBack("a")
	reference.PushBack("b")
	reference.PushBack("c")
	result := Tokenize(example)
	testEqual(result, reference, t)
}

// Test if tokenizer parses correctly an empty list
func TestTokenizerWithEmptyList(t *testing.T) {
	example := "()"
	result := Tokenize(example)
	reference := list.New()

	testEqual(result, reference, t)
}

func TestTokenizerWithStrings(t *testing.T) {
	example := `(a "b" cde "efg")`
	result := Tokenize(example)
	reference := list.New()
	reference.PushBack("a")
	reference.PushBack(`"b"`)
	reference.PushBack(`cde`)
	reference.PushBack(`"efg"`)

	testEqual(result, reference, t)
}

func TestTokenizerWithNestedLists(t *testing.T) {
	example := `(a "b" (cde "efg"))`
	result := Tokenize(example)
	reference := list.New()
	sublist := list.New()
	sublist.PushBack(`cde`)
	sublist.PushBack(`"efg"`)

	reference.PushBack(`a`)
	reference.PushBack(`"b"`)
	reference.PushBack(sublist)

	t.Log("reference:")
	printTokenList(reference, t)
	t.Log("result:")
	printTokenList(result, t)

	testEqual(result, reference, t)
}

func TestTokenizerWithDeepNestedLists(t *testing.T) {
	example := `(a "b" ((cde hij) "efg") sausage)`
	result := Tokenize(example)
	reference := list.New()
	reference.PushBack(`a`)
	reference.PushBack(`"b"`)
	sublist := list.New()
	subsublist := list.New()
	subsublist.PushBack(`cde`)
	subsublist.PushBack(`hij`)
	sublist.PushBack(subsublist)
	sublist.PushBack(`"efg"`)
	reference.PushBack(sublist)
	reference.PushBack(`sausage`)

	t.Log("reference:")
	printTokenList(reference, t)
	t.Log("result:")
	printTokenList(result, t)

	testEqual(result, reference, t)
}
