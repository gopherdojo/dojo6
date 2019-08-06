package wordsreader

import (
	"reflect"
	"testing"
)

func TestWordsReader_Read(t *testing.T) {
	cases := []struct {
		fileName string
		expected []string
	}{
		{fileName: "./testdata/test1.txt", expected: []string{"a", "b", "c"}},
		{fileName: "./testdata/test2.txt", expected: []string{"a"}},
	}

	for _, c := range cases {
		c := c
		t.Run(c.fileName, func(t *testing.T) {
			t.Parallel()
			wr := WordsReader{FileName: c.fileName}
			if actual, _ := wr.Read(); !reflect.DeepEqual(c.expected, actual) {
				t.Errorf("want %v, got %v", c.expected, actual)
			}
		})
	}
}
