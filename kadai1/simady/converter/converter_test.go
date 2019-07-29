package converter

import (
	"testing"
)

func TestGetConverter(t *testing.T) {
	var c IConverter
	c = GetConverter("jpeg", "png")
	if _, ok := c.(pngConverter); !ok {
		t.Fatal("c is not pngConverter.")
	}

	c = GetConverter("png", "gif")
	if _, ok := c.(gifConverter); !ok {
		t.Fatal("c is not gifConverter.")
	}

	c = GetConverter("gif", "jpeg")
	if _, ok := c.(jpegConverter); !ok {
		t.Fatal("c is not jpegConverter.")
	}
}

func TestValidate(t *testing.T) {
	var src, dst string

	// 正常系
	src, dst = "jpeg", "png"
	if !(converter{src: src, dst: dst}).validate() {
		t.Fatalf("Expect src=%v dst=%v to be supported.\n", src, dst)
	}

	src, dst = "png", "gif"
	if !(converter{src: src, dst: dst}).validate() {
		t.Fatalf("Expect src=%v dst=%v to be supported.\n", src, dst)
	}

	src, dst = "gif", "jpeg"
	if !(converter{src: src, dst: dst}).validate() {
		t.Fatalf("Expect src=%v dst=%v to be supported.\n", src, dst)
	}

	// 異常系
	src, dst = "png", "webp"
	if (converter{src: src, dst: dst}).validate() {
		t.Fatalf("src=%v dst=%v not supported.\n", src, dst)
	}
}

func TestConvert(t *testing.T) {
	var path, fileName string
	path, fileName = "../src/gopher.jpeg", "../dst/test/test.png"
	if err := Convert(path, pngConverter{}, fileName); err != nil {
		t.Fatal("Failed to convert.")
	}

	path, fileName = "../src/gopher.png", "../dst/test/test.gif"
	if err := Convert(path, gifConverter{}, fileName); err != nil {
		t.Fatal("Failed to convert.")
	}

	path, fileName = "../src/gopher.gif", "../dst/test/test.jpeg"
	if err := Convert(path, jpegConverter{}, fileName); err != nil {
		t.Fatal("Failed to convert.")
	}
}
