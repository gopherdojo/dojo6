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
	bodyStr := fmt.Sprintf("%v%v%v%v%v",
		"0000000000",
		"1111111111",
		"2222222222",
		"3333333333",
		"4444444444",
	)

	cases := []struct {
		canGet      bool
		canGetRange bool
		from        int64
		to          int64
		rngStr      string
		expected    []byte
	}{
		{
			canGet:      true,
			canGetRange: true,
			from:        10,
			to:          19,
			rngStr:      "bytes=10-19",
			expected:    []byte(bodyStr[10:20]),
		},
		{
			canGet:      true,
			canGetRange: false,
			expected:    []byte(bodyStr),
		},
		{
			canGet: false,
		},
	}

	for i, c := range cases {
		c := c
		bodyStr := bodyStr

		var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "HEAD":
				if c.canGetRange {
					w.Header().Add("Accept-Ranges", "bytes")
				}
				w.Header().Add("Content-Length", fmt.Sprint(len(bodyStr)))
				fmt.Print(w, "")
			case "GET":
				if c.canGet {
					rng := r.Header.Get("Range")
					if rng != "" && rng == c.rngStr {
						fmt.Fprint(w, bodyStr[c.from:c.to+1])
					} else {
						fmt.Fprint(w, bodyStr)
					}
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprint(w, "")
				}
			}
		})
		ts := httptest.NewServer(testHandler)
		defer ts.Close()

		t.Run(fmt.Sprintf("case %v", i), func(t *testing.T) {
			req, _ := divdl.NewRequest(ts.URL)
			var (
				actual []byte
				err    error
			)
			if req.CanAcceptRangeRequest() {
				actual, err = req.DownloadPartially(c.from, c.to)
			} else {
				actual, err = req.Download()
			}

			if !c.canGet {
				if err == nil {
					t.Error("Unexpected http status")
				}
			}

			if !cmp.Equal(c.expected, actual) {
				t.Errorf("Unexpected response: Diff\n%v", cmp.Diff(c.expected, actual))
			}
		})
	}
}
