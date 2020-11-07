package main

import (
	"flag"
	"fmt"
	. "helpnow/internal/database"
	"helpnow/internal/handlers"
	. "helpnow/internal/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	port    int
	logfile string
	ver     bool
)

//Init To Initialize Cache Maps and flags require on startup
func init() {
	flag.IntVar(&port, "port", 9090, "The port to listen on.")
	flag.StringVar(&logfile, "logfile", "", "Location of the logfile.")
	flag.BoolVar(&ver, "version", false, "Print server version.")
}

const (
	// base HTTP paths.
	apiVersion  = "v1"
	apiBasePath = "/api/" + apiVersion + "/"

	scheduleMeeting = apiBasePath + "meeting/schedule"
	searchMeeting = apiBasePath + "meetings/"
	addmeetingsPath = apiBasePath + "meetings"
	searchMeetingByDate = apiBasePath + "meetings?"
	// server version.
	version = "1.0.0"
)

//Main Function: Starts Server,Exposes Endpoint and Initialized DataBase Connection
func main() {
	flag.Parse()
	if ver {
		fmt.Printf("HTTP Server v%s", version)
		os.Exit(0)
	}
	var logger *log.Logger
	if logfile == "" {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	} else {
		f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		logger = log.New(f, "", log.LstdFlags)
	}
	err := InitConnection()
	if err != nil {
		panic(err)
	}
	http.Handle(scheduleMeeting,ServiceLoader(http.HandlerFunc(handlers.Meetings),RequestMetrics(logger)))
	http.Handle(searchMeeting,ServiceLoader(http.HandlerFunc(handlers.Meetings),RequestMetrics(logger)))	
	http.Handle(addmeetingsPath,ServiceLoader(http.HandlerFunc(handlers.Meetings),RequestMetrics(logger)))
	
	logger.Printf("starting server on :%d", port)

	strPort := ":" + strconv.Itoa(port)
	logger.Fatal(http.ListenAndServe(strPort, nil))
	logger.Printf("started server on :%d", port)
}
