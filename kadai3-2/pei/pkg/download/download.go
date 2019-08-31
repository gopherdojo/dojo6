package download

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

// Downloader Interface
type Downloader interface {
	Do() (string, error)
}

// NonRangeDownloader has info for downloading
type NonRangeDownloader struct {
	url, outputPath string
}

// RangeDownloader has info for split downloading
type RangeDownloader struct {
	splitNum        int
	ranges          []*Range
	url, outputPath string
}

// Range has rangestart and rangeend
type Range struct {
	start int64
	end   int64
}

// NewDownloader creates Downloader
func NewDownloader(splitNum int, url, outputPath string) (Downloader, error) {
	dir, fileName := parseDirAndFileName(outputPath)
	if fileName == "" {
		outputPath = filepath.Join(dir, parseFileName(url))
	}

	res, err := http.Head(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.Header.Get("Accept-Ranges") != "bytes" {
		return &NonRangeDownloader{url: url, outputPath: outputPath}, nil
	}

	contentLength := res.ContentLength
	unit := contentLength / int64(splitNum)
	ranges := make([]*Range, splitNum)

	for i := range ranges {
		var start, end int64
		if i != 0 {
			start = int64(i)*unit + 1
		}
		end = int64(i+1) * unit
		if i == splitNum-1 {
			end = contentLength
		}

		ranges[i] = &Range{start: start, end: end}
	}

	return &RangeDownloader{
		splitNum:   splitNum,
		ranges:     ranges,
		url:        url,
		outputPath: outputPath,
	}, nil
}

// Do download
func (d *NonRangeDownloader) Do() (string, error) {
	req, err := http.NewRequest(http.MethodGet, d.url, nil)
	if err != nil {
		return "", err
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	return d.outputPath, saveResponseBody(d.outputPath, res)
}

// Do split download
func (d *RangeDownloader) Do() (string, error) {
	eg, ctx := errgroup.WithContext(context.TODO())

	for i := range d.ranges {
		i := i
		eg.Go(func() error {
			return d.do(ctx, i)
		})
	}

	if err := eg.Wait(); err != nil {
		return "", err
	}

	return d.outputPath, d.mergeFiles()
}

func (d *RangeDownloader) do(ctx context.Context, idx int) error {
	req, err := http.NewRequest(http.MethodGet, d.url, nil)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)

	ran := d.ranges[idx]
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", ran.start, ran.end))

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	tmpFileName := fmt.Sprintf("%s.%d", d.outputPath, idx)
	return saveResponseBody(tmpFileName, res)
}

func (d *RangeDownloader) mergeFiles() error {
	file, err := os.Create(d.outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	for i := range d.ranges {
		tmpFileName := fmt.Sprintf("%s.%d", d.outputPath, i)
		tmpFile, err := os.Open(tmpFileName)
		if err != nil {
			return err
		}

		io.Copy(file, tmpFile)
		tmpFile.Close()
		if err := os.Remove(tmpFileName); err != nil {
			return err
		}
	}

	return nil
}

func saveResponseBody(fileName string, response *http.Response) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(file, response.Body); err != nil {
		return err
	}

	return nil
}
