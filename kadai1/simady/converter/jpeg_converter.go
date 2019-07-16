package converter

import (
	"image"
	"image/jpeg"
	"os"
)

// jpegConverter Jpeg用コンバーター
type jpegConverter struct {
	converter
}

func (jc jpegConverter) convert(file *os.File, img image.Image) error {
	return jpeg.Encode(file, img, &jpeg.Options{Quality: 100})
}
