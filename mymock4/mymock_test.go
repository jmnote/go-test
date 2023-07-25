package mymock

import (
	"testing"

	"github.com/jmnote/go-test/mymock4/catclient"
	"github.com/stretchr/testify/assert"
)

type mockCatClient struct{}

func newMockCatClient() catclient.ICatClient {
	return &mockCatClient{}
}

func (m *mockCatClient) HTTPGet(name string) (int, string) {
	switch name {
	case "Alice":
		return 200, "Hello Alice"
	case "Bob":
		return 200, "Bye Bob"
	case "Carol":
		return 403, "limit exceeded"
	}
	return 0, ""
}

func TestGet(t *testing.T) {
	catClient = newMockCatClient()
	testCases := []struct {
		name      string
		want      string
		wantError string
	}{
		{"Alice", "Hello Alice", ""},
		{"Bob", "Bye Bob", ""},
		{"Carol", "", "err: limit exceeded"},
	}
	for _, tc := range testCases {
		t.Run("", func(t *testing.T) {
			got, err := Get(tc.name)
			if tc.wantError == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.wantError)
			}
			assert.Equal(t, tc.want, got)
		})
	}
}
