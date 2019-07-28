package dir

import "testing"

var directory = "testdata"

type testcase struct {
	desc   string
	before string
	after  string
	err    bool
}

func TestWalk(t *testing.T) {

	tests := []testcase{
		{
			desc:   "Succsess",
			before: "jpg",
			after:  "png",
			err:    false,
		}, {
			desc:   "Fail",
			before: "jpg",
			after:  "gif",
			err:    true,
		},
	}

	for _, tc := range tests {
		testWalk(t, tc)
	}
}

func testWalk(t *testing.T, tc testcase) {

	err := Walk(directory, tc.before, tc.after)
	if !(err != nil) == tc.err {
		t.Errorf("failed test %s: %#v", tc.desc, err)
	}
}
