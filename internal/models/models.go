package models

// Item
type Item struct {
	Oject      []byte
	Expiration int64
}

// SessionRequest  represents th etype for the TTL as an optional param to create
type SessionRequest struct {
	TTL int64 `json:"ttl"`
}

type DestroyRequest struct {
	SessionId string `json:"session_id" validate:"required"`
}

type ExtendRequest struct {
	TTL int64 `json:"ttl"`
	SessionId string `json:"session_id" validate:"required"`
}

// Session  represents th etype for the TTL as an optional param to create
type Session struct {
	SessionId string `json:"session_id" validate:"required"`
}

type Sessions struct {
	List []string `json:"list"`
}

