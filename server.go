package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/danackerson/outlyer/commands"
	"github.com/gorilla/mux"
	"github.com/jasonlvhit/gocron"
	"github.com/urfave/negroni"
)

var httpPort = ":8080"

func getHTTPPort() string {
	return httpPort
}

func main() {
	go startMetricsDaemon()
	startAPIServer()
}

func startMetricsDaemon() {
	gocron.Every(1).Seconds().Do(commands.StoreMetricMeasurement)
	<-gocron.Start()
}

func startAPIServer() {
	r := mux.NewRouter()
	setUpRoutes(r)
	n := negroni.Classic()

	n.UseHandler(r)
	log.Printf("SERVING HTTP at http://localhost:%s/metrics", getHTTPPort())
	http.ListenAndServe(httpPort, n)
}

func setUpRoutes(router *mux.Router) {
	router.HandleFunc("/metrics",
		func(w http.ResponseWriter, r *http.Request) {
			getAllMetrics(w, r)
		})
}

func getAllMetrics(w http.ResponseWriter, req *http.Request) {
	data, err := json.Marshal(commands.GetAllMetrics())
	if err != nil {
		errText := fmt.Sprintf(
			"ERR: unable to marshal all JSON metrics object: %s", err.Error())
		http.Error(w,
			errText,
			http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}
