package myhttp3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginEndpoint(t *testing.T) {
	assert.HTTPStatusCode(t, loginEndpoint, "GET", "", nil, 200)
	assert.HTTPBodyContains(t, loginEndpoint, "GET", "", nil, "pong")
}

func TestSubmitEndpoint(t *testing.T) {
	assert.HTTPStatusCode(t, submitEndpoint, "GET", "", nil, 200)
	assert.HTTPBodyContains(t, submitEndpoint, "GET", "", nil, "pong")
}

func TestReadEndpoint(t *testing.T) {
	assert.HTTPStatusCode(t, readEndpoint, "GET", "", nil, 200)
	assert.HTTPBodyContains(t, readEndpoint, "GET", "", nil, "pong")
}
