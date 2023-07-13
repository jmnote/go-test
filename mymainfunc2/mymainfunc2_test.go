package main

import (
	"os"
	"testing"

	"github.com/kuoss/common/tester"
	"github.com/stretchr/testify/assert"
)

func TestMainFunc(t *testing.T) {
	stdout, stderr, err := tester.CaptureChildTest(func() {
		main()
	})
	assert.Equal(t, "loadKubeconfig ok\n", stdout)
	assert.Equal(t, "", stderr)
	assert.NoError(t, err)
}

func TestMainFunc_2(t *testing.T) {
	stdout, stderr, err := tester.CaptureChildTest(func() {
		os.Setenv("KUBECONFIG", "")
		main()
	})
	assert.Equal(t, "loadKubeconfig ok\n", stdout)
	assert.Equal(t, "", stderr)
	assert.NoError(t, err)
}

func TestMainFunc_3(t *testing.T) {
	stdout, stderr, err := tester.CaptureChildTest(func() {
		os.Setenv("KUBECONFIG", "exists.yaml")
		main()
	})
	assert.Equal(t, "loadKubeconfig ok\n", stdout)
	assert.Equal(t, "", stderr)
	assert.NoError(t, err)
}

func TestMainFunc_4(t *testing.T) {
	stdout, stderr, err := tester.CaptureChildTest(func() {
		os.Setenv("KUBECONFIG", "not_exists.yaml")
		main()
	})
	assert.Equal(t, "", stdout)
	// "2023/06/25 18:45:24 loadKubeconfig err: ReadFile err: open not_exists.yaml: no such file or directory\n"
	assert.Contains(t, stderr, "loadKubeconfig err: ReadFile err: open not_exists.yaml")
	assert.EqualError(t, err, "exit status 1")
}
