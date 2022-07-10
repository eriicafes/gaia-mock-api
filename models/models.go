package models

import "github.com/eriicafes/filedb"

type model struct {
	db *filedb.Database
}

func New(db *filedb.Database) *model {
	return &model{db}
}
