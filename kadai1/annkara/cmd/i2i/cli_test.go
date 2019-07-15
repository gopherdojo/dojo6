package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		desc           string
		arg            string
		expectedStatus int
	}{
		{
			desc:           "正常終了 : コマンド名のみを指定",
			arg:            "i2i",
			expectedStatus: exitCodeOK,
		}, {
			desc:           "正常終了 : ショートオプションを指定",
			arg:            "i2i -b jpeg -a png",
			expectedStatus: exitCodeOK,
		}, {
			desc:           "正常終了 : ロングオプションを指定",
			arg:            "i2i -before jpeg -after png",
			expectedStatus: exitCodeOK,
		}, {
			desc:           "異常終了：不正なオプションを指定",
			arg:            "i2i -u unkonwn",
			expectedStatus: exitCodeErr,
		},
	}

	outSteam, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &cli{outStream: outSteam, errStream: errStream}

	for _, tc := range tests {
		args := strings.Split(tc.arg, " ")
		status := cli.run(args)
		if status != tc.expectedStatus {
			t.Errorf("desc: %v, status should be %v, not %v", tc.desc, tc.expectedStatus, status)
		}
	}
}
