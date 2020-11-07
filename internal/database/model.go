package database

import (

	"fmt"
	"helpnow/internal/object"
	"time"

)


//AddMeetings Adds Movie into Database, Returns Success Message or Error as Message
func AddMeetings(request *object.AddMeetingRequest) *object.Response {
	db = GetConnection()
	startTime, err := time.Parse("2006-01-02 15:04:04", request.Meetings.StartTIme)
	if err != nil {
		panic(err)
	}
	endTime, err := time.Parse("2006-01-02 15:04:04", request.Meetings.EndTime)
	if err != nil {
		panic(err)
	}
	if endTime.Before(startTime) {
		return &object.Response{
			Message: "end_time is smaller than start_time",
		}
	}
	exists, err := CheckMeetingExists(request,startTime,endTime)
	if len(exists.Meetings) > 0 {
		return &object.Response{
			Message: "Meeting overlapped",
		}
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}

	db = GetConnection()
	insertStatement := `INSERT INTO meetings (title, participant_name, participant_email, start_time, end_time, rsvp)  VALUES($1,$2,$3,$4,$5,$6)`

	_, err = db.Exec(insertStatement,request.Meetings.Title,request.Meetings.Participants.Name, request.Meetings.Participants.Email,startTime,endTime,request.Meetings.Rsvp)
	if err != nil {
		return &object.Response{
			Message: err.Error(),
		}
	}
	
	return &object.Response{
		Message: "Meetings Added",
	}
}


func CheckMeetingExists(request *object.AddMeetingRequest, startTime time.Time, endTime time.Time) (*object.GetAllMeetingResponse, error) {
	db = GetConnection()
	var meeting object.Meeting
	var meetings []object.Meeting
	rsvp := "yes"
	query := `Select id, title, participant_name, participant_email, start_time,end_time, timestamp from meetings where participant_email = $1 and ((start_time <= $2 and end_time >= $3 ) or (start_time >= $2 and end_time >= $3) or (start_time >= $2 and start_time <= $3 ) or (end_time >= $2 and end_time <= $3)) and rsvp = $4`
	rows, err := db.Query(query,request.Meetings.Participants.Email,startTime,endTime,rsvp)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&meeting.Id,&meeting.Title,&meeting.Participants.Name, &meeting.Participants.Email,&meeting.StartTIme,&meeting.EndTime,&meeting.Timestamp)
		if err != nil {
			return nil, err
		}
		meetings = append(meetings, meeting)
	}

	return &object.GetAllMeetingResponse{
		Message: "overlapped meetings",
		Meetings:  meetings,
	}, nil
	
}


//FetchAllMeetings Fetches All Meetings from Db and returns list of meetings and success message or error as message
func FetchAllMeetings() *object.GetAllMeetingResponse {
	db = GetConnection()
	var meeting object.Meeting
	var meetings []object.Meeting
	
	query := "Select id, title, participant_name, participant_email, start_time,end_time, timestamp from meetings"
	rows, err := db.Query(query)
	if err != nil {
		return &object.GetAllMeetingResponse{
			Message: err.Error(),
		}
	}
	for rows.Next() {
		err := rows.Scan(&meeting.Id,&meeting.Title,&meeting.Participants.Name, &meeting.Participants.Email,&meeting.StartTIme,&meeting.EndTime,&meeting.Timestamp)
		if err != nil {
			return &object.GetAllMeetingResponse{
				Message: err.Error(),
			}
		}
		meetings = append(meetings, meeting)
	}
	return &object.GetAllMeetingResponse{
		Message: "SuccessFully Retrieved All Meetings",
		Meetings:  meetings,
	}
}



func FetchAllMeetingsForParticipant(participant string) *object.GetAllMeetingResponse {
	db = GetConnection()
	var meeting object.Meeting
	var meetings []object.Meeting
	query := "Select id, title, participant_name, participant_email, start_time,end_time, timestamp, rsvp from meetings where participant_email = $1"
	rows, err := db.Query(query, participant)
	if err != nil {
		return &object.GetAllMeetingResponse{
			Message: err.Error(),
		}
	}
	for rows.Next() {
		err := rows.Scan(&meeting.Id,&meeting.Title,&meeting.Participants.Name, &meeting.Participants.Email,&meeting.StartTIme,&meeting.EndTime,&meeting.Timestamp, &meeting.Rsvp)
		if err != nil {
			return &object.GetAllMeetingResponse{
				Message: err.Error(),
			}
		}
		meetings = append(meetings, meeting)
	}
	return &object.GetAllMeetingResponse{
		Message: "SuccessFully Retrieved Meetings",
		Meetings:  meetings,
	}
}


func FetchMeetingsByID(id int) *object.GetAllMeetingResponse {
	db = GetConnection()
	var meeting object.Meeting
	var meetings []object.Meeting
query := "Select id, title, participant_name, participant_email, start_time,end_time, timestamp, rsvp from meetings where id = $1"
	rows, err := db.Query(query, id)
	if err != nil {
		return &object.GetAllMeetingResponse{
			Message: err.Error(),
		}
	}
	for rows.Next() {
		err := rows.Scan(&meeting.Id,&meeting.Title,&meeting.Participants.Name, &meeting.Participants.Email,&meeting.StartTIme,&meeting.EndTime,&meeting.Timestamp,  &meeting.Rsvp)
		if err != nil {
			return &object.GetAllMeetingResponse{
				Message: err.Error(),
			}
		}
		meetings = append(meetings, meeting)
	}
	return &object.GetAllMeetingResponse{
		Message: "SuccessFully Retrieved Meetings",
		Meetings:  meetings,
	}
}


func FetchMeetingsWithinRange(startTime time.Time, endTime time.Time) *object.GetAllMeetingResponse {
	db = GetConnection()
	var meeting object.Meeting
	var meetings []object.Meeting
	query := "Select id, title, participant_name, participant_email, start_time,end_time, timestamp, rsvp from meetings where start_time >= $1 and end_time <= $2"
	rows, err := db.Query(query, startTime, endTime)
	if err != nil {
		return &object.GetAllMeetingResponse{
			Message: err.Error(),
		}
	}
	for rows.Next() {
		err := rows.Scan(&meeting.Id,&meeting.Title,&meeting.Participants.Name, &meeting.Participants.Email,&meeting.StartTIme,&meeting.EndTime,&meeting.Timestamp,  &meeting.Rsvp)
		if err != nil {
			return &object.GetAllMeetingResponse{
				Message: err.Error(),
			}
		}
		meetings = append(meetings, meeting)
	}
	return &object.GetAllMeetingResponse{
		Message: "SuccessFully Retrieved Meetings",
		Meetings:  meetings,
	}
}
