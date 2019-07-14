package imgcnv

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// ImageFile expresses the i
type ImageFile struct {
	image *image.Image
	path  string
}

// NewImageFile is a constructor of ImageFile
func NewImageFile(path string) *ImageFile {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var (
		image image.Image
	)
	switch filepath.Ext(path) {
	case ".jpg", ".jpeg":
		image, err = jpeg.Decode(file)
	case ".png":
		image, err = png.Decode(file)
	default:
		return nil
	}

	return &ImageFile{
		image: &image,
		path:  path,
	}
}

// AbsPath returns the absolute path of the input file
func (img *ImageFile) AbsPath() string {
	return img.path
}

// SaveAs oututs a file to the specified path after convering to the specified exteinsion.
func (img *ImageFile) SaveAs(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return nil
	}
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".jpeg":
		return jpeg.Encode(file, *img.image, nil)
	case ".png":
		return png.Encode(file, *img.image)
	default:
		return nil
	}
}
