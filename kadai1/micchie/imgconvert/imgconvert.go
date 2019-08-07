package imgconvert

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// ImageFile is the Image file.
type ImageFile string

// Directory is the File directory.
func (i ImageFile) Directory() (string, error) {
	p, err := filepath.Abs(filepath.Clean(string(i)))
	return filepath.Dir(p), err
}

// Ext is the File extention.
func (i ImageFile) Ext() string {
	return strings.ToLower(filepath.Ext(string(i)))
}

// Name is the File name without extension.
func (i ImageFile) Name() string {
	file := string(i)
	return filepath.Base(file[:len(file)-len(i.Ext())])
}

// ConvertImage configures convert images.
type ConvertImage struct {
	Before ImageFile
	After  ImageFile
}

// Convert writes the Image to format.
func (cnv ConvertImage) Convert(format string) error {
	bp, err := os.Open(string(cnv.Before))
	if err != nil {
		return err
	}
	defer bp.Close()

	ap, err := os.Create(string(cnv.After))
	if err != nil {
		return err
	}
	defer ap.Close()

	img, _, err := image.Decode(bp)
	if err != nil {
		return err
	}

	switch format {
	case "png":
		err = png.Encode(ap, img)
	case "jpg":
		opts := &jpeg.Options{Quality: jpeg.DefaultQuality}
		err = jpeg.Encode(ap, img, opts)
	}

	if err != nil {
		return err
	}
	return nil
}
