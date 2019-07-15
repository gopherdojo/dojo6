package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo6/kadai1/micchie/imgconvert"
	"github.com/pkg/errors"
)

var (
	dir          string
	beforeFormat string
	afterFormat  string
)

func init() {
	flag.StringVar(&dir, "d", "", "path of images to convert [required]")
	flag.StringVar(&beforeFormat, "before", "jpg", "image format before conversion")
	flag.StringVar(&afterFormat, "after", "png", "image format after conversion")
	flag.Parse()
}

func main() {
	if err := exec(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
}

func exec() error {
	d := strings.TrimSpace(dir)
	if d == "" {
		return errors.New("directory is required")
	}

	info, err := os.Stat(d)
	if os.IsNotExist(err) {
		return errors.Wrap(err, "directory does not exist")
	}

	if !info.IsDir() {
		return errors.New("set a directory")
	}

	var bf, af string
	bf, err = validationFormat(beforeFormat)
	if err != nil {
		return errors.Wrapf(err, "format: %v", beforeFormat)
	}

	af, err = validationFormat(afterFormat)
	if err != nil {
		return errors.Wrapf(err, "format: %v", afterFormat)
	}

	return filepath.Walk(d, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		img := &imgconvert.ConvertImage{}
		img.Before = imgconvert.ImageFile(p)

		if img.Before.Ext() != fmt.Sprintf(".%v", bf) {
			return nil
		}

		dir, err := img.Before.Directory()
		if err != nil {
			return nil
		}

		img.After = imgconvert.ImageFile(fmt.Sprintf("%v/%v.%v", dir, img.Before.Name(), af))

		return img.Convert(af)
	})
}

func validationFormat(format string) (string, error) {
	f := strings.ToLower(format)
	switch {
	case f == "jpg" || f == "png":
		return f, nil
	}
	return "", errors.New("format is not valid")
}
