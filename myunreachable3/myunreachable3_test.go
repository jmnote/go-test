package myunreachable3

import (
	"testing"

	"github.com/kuoss/common/tester"
	"github.com/stretchr/testify/assert"
)

func ExampleGreet() {
	Greet(true)
	// Output: hello
}
func TestGreet_hello(t *testing.T) {
	stdout, stderr, err := tester.CaptureChildTest(func() {
		Greet(true)
	})
	assert.Equal(t, "hello\n", stdout)
	assert.Equal(t, "", stderr)
	assert.NoError(t, err)
}

func TestGreet_bye(t *testing.T) {
	stdout, stderr, err := tester.CaptureChildTest(func() {
		Greet(false)
	})
	assert.Equal(t, "", stdout)
	// 2023/07/13 22:13:07 bye
	assert.Regexp(t, `^[0-9:/ ]+ bye\n$`, stderr)
	assert.EqualError(t, err, "exit status 1")
}
