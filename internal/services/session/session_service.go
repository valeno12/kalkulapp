package services

import (
	"context"
	"database/sql"
	"math/rand"
	"time"

	db "github.com/valeno12/kalkulapp/internal/models"
)

type SessionService struct {
	queries *db.Queries
}

func NewSessionService(q *db.Queries) *SessionService {
	return &SessionService{queries: q}
}

func (s *SessionService) CreateSession(ctx context.Context, req CreateSessionRequest) (int64, string, error) {
	code := generateSessionCode()

	params := db.CreateSessionParams{
		Name:            req.SessionName,
		Code:            code,
		CreatedBy:       0,
		MaxParticipants: sqlNullInt32(req.MaxParticipants),
	}

	result, err := s.queries.CreateSession(ctx, params)
	if err != nil {
		return 0, "", err
	}

	sessionID, err := result.LastInsertId()
	if err != nil {
		return 0, "", err
	}

	_, err = s.queries.CreateUser(ctx, db.CreateUserParams{
		SessionID: sessionID,
		Name:      req.UserName,
	})
	if err != nil {
		return 0, "", err
	}

	return sessionID, code, nil
}


func sqlNullInt32(val *int) sql.NullInt32 {
	if val == nil {
		return sql.NullInt32{Valid: false}
	}
	return sql.NullInt32{Int32: int32(*val), Valid: true}
}

func generateSessionCode() string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, 6)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}
