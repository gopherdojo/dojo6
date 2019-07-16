package converter

import (
	"image"
	"image/png"
	"os"
)

// pngConverter Png用コンバーター
type pngConverter struct {
	converter
}

func (pc pngConverter) convert(file *os.File, img image.Image) error {
	return png.Encode(file, img)
}
