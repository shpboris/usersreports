package reportsvc

import (
	"github.com/shpboris/usersdata"
	"usersreports/reportdata"
)

func GetReportData(users []usersdata.User) []*reportdata.ReportData {
	var unitToReportDataMap = make(map[string]*reportdata.ReportData)
	var reportSummary = make([]*reportdata.ReportData, 0)
	if len(users) > 0 {
		for _, user := range users {
			if reportData, ok := unitToReportDataMap[user.Unit]; ok {
				reportData.Budget += user.Salary
			} else {
				unitToReportDataMap[user.Unit] = &reportdata.ReportData{Unit: user.Unit, Budget: user.Salary}
			}
		}
		for _, reportData := range unitToReportDataMap {
			reportSummary = append(reportSummary, reportData)
		}
	}
	return reportSummary
}
