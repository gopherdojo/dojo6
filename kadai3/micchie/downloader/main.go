package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo6/kadai3/micchie/downloader/download"
	"github.com/pkg/errors"
)

var downloadURL string

func init() {
	flag.StringVar(&downloadURL, "url", "", "download [required]")
	flag.Parse()
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	url := strings.TrimSpace(downloadURL)
	if url == "" {
		return errors.New("download url is required")
	}

	// TODO: URL の整合性

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	d := download.NewDownload(ctx, url)

	f, err := os.Create(filepath.Base(url))
	if err != nil {
		return errors.Wrap(err, "failed to write")
	}
	defer f.Close()

	err := d.Get(f)
	if err != nil {
		return errors.Wrapf(err, "failed to get content: %v", url)
	}

	return nil
}
