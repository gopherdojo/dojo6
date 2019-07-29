package converter

import (
	"image"
	"image/png"
	"os"
)

// pngEncoder Png用エンコーダー
type pngEncoder struct {
}

func (*pngEncoder) encode(file *os.File, img image.Image) error {
	return png.Encode(file, img)
}
