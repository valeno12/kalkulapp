package session

import (
	session "github.com/valeno12/kalkulapp/internal/services/session"
)

type SessionHandler struct {
	service *session.SessionService
}

func NewSessionHandler(service *session.SessionService) *SessionHandler {
	return &SessionHandler{service: service}
}
