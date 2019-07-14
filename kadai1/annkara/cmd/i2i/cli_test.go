package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	outSteam, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &cli{outStream: outSteam, errStream: errStream}
	status := cli.run()

	if status != 0 {
		t.Errorf("desc: Success, status should be %v, not %v", exitCodeOK, status)
	}
}
