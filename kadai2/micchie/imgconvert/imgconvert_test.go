package imgconvert

import "testing"

func TestImageFile_Ext(t *testing.T) {
	tests := []struct {
		file string
		ext  string
	}{
		{file: "../testdata/test1.png", ext: ".png"},
		{file: "../testdata/test2.png", ext: ".png"},
		{file: "../testdata/test_directory/test3.jpg", ext: ".jpg"},
	}

	for _, test := range tests {
		ext := ImageFile(test.file).Ext()
		if ext != test.ext {
			t.Errorf("file: %#v (want: %v, got: %v)", test.file, test.ext, ext)
			continue
		}
	}
}

func TestImageFile_Name(t *testing.T) {
	tests := []struct {
		file string
		name string
	}{
		{file: "../testdata/test1.png", name: "test1"},
		{file: "../testdata/test2.png", name: "test2"},
		{file: "../testdata/test_directory/test3.jpg", name: "test3"},
	}

	for _, test := range tests {
		name := ImageFile(test.file).Name()
		if name != test.name {
			t.Errorf("file: %#v (want: %v, got: %v)", test.file, test.name, name)
			continue
		}
	}
}

func TestConvertImage_Convert(t *testing.T) {
	tests := []struct {
		format   string
		before   string
		after    string
		hasError bool
	}{
		{format: "jpg", before: "../testdata/test1.png", after: "../testdata/test1.jpg", hasError: false},
		{format: "png", before: "../testdata/test2.jpg", after: "../testdata/test2.png", hasError: false},
		{format: "png", before: "../testdata/test_directory/test3.jpg", after: "../testdata/test_directory/test3.png", hasError: false},
		{format: "pdf", before: "../testdata/test_directory/test4.pdf", after: "../testdata/test_directory/test4.pdf", hasError: true},
		{format: "png", before: "../testdata/test/test5.png", after: "../testdata/test/test5.png", hasError: true},
	}

	for _, test := range tests {
		images := &ConvertImage{
			Before: ImageFile(test.before),
			After:  ImageFile(test.after),
		}

		var hasError bool

		if err := images.Convert(test.format); err != nil {
			hasError = true
		}

		if hasError != test.hasError {
			t.Errorf("images: %#v (want: %v, got: %v)", images, test.hasError, hasError)
			continue
		}
	}
}
