package imgconv

import "testing"

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
