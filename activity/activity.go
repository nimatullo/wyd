package activity

type Activity struct {
	Name    string `json:"name"`
	Website string `json:"website"`
	Since   string `json:"since"`
	Ready   bool   `json:"ready"`
}

var CURRENT_ACTIVITY Activity = Activity{
	Name:    "",
	Website: "",
	Since:   "",
	Ready:   false,
}
