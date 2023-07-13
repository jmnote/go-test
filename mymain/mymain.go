package mymain

import (
	"os"
	"testing"

	ms "github.com/kuoss/venti/pkg/mock/servers"
)

var servers *ms.Servers

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	servers = ms.New(ms.Requirements{
		{Type: ms.TypeAlertmanager, Port: 0, Name: "alertmanager1", IsMain: false, BasicAuth: false},
		{Type: ms.TypeLethe, Port: 0, Name: "lethe1", IsMain: true, BasicAuth: false},
		{Type: ms.TypeLethe, Port: 0, Name: "lethe2", IsMain: false, BasicAuth: false},
		{Type: ms.TypePrometheus, Port: 0, Name: "prometheus1", IsMain: true, BasicAuth: false},
		{Type: ms.TypePrometheus, Port: 0, Name: "prometheus2", IsMain: false, BasicAuth: false},
		{Type: ms.TypePrometheus, Port: 0, Name: "prometheus3", IsMain: false, BasicAuth: true},
	})
}

func teardown() {
	servers.Close()
}
