package dto

type CreateSessionRequest struct {
	SessionName     string `json:"session_name" validate:"required"`
	UserName        string `json:"user_name" validate:"required"`
	MaxParticipants *int   `json:"max_participants"`
}

type CreateSessionResponse struct {
	SessionID int64  `json:"session_id"`
	Code      string `json:"code"`
}

type JoinSessionRequest struct {
	UserName string `json:"user_name" validate:"required"`
}
