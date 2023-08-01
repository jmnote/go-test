package mylogrus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogString(t *testing.T) {
	got := LogString("Hello")
	t.Log(got)
	assert.Contains(t, got, "level=info msg=Hello")
	assert.Regexp(t, `time="[^"]+" level=info msg=Hello`, got)
}
