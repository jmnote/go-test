package mymock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockCatClient struct{ mock.Mock }

func newMockCatClient() *mockCatClient { return &mockCatClient{} }

func (m *mockCatClient) HTTPGet(name string) (int, string) {
	args := m.Called(name)
	return args.Int(0), args.String(1)
}

func TestGetWithMock(t *testing.T) {
	m := newMockCatClient()
	m.On("HTTPGet", "Alice").Return(200, "Hello Alice").Once()
	m.On("HTTPGet", "Bob").Return(200, "Bye Bob").Once()
	m.On("HTTPGet", "Carol").Return(403, "limit exceeded")
	catClient = m

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
