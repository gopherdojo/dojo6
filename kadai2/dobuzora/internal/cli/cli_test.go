package cli_test

import (
	"testing"

	"github.com/gopherdojo/dojo6/kadai2/dobuzora/internal/cli"
)

var inputtests = []struct {
	in       string
	expected string
}{
	{"ninja.png", "ninja"},
	{"ninja.jpeg", "ninja"},
	{"n.jpg", "n"},
}

func TestGetFileNameWithoutExt(t *testing.T) {
	for _, c := range inputtests {
		actual := cli.GetFileNameWithoutExt(c.in)
		if actual != c.expected {
			t.Errorf("getFileNameWithoutExt(%q) == %q, expect %q", c.in, actual, c.expected)
		}
	}
}
