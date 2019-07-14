package convert

import (
	"os"
	"testing"
)


// kadai1/test_images/test.pngファイルが存在しています。
func TestSuccess(t *testing.T) {
	myflag := FlagOps{"../test_images", "png", "jpg", false}

	err := Convert(myflag)
	defer os.Remove("../test_images/test.jpg")

	if err != nil {
		t.Fatal("failed test")
	}

	if _, err := os.Stat("../test_images/test.jpg"); os.IsNotExist(err) {
		t.Fatal("failed test")
	}
}

func TestFail(t *testing.T) {
	myflag := FlagOps{"../test_images", "png", "img", false}

	err := Convert(myflag)
	defer os.Remove("../test_images/test.img")

	if err == nil {
		t.Fatal("failed test")
	}
}