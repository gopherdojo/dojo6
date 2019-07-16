package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"
)

type ImageFile struct {
	Source       string
	BeforeFormat string
	AfterFormat  string
}

func (i *ImageFile) addOrReplaceExtension() string {
	oldExt := "." + i.BeforeFormat
	newExt := "." + i.AfterFormat
	if strings.HasSuffix(i.Source, oldExt) {
		return strings.TrimSuffix(i.Source, oldExt) + newExt
	}

	return i.Source + newExt
}

func decode(r io.Reader, format string) (image.Image, error) {
	switch format {
	case "jpg", "jpeg":
		return jpeg.Decode(r)
	case "png":
		return png.Decode(r)
	case "gif":
		return gif.Decode(r)
	default:
		return nil, fmt.Errorf("unsupport format: %s", format)
	}
}

func encode(w io.Writer, img image.Image, format string) error {
	switch format {
	case "jpg", "jpeg":
		return jpeg.Encode(w, img, nil)
	case "png":
		return png.Encode(w, img)
	case "gif":
		return gif.Encode(w, img, nil)
	default:
		return fmt.Errorf("unsupport format: %s", format)
	}
}

func Convert(m ImageFile) error {
	from, err := os.Open(m.Source)
	if err != nil {
		return err
	}
	defer from.Close()

	img, err := decode(from, m.BeforeFormat)
	if err != nil {
		return err
	}

	to, err := os.Create(m.addOrReplaceExtension())
	if err != nil {
		return err
	}
	defer to.Close()

	err = encode(to, img, m.AfterFormat)
	if err != nil {
		return err
	}

	return nil
}
