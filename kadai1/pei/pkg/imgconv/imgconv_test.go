package imgconv_test

import (
	"os"
	"testing"

	"github.com/gopherdojo/dojo6/kadai1/pei/pkg/imgconv"
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
		ic  imgconv.ImgConverter
		err error
	)

	// jpg -> png
	ic = imgconv.ImgConverter{
		InputPath:       jpgfile,
		InputExtension:  imgconv.JPEG,
		OutputPath:      pngfile,
		OutputExtension: imgconv.PNG,
		LeaveInput:      true,
	}
	err = ic.Convert()
	if err != nil {
		t.Errorf(errorMessage)
	}

	// png -> gif
	ic = imgconv.ImgConverter{
		InputPath:       pngfile,
		InputExtension:  imgconv.PNG,
		OutputPath:      giffile,
		OutputExtension: imgconv.GIF,
		LeaveInput:      false,
	}
	err = ic.Convert()
	if err != nil {
		t.Errorf(errorMessage)
	}
}
