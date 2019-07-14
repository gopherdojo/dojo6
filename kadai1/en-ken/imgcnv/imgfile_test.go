package imgcnv

import (
	"os"
	"testing"
)

const outputDir = "../out"

func TestCanConstructImageFileFromJpegToPng(t *testing.T) {
	os.RemoveAll(outputDir)
	os.MkdirAll(outputDir, 0777)
	imageFile := NewImageFile("../testdata/lenna_color.jpg")
	err := imageFile.SaveAs(outputDir + "/lenna_color.png")
	if err != nil {
		t.Error(err)
	}
}
