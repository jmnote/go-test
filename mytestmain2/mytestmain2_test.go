package mytestmain2

import (
	"testing"

	"github.com/kuoss/venti/pkg/mocker"
	"github.com/kuoss/venti/pkg/mocker/prometheus"
	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	server *mocker.Server
}

func (suite *ExampleTestSuite) SetupSuite() {
	var err error
	suite.server, err = prometheus.New(0)
	if err != nil {
		panic(err)
	}
	host = suite.server.URL
}

func (suite *ExampleTestSuite) TearDownSuite() {
	suite.server.Close()
}

func (suite *ExampleTestSuite) TestGetMetadata() {
	want := `{"status":"success","data":{"apiserver_audit_event_total":[{"type":"counter","help":"[ALPHA] Counter of audit events generated and sent to the audit backend.","unit":""}]}}`
	got, err := GetMetadata()
	suite.NoError(err)
	suite.Equal(want, got)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}
