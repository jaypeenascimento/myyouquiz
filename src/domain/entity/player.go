package entity

type Player struct {
	Name           string `json:"name"`
	Score          int    `json:"score"`
	CanGetQuestion bool   `json:"canGetQuestion"`
}

func NewPlayer(name string) *Player {
	return &Player{Name: name}
}
