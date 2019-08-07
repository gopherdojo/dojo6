package typing_test

import (
	"testing"

	typing "github.com/gopherdojo/dojo6/kadai3/en-ken"
)

var dict = []string{"ABC", "DEF", "GHI", "JKL", "MNO", "PQR", "STU", "VWX", "YZ"}

func contains(s string) bool {
	for _, v := range dict {
		if v == s {
			return true
		}
	}
	return false
}

func TestTypingCanGetNextText(t *testing.T) {
	typ := typing.NewTyping(dict)
	for i := 0; i < 100; i++ {
		txt := typ.GetNextText()
		if !contains(txt) {
			t.Errorf("actual: %v\n", txt)
		}
	}
}

func TestTypingIsCorrect(t *testing.T) {
	typ := typing.NewTyping(dict)
	txt := typ.GetNextText()

	if !typ.IsCorrect(txt) {
		t.Errorf("IsCorrect() must be true")
	}
}

func TestTypingIsCorrectFailed(t *testing.T) {
	typ := typing.NewTyping(dict)
	typ.GetNextText()

	if typ.IsCorrect("YZZ") {
		t.Errorf("IsCorrect() must be false")
	}
}
