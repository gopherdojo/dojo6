package imgconv

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"testing"
)

func TestImgconv_addReplaceExtension(t *testing.T) {
	i := ImageFile{"test.jpg", "jpg", "png"}
	r := i.addOrReplaceExtension()
	if r != "test.png" {
		t.Fatalf("failed replace extension: actual %s", r)
	}

	i = ImageFile{"test", "jpg", "png"}
	r = i.addOrReplaceExtension()
	if r != "test.png" {
		t.Fatalf("failed add extension: actual %s", r)
	}
}

func decodeFile(filename string, format string) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return decode(f, format)
}

func TestImgconv_Decode_Success(t *testing.T) {
	testCases := map[string][]string{
		"testdata/1px.png": {"png"},
		"testdata/1px.gif": {"gif"},
		"testdata/1px.jpg": {"jpg", "jpeg"},
	}

	for file, formats := range testCases {
		for _, format := range formats {
			_, err := decodeFile("testdata/1px.png", "png")
			if err != nil {
				t.Fatalf("failed decode %s format, %s file assigned", format, file)
			}
		}
	}
}

func TestImgconv_Decode_Failure(t *testing.T) {
	buf := new(bytes.Buffer)
	_, err := decode(buf, "test")
	if err == nil {
		t.Fatal("failed format handling")
	}

	testCases := map[string][]string{
		"testdata/1px.png": {"jpg", "jpeg", "gif"},
		"testdata/1px.jpg": {"gif", "png"},
		"testdata/1px.gif": {"jpg", "jpeg", "png"},
	}

	for file, formats := range testCases {
		for _, format := range formats {
			_, err := decodeFile(file, format)
			if err == nil {
				t.Fatalf("failed %s format handling when %s file assigned", format, file)
			}
		}
	}
}

func readImage(filename string, fn func(r io.Reader) (image.Image, error)) (image.Image, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return fn(f)
}

func TestImgconv_Encode_Success(t *testing.T) {
	buf := new(bytes.Buffer)

	formats := []string{"jpg", "jpeg", "gif", "png"}
	testCases := map[string]func(r io.Reader) (image.Image, error){
		"testdata/1px.png": png.Decode,
		"testdata/1px.jpg": jpeg.Decode,
		"testdata/1px.gif": gif.Decode,
	}

	for file, decoder := range testCases {
		img, err := readImage(file, decoder)
		if err != nil {
			t.Fatalf("failed read image: %v", err)
		}

		for _, format := range formats {
			err = encode(buf, img, format)
			if err != nil {
				t.Fatalf("failed decode %s format, %s file assigned", format, file)
			}
		}
	}
}

func TestImgconv_Encode_Failure(t *testing.T) {
	// FIXME: 失敗することを確認したいので、image.Imageを満たすモックを作る方が効率がいい
	img, err := readImage("testdata/1px.png", png.Decode)
	if err != nil {
		t.Fatalf("failed read image: %v", err)
	}

	buf := new(bytes.Buffer)
	err = encode(buf, img, "test")
	if err == nil {
		t.Fatal("failed format handling")
	}
}
