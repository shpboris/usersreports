package reportsvc

import (
	"github.com/shpboris/usersdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetReportData(t *testing.T) {
	users := []usersdata.User{
		{
			Id:     "1",
			Name:   "alex",
			Unit:   "pex",
			Salary: 1000,
		},
		{
			Id:     "2",
			Name:   "boris",
			Unit:   "edge",
			Salary: 2000,
		},
		{
			Id:     "3",
			Name:   "roni",
			Unit:   "edge",
			Salary: 5000,
		},
	}
	reportData := GetReportData(users)
	assert.Equal(t, 2, len(reportData))
	found := false
	for _, curr := range reportData {
		if curr.Unit == "edge" {
			found = true
			assert.Equal(t, curr.Budget, 7000)
		}
	}
	assert.Equal(t, true, found)
}
