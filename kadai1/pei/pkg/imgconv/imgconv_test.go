package imgconv

import (
	"os"
	"testing"
)

func TestImgConverter_Convert(t *testing.T) {
	const (
		jpgfile = "../../testdata/sample.jpg"
		pngfile = "../../testdata/sample.png"
		giffile = "../../testdata/sample.gif"

		errorMessage = "Unexpected behavior"
	)
	defer func() {
		os.Remove(pngfile)
		os.Remove(giffile)
	}()

	var (
		ic  ImgConverter
		err error
	)

	// jpg -> png
	ic = ImgConverter{
		InputPath:       jpgfile,
		InputExtension:  JPEG,
		OutputPath:      pngfile,
		OutputExtension: PNG,
		LeaveInput:      true,
	}
	err = ic.Convert()
	if err != nil {
		t.Errorf(errorMessage)
	}

	// png -> gif
	ic = ImgConverter{
		InputPath:       pngfile,
		InputExtension:  PNG,
		OutputPath:      giffile,
		OutputExtension: GIF,
		LeaveInput:      false,
	}
	err = ic.Convert()
	if err != nil {
		t.Errorf(errorMessage)
	}
}
