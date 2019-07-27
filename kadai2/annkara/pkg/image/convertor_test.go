package image

import (
	"os"
	"path/filepath"
	"testing"
)

var path = filepath.Join("testdata")

func TestConvertSuccess(t *testing.T) {
	in, err := os.Open(filepath.Join(path, "test.jpg"))
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer in.Close()

	out, err := os.Create(filepath.Join(path, "test.png"))
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer out.Close()

	err = Convert(in, out, "jpg")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
