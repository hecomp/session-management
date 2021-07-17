package repository

import (
	"github.com/go-kit/kit/log"

	. "github.com/hecomp/session-management/internal/models"
)

// SessionMgmntRepository
type SessionMgmntRepository interface {
	Create(request Session) error
	Destroy(session *DestroyRequest) error
	Extend(request *ExtendRequest) error
	List() Sessions
}

// AuthRepository has the implementation of the db methods.
type sessionMgmntRepository struct {
	logger log.Logger
}

// NewSessionMgmntRepository
func NewSessionMgmntRepository(logger log.Logger)  SessionMgmntRepository{
	return &sessionMgmntRepository{logger}
}

// Create session is stored in-memory
func (s sessionMgmntRepository) Create(request Session) error {
	panic("implement me")
}

// Destroy remove the session from its cache
func (s sessionMgmntRepository) Destroy(session *DestroyRequest) error {
	panic("implement me")
}

//Extend
func (s sessionMgmntRepository) Extend(request *ExtendRequest) error {
	panic("implement me")
}

//List returns a list of all the sessions that the service is currently tracking
func (s sessionMgmntRepository) List() Sessions {
	panic("implement me")
}
