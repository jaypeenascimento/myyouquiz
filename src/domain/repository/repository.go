// Package repository groups the repositories from the services at domain
package repository

import "github.com/jaypeenascimento/myyouquiz/src/domain/entity"

type PlayerRepository interface {
	PutPlayer(entity.Player) (entity.Player, error)
	FindPlayer(name string) (entity.Player, error)
	FetchAllPlayers() ([]entity.Player, error)
}
