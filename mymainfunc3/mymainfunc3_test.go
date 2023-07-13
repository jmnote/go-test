package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainFunc(t *testing.T) {
	main()
}

func TestMainFunc_2(t *testing.T) {
	os.Setenv("KUBECONFIG", "")
	main()
}

func TestMainFunc_3(t *testing.T) {
	os.Setenv("KUBECONFIG", "exists.yaml")
	main()
}

func TestMainFunc_4(t *testing.T) {
	if os.Getenv("CHILD") == "1" {
		os.Setenv("KUBECONFIG", "not_exists.yaml")
		main()
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMainFunc_4")
	cmd.Env = append(os.Environ(), "CHILD=1")
	err := cmd.Run()
	assert.EqualError(t, err, "exit status 1")
}
