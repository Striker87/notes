package storage

import (
	"github.com/Striker87/notes/pkg/client/postgresql"
	"github.com/Striker87/notes/pkg/logging"
)

type storage struct {
	client postgresql.Client
	logger *logging.Logger
}
