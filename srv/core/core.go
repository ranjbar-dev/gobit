package core

import "github.com/ranjbar-dev/gobit/srv/repository"

type Core struct {
	repo *repository.Repository
	// state
	currentBlockNumber int64
}

func NewCore(repo *repository.Repository) *Core {

	return &Core{
		repo: repo,
	}
}
