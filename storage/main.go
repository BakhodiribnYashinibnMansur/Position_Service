package storage

import (
	"position_server/storage/postgres"
	"position_server/storage/repo"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Profession() repo.ProfessionRepoI
}

type storagePG struct {
	profession repo.ProfessionRepoI
}

func NewStoragePG(db *sqlx.DB) StorageI {
	return &storagePG{
		profession: postgres.NewProfessionRepo(db),
	}
}

func (s *storagePG) Profession() repo.ProfessionRepoI {
	return s.profession
}
