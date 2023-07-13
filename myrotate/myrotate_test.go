package myrotate

import (
	"os/exec"
	"testing"
	"time"

	"github.com/jmnote/go-test/myclock"
	"github.com/stretchr/testify/assert"
)

func init() {
	myclock.SetPlaygroundMode(true)
}

func resetData() {
	cmd := exec.Command("sh", "-c", "rm -rf data; cp -a testdata data")
	_, err := cmd.Output()
	if err != nil {
		panic(err)
	}
}

func TestDeleteByAge(t *testing.T) {
	testCases := []struct {
		retentionTime time.Duration
		fileExists    map[string]bool
	}{
		{
			24 * time.Hour, // 1 day
			map[string]bool{"2009-10-30.log": false, "2009-11-05.log": false, "2009-11-10.log": true},
		},
		{
			7 * 24 * time.Hour, // 7 days
			map[string]bool{"2009-10-30.log": false, "2009-11-05.log": true, "2009-11-10.log": true},
		},
		{
			14 * 24 * time.Hour, // 14 days
			map[string]bool{"2009-10-30.log": true, "2009-11-05.log": true, "2009-11-10.log": true},
		},
	}
	for _, tc := range testCases {
		resetData()
		deleteByAge(tc.retentionTime)
		for f, exist := range tc.fileExists {
			if exist {
				assert.FileExists(t, "data/"+f)
			} else {
				assert.NoFileExists(t, "data/"+f)
			}
		}
	}
}
