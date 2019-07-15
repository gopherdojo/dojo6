package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo6/kadai1/ikkun/imgconv"
)

func main() {
	from := flag.String("f", "jpg", "from format(jpg|jpeg|png|gif) default jpg.")
	to := flag.String("t", "png", "to format(jpg|jpeg|png|gif) default png.")

	flag.Parse()

	if !validateFormat(*from) {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("unsupport format: %s", *from))
		os.Exit(1)
	}

	if !validateFormat(*to) {
		fmt.Fprintln(os.Stderr, fmt.Sprintf("unsupport format: %s", *to))
		os.Exit(1)
	}

	args := flag.Args()
	if len(args) < 1 {
		fmt.Fprintln(os.Stderr, "required directory assigned")
		os.Exit(1)
	}

	dir := args[0]
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for _, path := range paths {
		img := imgconv.ImageFile{path, *from, *to}
		err = imgconv.Convert(img)
		if err != nil {
			fmt.Fprintf(os.Stdout, "%s: %s\n", err, img.Source)
		}
	}
}

func validateFormat(f string) bool {
	switch f {
	case "jpg", "jpeg", "png", "gif":
		return true
	}

	return false
}
