package models

// Session  represents th etype for the TTL as an optional param to create
type Session struct {
	TTL string `json:"ttl"`
}
type DestroyRequest struct {
	SessionId string `json:"session_id" validate:"required"`
}

type ExtendRequest struct {
	TTL string `json:"ttl"`
	SessionId string `json:"session_id" validate:"required"`
}

type Sessions struct {
	List []string `json:"list"`
}

