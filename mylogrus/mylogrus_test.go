package mylogrus

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogHello(t *testing.T) {
	got := LogHello()
	t.Log(got)
	assert.Contains(t, got, "level=info msg=Hello")
	assert.Regexp(t, `time="[^"]+" level=info msg=Hello`, got)
}
