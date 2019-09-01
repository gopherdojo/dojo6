package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

// Request wraps downloading proceess.
type Request struct {
	url                   string
	contentLength         int64
	canAcceptRangeRequest bool
}

// NewRequest is a constructor of Request.
func NewRequest(url string) (*Request, error) {
	resp, err := http.Head(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	// Header request before downloading
	acceptRanges := resp.Header.Get("Accept-Ranges")
	contentLength, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 0)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Request{
		url:                   url,
		contentLength:         contentLength,
		canAcceptRangeRequest: acceptRanges == "bytes",
	}, nil
}

// DownloadPartially downloads specified part of data from the specified url.
func (r *Request) DownloadPartially(from int64, to int64) ([]byte, error) {
	if !r.canAcceptRangeRequest {
		return nil, fmt.Errorf("This file cannot download with Range")
	}

	// Set Range Header if Range request is accepted.
	headers := map[string]string{}
	headers["Range"] = fmt.Sprintf("bytes=%d-%d", from, to)

	return r.download(headers)
}

// Download downloads data from the specified url.
func (r *Request) Download() ([]byte, error) {
	return r.download(nil)
}

func (r *Request) download(extraHeaders map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", r.url, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for k, v := range extraHeaders {
		req.Header.Add(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if resp.Status[:1] != "2" {
		return nil, errors.Errorf("Faild to request: %v", resp.StatusCode)
	}
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll((resp.Body))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return body, nil
}

// CanAcceptRangeRequest is getter of canAcceptRengeRequest
func (r *Request) CanAcceptRangeRequest() bool {
	return r.canAcceptRangeRequest
}

// GetContentLength is getter to content length of data.
func (r *Request) GetContentLength() int64 {
	return r.contentLength
}
