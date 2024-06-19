package domain

type Profile struct {
	Date    string `json:"date"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Skills  string `json:"skills"`
	Age     int    `json:"age"`
	Birth   string `json:"birth"`
}

type Profiles struct {
	Date      string `json:"date"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       string `json:"age"`
}
