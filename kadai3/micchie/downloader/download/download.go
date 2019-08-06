package download

import (
	"context"
	"net/http"
)

type Downloader struct {
	URL     string
	Context context.Context
	Client  *http.Client
}

func NewDownload(ctx context.Context, url string) *Downloader {
	return &Download{
		URL:     url,
		Context: ctx,
		Client:  &http.Client{},
	}
}

func (d *Downloader) Get(ctx context, file string) error {

	return nil
}
