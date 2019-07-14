package image

import (
	"os"
	"testing"
)

var path = "../../test/"

func TestConvertSuccess(t *testing.T) {
	in, err := os.Open(path + "test.jpg")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer in.Close()

	out, err := os.Create(path + "test.png")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	defer out.Close()

	err = Convert(in, out, "jpg")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}
