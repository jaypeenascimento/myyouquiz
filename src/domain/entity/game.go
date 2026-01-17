package entity

type Game struct {
	KingPlayer  string   `json:"kingPlayer"`
	LetterQueue []Letter `json:"letterQueue"`
	ScoreToWin  int      `json:"scoreToWin"`
}
