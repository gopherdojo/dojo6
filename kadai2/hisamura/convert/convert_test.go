package convert

import (
	"testing"
)

type testImage struct {
	name     string
	src      string
	dest     string
	expected string
}

func TestSuccess(t *testing.T) {

	cases := []testImage{
		testImage{"jpgToGif", "jpg", "gif", "../testdata/test.gif"},
		testImage{"jpgToGif", "jpg", "gif", "../testdata/test.gif"},
		testImage{"jpgToPng", "jpg", "png", "../testdata/test.png"},
		testImage{"gifToJpg", "gif", "jpg", "../testdata/test.jpg"},
		testImage{"gifToPng", "gif", "png", "../testdata/test.png"},
		testImage{"pngToGif", "png", "gif", "../testdata/test.gif"},
		testImage{"pngToJpg", "gif", "png", "../testdata/test.png"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			testSuccessConvert(t, c)
		})
	}
}

func testSuccessConvert(t *testing.T, c testImage) {
	t.Helper()
	myflag := FlagOps{"../testdata", c.src, c.dest, false}
	convertedNmae, err := Convert(myflag)

	if convertedNmae != c.expected {
		t.Fatal("differnt name")
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestFail(t *testing.T) {
	cases := []testImage{
		testImage{"jpgToGif", "jpg", "img", "../testdata/test.img"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			testFailConvert(t, c)
		})
	}

}

func testFailConvert(t *testing.T, c testImage) {
	t.Helper()
	myflag := FlagOps{"../testdata", c.src, c.dest, false}
	convertedNmae, err := Convert(myflag)

	if convertedNmae != "" {
		t.Fatal("not empty name")
	}
	if err == nil {
		t.Fatal("err nil")
	}
}
