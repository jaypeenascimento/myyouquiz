package entities

type Game struct {
	KingPlayer  string `json:"kingPlayer"`
	LetterQueue string `json:"letterQueue"`
	ScoreToWin  string `json:"scoreToWin"`
}
