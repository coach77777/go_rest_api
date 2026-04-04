package models


type Event struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Location    string `json:"location"`
	DateTime    time.Time `json:"timeTime"`
	UserID      int    `json:"userId"`


}

var events = []Event{}

func (e Event) Save(){

	//to dataabase


	events = append(events, e)
}