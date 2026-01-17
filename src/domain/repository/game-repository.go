package repository

import "github.com/jaypeenascimento/myyouquiz/src/domain/entity"

type GameRepository interface {
	GetGame() (entity.Game, error)
}
