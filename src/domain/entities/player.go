package entities

type Player struct {
	Score          int  `json:"score"`
	CanGetQuestion bool `json:"canGetQuestion"`
}
