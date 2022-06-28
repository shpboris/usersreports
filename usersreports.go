package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shpboris/usersdata"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"usersreports/logger"
	"usersreports/reportsvc"
)

const (
	acceptHeader    = "Accept"
	applicationJson = "application/json"
	GET             = "GET"
	URL             = "http://localhost:8000/users"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/report", GenerateReport).Methods(GET).Headers(acceptHeader, applicationJson)
	logger.Log.Info("Started the server on port 8001")
	log.Fatal(http.ListenAndServe(":8001", router))
}

func GenerateReport(w http.ResponseWriter, r *http.Request) {
	logger.Log.Debug("Started GenerateReport")
	client := http.Client{}
	req, err := http.NewRequest(GET, URL, nil)
	if handleError(err, w) {
		return
	}
	req.Header = http.Header{
		acceptHeader: {applicationJson},
	}
	res, err := client.Do(req)
	if handleError(err, w) {
		return
	}
	responseData, err := ioutil.ReadAll(res.Body)
	if handleError(err, w) {
		return
	}
	var users []usersdata.User
	err = json.Unmarshal(responseData, &users)
	if handleError(err, w) {
		return
	}
	var reportSummary = reportsvc.GetReportData(users)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reportSummary)
	logger.Log.Debug("Completed GenerateReport")
}

func handleError(err error, w http.ResponseWriter) bool {
	if err != nil {
		logger.Log.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return true
	}
	return false
}
