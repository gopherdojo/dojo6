package imgconv

import (
	"strings"
)

// ImgExtension is enum type for image extension.
type ImgExtension int

const (
	UNDEF ImgExtension = iota
	JPEG
	PNG
	GIF
)

// ImgExtension => String
func (ie ImgExtension) String() string {
	switch ie {
	case JPEG:
		return "jpeg"
	case PNG:
		return "png"
	case GIF:
		return "gif"
	default:
		return "undefined"
	}
}

// String => ImgExtension
func ParseImgExtension(s string) ImgExtension {
	stringList := strings.Split(s, ".")
	ext := stringList[len(stringList)-1]
	ext = strings.ToLower(ext)
	switch ext {
	case "jpg", "jpeg":
		return JPEG
	case "png":
		return PNG
	case "gif":
		return GIF
	default:
		return UNDEF
	}
}
