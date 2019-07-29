package converter

import (
	"image"
	"image/gif"
	"os"
)

// gifConverter Gif用コンバーター
type gifConverter struct {
	converter
}

func (gc gifConverter) convert(file *os.File, img image.Image) error {
	return gif.Encode(file, img, &gif.Options{NumColors: 256})
}
