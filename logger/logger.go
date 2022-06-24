package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

var Log = logrus.New()

func init() {
	Log.SetLevel(getLogLevel())
	f, err := os.OpenFile(getLogFilePath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Log.Fatalf("Error opening file: %v", err)
	}
	wrt := io.MultiWriter(os.Stdout, f)
	Log.SetOutput(wrt)
	Log.SetReportCaller(getReportCaller())
}

func getLogLevel() logrus.Level {
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "debug"
	}
	ll, err := logrus.ParseLevel(lvl)
	if err != nil {
		ll = logrus.DebugLevel
	}
	return ll
}

func getLogFilePath() string {
	path, ok := os.LookupEnv("LOG_FILE_PATH")
	if !ok {
		path = "app.log"
	}
	return path
}

func getReportCaller() bool {
	reportCaller, ok := os.LookupEnv("REPORT_CALLER")
	if !ok || strings.EqualFold(reportCaller, "y") {
		return true
	}
	return false
}
