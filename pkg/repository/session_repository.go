package repository

import (
	"bytes"
	"errors"
	"time"

	"github.com/go-kit/kit/log"

	. "github.com/hecomp/session-management/internal/models"
	"github.com/hecomp/session-management/pkg/in_memory"
)

var (
	// ErrEmpty is returned when input string is empty
	ErrEmpty = errors.New("Empty session id")
	ErrNotFound = errors.New("Session id not found")
	ErrInvalidSessionId = errors.New("Invalid session id")
)

// SessionMgmntRepository
type SessionMgmntRepository interface {
	Create(sessionId string, request *SessionRequest) error
	Destroy(session *DestroyRequest) error
	Extend(request *ExtendRequest) error
	Exist(sessionId string) (bool, error)
	List() (*Sessions, error)
}

// AuthRepository has the implementation of the db methods.
type sessionMgmntRepository struct {
	store *in_memory.InMemStore
	logger log.Logger
}

// NewSessionMgmntRepository
func NewSessionMgmntRepository(store *in_memory.InMemStore, logger log.Logger) SessionMgmntRepository {
	return &sessionMgmntRepository{store: store, logger: logger}
}

// Create session is stored in-memory
func (s *sessionMgmntRepository) Create(sessionId string, request *SessionRequest) error {
	if sessionId == "" {
		return ErrEmpty
	}

	expiration := time.Unix(request.TTL, 0).UTC()
	if err := s.store.Commit(sessionId, []byte(sessionId), expiration); err != nil {
		return err
	}
	return nil
}

// Destroy remove the session from its cache
func (s *sessionMgmntRepository) Destroy(session *DestroyRequest) error {
	if err := s.store.Delete(session.SessionId); err != nil {
		return err
	}
	return nil
}

// Extend with the provided TTL
func (s *sessionMgmntRepository) Extend(request *ExtendRequest) error {

	expiration := time.Unix(request.TTL, 0).UTC()
	s.store.Reset(request.SessionId, expiration)

	b, found, err := s.store.Reset(request.SessionId, expiration)
	if err != nil {
		return err
	}
	if found != true {
		return ErrNotFound
	}
	if bytes.Equal(b, []byte(request.SessionId)) == false {
		return ErrInvalidSessionId
	}
	return nil
}

// Exist
func (s *sessionMgmntRepository) Exist(sessionId string) (bool, error) {
	b, found, err := s.store.Find(sessionId)
	if err != nil {
		return false, err
	}
	if found != true {
		return false, ErrNotFound
	}
	if bytes.Equal(b, []byte(sessionId)) == false {
		return false, ErrInvalidSessionId
	}
	return true, nil
}

//List returns a list of all the sessions that the service is currently tracking
func (s *sessionMgmntRepository) List() (*Sessions, error) {
	panic("implement me")
}
