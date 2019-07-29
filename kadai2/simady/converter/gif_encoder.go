package converter

import (
	"image"
	"image/gif"
	"os"
)

// gifEncoder Gif用エンコーダー
type gifEncoder struct {
}

func (*gifEncoder) encode(file *os.File, img image.Image) error {
	return gif.Encode(file, img, &gif.Options{NumColors: 256})
}
