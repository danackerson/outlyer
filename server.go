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
	commands.StoreMetrics() // initialize the metrics storage registry
	gocron.Every(1).Seconds().Do(commands.StoreMetrics)
	<-gocron.Start()
}

func startAPIServer() {
	r := mux.NewRouter()
	setUpRoutes(r)
	n := negroni.Classic()

	n.UseHandler(r)
	log.Printf("SERVING HTTP...")
	http.ListenAndServe(httpPort, n)
}

func setUpRoutes(router *mux.Router) {
	router.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		getMetrics(w, r)
	})
}

func getMetrics(w http.ResponseWriter, req *http.Request) {
	lastBaseMetricSample := commands.GetLastMetricSample()

	data, err := json.Marshal(lastBaseMetricSample)
	if err != nil {
		errText := fmt.Sprintf("ERR: unable to marshal JSON baseMetrics object: %s", err.Error())
		http.Error(w,
			errText,
			http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(data)
}
