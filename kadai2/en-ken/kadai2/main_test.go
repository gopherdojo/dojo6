package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	exec.Command("go", "get", "github.com/gopherdojo/dojo6/kadai2/en-ken/kadai2").Run()
	m.Run()
}

func TestMainSuccess(t *testing.T) {

	cmdString := "kadai2 ./testdata -input-ext .jpg -output-dir ./out -output-ext .png"
	cmd := strings.Split(cmdString, " ")
	err := exec.Command(cmd[0], cmd[1:]...).Run()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMainFailure(t *testing.T) {

	cmdString := "kadai2"
	cmd := strings.Split(cmdString, " ")
	err := exec.Command(cmd[0], cmd[1:]...).Run()
	if err == nil {
		t.Fatal(err)
	}
}
