package searchfile_test

import (
	"dojo6/kadai1/uetsu/ImgConv/searchfile"
	"sort"
	"testing"
)

var TestCase = []struct {
	in  string
	out []string
}{
	{"sample", []string{"sample/hoge.jpg"}},
	{"sample2", []string{"sample2/hoge.jpg", "sample2/hogo/hoge2.jpg"}},
}

func Test_RecursionFile(t *testing.T) {
	for _, tt := range TestCase {
		expect := tt.out
		actual := searchfile.RecursionFile(tt.in)
		sort.Strings(expect)
		sort.Strings(actual)
		if len(actual) != len(expect) {
			t.Errorf(`expect="%s" actual="%s"`, expect, actual)
			break
		}
		for i := range expect {
			if actual[i] != expect[i] {
				t.Errorf(`expect="%s" actual="%s"`, expect, actual)
				break
			}
		}
	}
}
func Test_FileNotExists(t *testing.T) {
	t.Helper()
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("Unexpected File Exists")
		}
	}()
	searchfile.RecursionFile("test")
}
