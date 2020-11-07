package object

import(
	"time"
)

//Response object for Common Response
type Response struct {
	Message string `json:"message,omitempty"`
}


type Participant struct {
	Name	string	`json:"name"`
	Email	string	`json:"email"`
}

type Meeting struct{
	Id				int			`json:"id,omitempty`
	Title			string		`json:"title"`
	StartTIme		string	`json:"start_time"`
	EndTime			string	`json:"end_time"`
	Participants	Participant	`json:"participants"`
	Timestamp		time.Time	`json:"time_stamp,omitempty"`
	Rsvp 			string		`json:"rsvp,omitempty"`
}

type GetAllMeetingResponse struct {
	Meetings []Meeting `json:"meetings"`
	Message	 string	   `json:"message"`
}

//AddMovieRequest Object For Add Movie Request
type AddMeetingRequest struct {
	Meetings Meeting `json:"meetings"`
}
