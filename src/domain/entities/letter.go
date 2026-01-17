package entities

type Letter struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Sender   string `json:"sender"`
}
