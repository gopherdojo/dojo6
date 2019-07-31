package cli

import "testing"

var inputtests = []struct {
	in       string
	expected string
}{
	{"ninja.png", "ninja"},
	{"ninja.jpeg", "ninja"},
	{".jpg", ""},
}

func TestgetFileNameWithoutExt(t *testing.T) {
	for _, c := range inputtests {
		actual := getFileNameWithoutExt(c.in)
		if actual != c.expected {
			t.Errorf("getFileNameWithoutExt(%q) == %q, expect %q", c.in, actual, c.expected)
		}
	}
}
