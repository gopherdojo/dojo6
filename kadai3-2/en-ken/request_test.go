package divdl_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	divdl "github.com/gopherdojo/dojo6/kadai3-2/en-ken"

	"github.com/google/go-cmp/cmp"
)

func TestNewRequest(t *testing.T) {
	cases := []struct {
		canHead        bool
		canAcceptRange bool
	}{
		{
			canHead:        true,
			canAcceptRange: true,
		},
		{
			canHead:        true,
			canAcceptRange: false,
		},
		{
			canHead:        false,
			canAcceptRange: false,
		},
	}

	for _, c := range cases {
		c := c

		body := []byte("success")
		var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !c.canHead {
				w.WriteHeader(http.StatusMethodNotAllowed)
				return
			}

			if c.canAcceptRange {
				w.Header().Add("Accept-Ranges", "bytes")
			}
			w.Header().Add("Content-Length", fmt.Sprint(len(body)))
			fmt.Fprint(w, "")
		})
		ts := httptest.NewServer(testHandler)
		defer ts.Close()

		req, err := divdl.NewRequest(ts.URL)
		if c.canHead {
			if err != nil {
				t.Errorf("Unexpected  result: %v", err)
			}
		} else {
			if err == nil {
				t.Error("Http status error epected.")
			}
			return
		}

		if req == nil {
			t.Errorf("Unexpected result")
		}
		if req.CanAcceptRangeRequest() != c.canAcceptRange {
			t.Errorf("Unexpected result")
		}
	}
}

func TestDownload(t *testing.T) {
	cases := []struct {
		canGet bool
	}{
		{
			canGet: true,
		},
		{
			canGet: false,
		},
	}

	for _, c := range cases {
		c := c

		bodyStr := fmt.Sprintf("%v%v%v%v%v",
			"0000000000",
			"1111111111",
			"2222222222",
			"3333333333",
			"4444444444",
		)
		var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "HEAD" {
				w.Header().Add("Content-Length", fmt.Sprint(len(bodyStr)))
				fmt.Print(w, "")
				return
			}

			if c.canGet {
				fmt.Fprintf(w, "%v", bodyStr)
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Print(w, "")
			}
		})
		ts := httptest.NewServer(testHandler)
		defer ts.Close()

		req, _ := divdl.NewRequest(ts.URL)
		data, err := req.Download()
		if !c.canGet {
			if err == nil {
				t.Error("Unexpected http status")
			}
			return
		}

		if !cmp.Equal([]byte(bodyStr), data) {
			t.Errorf("Unexpected response: Diff\n%v", cmp.Diff(bodyStr, data))
		}
	}
}

func TestDownloadPartially(t *testing.T) {
	cases := []struct {
		canGet   bool
		from     int64
		to       int64
		rngStr   string
		expected string
	}{
		{
			canGet:   true,
			from:     0,
			to:       10,
			rngStr:   "bytes=0-10",
			expected: "0000000000",
		},
		{
			canGet: false,
		},
	}

	for _, c := range cases {
		c := c

		bodyStr := fmt.Sprintf("%v%v%v%v%v",
			"0000000000",
			"1111111111",
			"2222222222",
			"3333333333",
			"4444444444",
		)
		var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "HEAD" {
				w.Header().Add("Accept-Ranges", "bytes")
				w.Header().Add("Content-Length", fmt.Sprint(len(bodyStr)))
				fmt.Print(w, "")
			}
			if c.canGet {
				if c.rngStr == r.Header.Get("Range") {
					fmt.Fprint(w, "0000000000")
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprint(w, "")
			}

		})
		ts := httptest.NewServer(testHandler)
		defer ts.Close()

		req, _ := divdl.NewRequest(ts.URL)
		actual, err := req.DownloadPartially(c.from, c.to)
		if !c.canGet {
			if err == nil {
				t.Error("Unexpected http status")
			}
		}

		if !cmp.Equal([]byte(c.expected), actual) {
			t.Errorf("Unexpected response: Diff\n%v", cmp.Diff(c.expected, actual))
		}
	}
}
