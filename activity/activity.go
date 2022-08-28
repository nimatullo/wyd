package activity
/* 
This struct keeps track of the current website that I'm on.
The ready field indicates to the event stream whenever this activity
needs to be sent to the frontend. After an activity is sent, the ready 
field is set to false.
*/
type Activity struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Since   string `json:"since"`
	Ready   bool   `json:"ready"`
}

/*
This is a global variable that keeps track of the current activity app-wide.
I chose to use a global variable to avoid having to read from the database 
everytime there is a read on the stream endpoint.
*/
var CURRENT_ACTIVITY Activity = Activity{
	Name:    "",
	Website: "",
	Since:   "",
	Ready:   false,
}
