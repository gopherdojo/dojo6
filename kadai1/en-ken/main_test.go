package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	exec.Command("go", "build", "-o", "kadai1").Run()
	m.Run()
}

func TestMainSuccess(t *testing.T) {

	cmdString := "./kadai1 ./testdata -input-ext .jpg -output-dir ./out -output-ext .png"
	cmd := strings.Split(cmdString, " ")
	err := exec.Command(cmd[0], cmd[1:]...).Run()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMainFailure(t *testing.T) {

	cmdString := "./kadai1"
	cmd := strings.Split(cmdString, " ")
	err := exec.Command(cmd[0], cmd[1:]...).Run()
	if err == nil {
		t.Fatal(err)
	}
}
