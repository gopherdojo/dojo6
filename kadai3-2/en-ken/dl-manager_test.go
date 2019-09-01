package divdl_test

import (
	"fmt"
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
					&divdl.TestRange{ID: 6, From: divdl.MaxRangeSize * 6, To: divdl.MaxRangeSize * 7},
				},
				{
					&divdl.TestRange{ID: 2, From: divdl.MaxRangeSize * 2, To: divdl.MaxRangeSize*3 - 1},
					&divdl.TestRange{ID: 7, From: divdl.MaxRangeSize * 7, To: divdl.MaxRangeSize*8 - 11},
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
