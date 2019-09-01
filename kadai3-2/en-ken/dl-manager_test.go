package divdl_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"regexp"
	"strconv"
	"testing"

	"github.com/google/go-cmp/cmp"
	divdl "github.com/gopherdojo/dojo6/kadai3-2/en-ken"
)

func TestDivideIntoRanges(t *testing.T) {

	type testCase struct {
		contentLength int64
		num           int
		expected      [][]*divdl.TestRange
		expectedNum   int
	}

	cases := []testCase{
		{
			contentLength: 100,
			num:           2,
			expected: [][]*divdl.TestRange{
				{
					&divdl.TestRange{ID: 0, From: 0, To: 49},
				},
				{
					&divdl.TestRange{ID: 1, From: 50, To: 99},
				},
			},
			expectedNum: 2,
		},
		{
			contentLength: 1000,
			num:           5,
			expected: [][]*divdl.TestRange{
				{
					&divdl.TestRange{ID: 0, From: 0, To: 199},
				},
				{
					&divdl.TestRange{ID: 1, From: 200, To: 399},
				},
				{
					&divdl.TestRange{ID: 2, From: 400, To: 599},
				},
				{
					&divdl.TestRange{ID: 3, From: 600, To: 799},
				},
				{
					&divdl.TestRange{ID: 4, From: 800, To: 999},
				},
			},
			expectedNum: 5,
		},
		{
			contentLength: 1001,
			num:           5,
			expected: [][]*divdl.TestRange{
				{
					&divdl.TestRange{ID: 0, From: 0, To: 200},
				},
				{
					&divdl.TestRange{ID: 1, From: 201, To: 401},
				},
				{
					&divdl.TestRange{ID: 2, From: 402, To: 602},
				},
				{
					&divdl.TestRange{ID: 3, From: 603, To: 803},
				},
				{
					&divdl.TestRange{ID: 4, From: 804, To: 1000},
				},
			},
			expectedNum: 5,
		},
		{
			contentLength: 1005,
			num:           5,
			expected: [][]*divdl.TestRange{
				{
					&divdl.TestRange{ID: 0, From: 0, To: 200},
				},
				{
					&divdl.TestRange{ID: 1, From: 201, To: 401},
				},
				{
					&divdl.TestRange{ID: 2, From: 402, To: 602},
				},
				{
					&divdl.TestRange{ID: 3, From: 603, To: 803},
				},
				{
					&divdl.TestRange{ID: 4, From: 804, To: 1004},
				},
			},
			expectedNum: 5,
		},
		{
			contentLength: divdl.MaxRangeSize*8 - 10,
			num:           5,
			expected: [][]*divdl.TestRange{
				{
					&divdl.TestRange{ID: 0, From: 0, To: divdl.MaxRangeSize - 1},
					&divdl.TestRange{ID: 5, From: divdl.MaxRangeSize * 5, To: divdl.MaxRangeSize*6 - 1},
				},
				{
					&divdl.TestRange{ID: 1, From: divdl.MaxRangeSize, To: divdl.MaxRangeSize*2 - 1},
					&divdl.TestRange{ID: 6, From: divdl.MaxRangeSize * 6, To: divdl.MaxRangeSize*7 - 1},
				},
				{
					&divdl.TestRange{ID: 2, From: divdl.MaxRangeSize * 2, To: divdl.MaxRangeSize*3 - 1},
					&divdl.TestRange{ID: 7, From: divdl.MaxRangeSize * 7, To: divdl.MaxRangeSize*8 - 10 - 1},
				},
				{
					&divdl.TestRange{ID: 3, From: divdl.MaxRangeSize * 3, To: divdl.MaxRangeSize*4 - 1},
				},
				{
					&divdl.TestRange{ID: 4, From: divdl.MaxRangeSize * 4, To: divdl.MaxRangeSize*5 - 1},
				},
			},
			expectedNum: 8,
		},
		{
			contentLength: divdl.MaxRangeSize * 3,
			num:           2,
			expected: [][]*divdl.TestRange{
				{
					&divdl.TestRange{ID: 0, From: 0, To: divdl.MaxRangeSize - 1},
					&divdl.TestRange{ID: 2, From: divdl.MaxRangeSize * 2, To: divdl.MaxRangeSize*3 - 1},
				},
				{
					&divdl.TestRange{ID: 1, From: divdl.MaxRangeSize, To: divdl.MaxRangeSize*2 - 1},
				},
			},
			expectedNum: 3,
		},
	}

	for i, c := range cases {
		c := c
		t.Run(fmt.Sprintf("case %v", i), func(t *testing.T) {
			n, actual := divdl.DivideIntoRanges(c.contentLength, c.num)

			if !cmp.Equal(actual, c.expected) || n != c.expectedNum {
				t.Errorf("failed. Diff:\n%v", cmp.Diff(actual, c.expected))
			}
		})
	}
}

func TestDo(t *testing.T) {
	type testCase struct {
		newRequestError bool
		downloadError   bool
		canGetRange     bool
	}

	cases := []testCase{
		{
			canGetRange: true,
		},
		{
			canGetRange: false,
		},
		{
			newRequestError: true,
		},
		{
			canGetRange:   true,
			downloadError: true,
		},
		{
			canGetRange:   false,
			downloadError: true,
		},
	}

	bodyStr := fmt.Sprintf("%v%v%v%v%v",
		"0000000000",
		"1111111111",
		"2222222222",
		"3333333333",
		"4444444444",
	)

	divdl.SetMaxRangeSize(10)
	tmpDir, _ := ioutil.TempDir("", ".tmp")
	fileName := filepath.Join(tmpDir, "test.txt")

	for i, c := range cases {
		c := c

		var testHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.Method {
			case "HEAD":
				if c.newRequestError {
					w.WriteHeader(http.StatusMethodNotAllowed)
					fmt.Fprint(w, "")
					return
				}

				if c.canGetRange {
					w.Header().Add("Accept-Ranges", "bytes")
				}
				w.Header().Add("Content-Length", fmt.Sprint(len(bodyStr)))
				fmt.Print(w, "")
			case "GET":
				if c.downloadError {
					w.WriteHeader(http.StatusForbidden)
					fmt.Fprint(w, "")
					return
				}

				rngStr := r.Header.Get("Range")
				if rngStr != "" {
					reg := regexp.MustCompile(`\d+`)
					rngs := reg.FindAllString(rngStr, -1)
					from, _ := strconv.ParseInt(rngs[0], 10, 0)
					to, _ := strconv.ParseInt(rngs[1], 10, 0)
					fmt.Fprint(w, bodyStr[from:to+1])
				} else {
					fmt.Fprint(w, bodyStr)
				}
			}
		})
		ts := httptest.NewServer(testHandler)
		defer ts.Close()

		t.Run(fmt.Sprintf("case %v", i), func(t *testing.T) {
			err := divdl.Do(ts.URL, fileName, 3)
			if err != nil {
				if !c.newRequestError &&
					!c.downloadError {
					t.Errorf("Unexpected result: %v", err)
				}
			} else {
				if c.newRequestError ||
					c.downloadError {
					t.Errorf("Unexpected result: newRequestError->%v, downloadError->%v",
						c.newRequestError, c.downloadError)
				}
			}

			actual, _ := ioutil.ReadFile(fileName)
			if string(actual) != bodyStr {
				t.Errorf("Downloading file is invalid data: %v", string(actual))
			}

		})
	}
}
