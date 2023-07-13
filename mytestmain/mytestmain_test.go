package mytestmain

import (
	"os"
	"testing"

	"github.com/kuoss/venti/pkg/mocker"
	"github.com/kuoss/venti/pkg/mocker/prometheus"
	"github.com/stretchr/testify/assert"
)

var server *mocker.Server

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	var err error
	server, err = prometheus.New(0)
	if err != nil {
		panic(err)
	}
	host = server.URL
}

func teardown() {
	server.Close()
}

func TestGetMetadata(t *testing.T) {
	want := `{"status":"success","data":{"apiserver_audit_event_total":[{"type":"counter","help":"[ALPHA] Counter of audit events generated and sent to the audit backend.","unit":""}]}}`
	got, err := GetMetadata()
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
