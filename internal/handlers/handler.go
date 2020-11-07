package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"helpnow/internal/database"
	"helpnow/internal/object"
	"helpnow/internal/util"
	"net/http"
	"strconv"
	"strings"
	"time"
)


func Meetings(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		var resp *object.GetAllMeetingResponse
		var id int
		var err error
		
		uri_date := strings.Split(util.After(r.RequestURI, "/api/v1/meetings?"), "?")
		if len(uri_date) > 0 {
			query := r.URL.Query()
			participant := query.Get("participant")
			if participant != "" {
				resp = database.FetchAllMeetingsForParticipant(participant)
		    	w.WriteHeader(http.StatusOK)
				byteresp, _ := json.Marshal(resp)
				w.Write([]byte(byteresp))
				return
			}

		    start_time := query.Get("start") 
		    end_time := query.Get("end")
			fmt.Println(start_time,end_time)
		    if start_time != "" || end_time != "" {
		    	startTime, err := time.Parse("2006-01-02T15:04:04", start_time)
				if err != nil {
					panic(err)
				}
				endTime, err := time.Parse("2006-01-02T15:04:04", end_time)
				if err != nil {
					panic(err)
				}
				if endTime.Before(startTime) {
					resp := object.Response{
						Message: "end_time should be greater than start_time",
					}
					w.WriteHeader(http.StatusMethodNotAllowed)
					byteresp, _ := json.Marshal(resp)
					w.Write([]byte(byteresp))
					return
				}
		    	resp = database.FetchMeetingsWithinRange(startTime,endTime)
		    	w.WriteHeader(http.StatusOK)
				byteresp, _ := json.Marshal(resp)
				w.Write([]byte(byteresp))
				return
			}
		}

		uri := strings.Split(util.After(r.RequestURI, "/api/v1/meetings/"), "/")
		if uri[0] != "" {
			id, err = CheckURIAndRetrieveID(uri)
			if err != nil {
				resp = &object.GetAllMeetingResponse{
					Message: err.Error(),
				}
				w.WriteHeader(http.StatusNotFound)
				byteresp, _ := json.Marshal(resp)
				w.Write(byteresp)
				return
			}
		}
		fmt.Println(uri_date)
		if id != 0 {
			resp = database.FetchMeetingsByID(id)
		} else {
			resp = database.FetchAllMeetings()
		}
		w.WriteHeader(http.StatusOK)
		byteresp, _ := json.Marshal(resp)
		w.Write([]byte(byteresp))
	case "POST":
		var reqobj object.AddMeetingRequest
		var resp *object.Response
		if err := json.NewDecoder(r.Body).Decode(&reqobj); err != nil {
			resp = &object.Response{
				Message: err.Error(),
			}
		}
		resp = database.AddMeetings(&reqobj)
		byteresp, _ := json.Marshal(resp)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(byteresp))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		response := object.Response{
			Message: "Method Not Supported",
		}
		byteresp, _ := json.Marshal(response)
		w.Write(byteresp)
	}
}


// CheckURIAndRetrieveID Method checks uri and Fetches Id when Operation By Id is requested
//It allows baseurl/{id} , returns error (method not found) if pattern is not there
//Else returns id
func CheckURIAndRetrieveID(uri []string) (int, error) {
	var id int
	var err error
	/*if len(uri) > 1 || len(uri) < 1 {
		return 0, errors.New("Method Not Found")
	}*/
	id, err = strconv.Atoi(uri[0])
	if err != nil {
		return 0, errors.New("Method Not Found")
	}
	return id, nil
}
