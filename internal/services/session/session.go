package session

import (
	"database/sql"

	db "github.com/valeno12/kalkulapp/internal/models"
)

type SessionService struct {
	db      *sql.DB
	queries *db.Queries
}

func NewSessionService(db *sql.DB, q *db.Queries) *SessionService {
	return &SessionService{db: db, queries: q}
}
