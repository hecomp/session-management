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
	ErrEmpty = errors.New("empty session id")
	ErrExist = errors.New("error during find")
	ErrNotFound = errors.New("session id not found")
	ErrInvalidSessionId = errors.New("envalid session id")
)

// SessionMgmntRepository
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SessionMgmntRepository
type SessionMgmntRepository interface {
	Create(sessionId string, expiration time.Time) error
	Destroy(session *DestroyRequest) error
	Extend(request *ExtendRequest) (bool, error)
	Exist(sessionId string) (bool, error)
	List() (*Sessions, error)
}

// AuthRepository has the implementation of the db methods.
type sessionMgmntRepository struct {
	store in_memory.MemStore
	logger log.Logger
}

// NewSessionMgmntRepository create a instance of session management repository
func NewSessionMgmntRepository(store in_memory.MemStore, logger log.Logger) SessionMgmntRepository {
	return &sessionMgmntRepository{store: store, logger: logger}
}

// Create session is stored in-memory
func (s *sessionMgmntRepository) Create(sessionId string, expiration time.Time) error {
	if sessionId == "" {
		return ErrEmpty
	}

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

// Extend session id with the provided TTL
func (s *sessionMgmntRepository) Extend(request *ExtendRequest) (bool, error) {
	expiration := time.Now().Add(time.Second * time.Duration(request.TTL))
	obj, found, err := s.store.Reset(request.SessionId, expiration)
	if err != nil {
		return false, err
	}
	if found != true {
		return false, nil
	}
	if bytes.Equal(obj, []byte(request.SessionId)) == false {
		return false, ErrInvalidSessionId
	}
	return true, nil
}

// Exist if the session exists
func (s *sessionMgmntRepository) Exist(sessionId string) (bool, error) {
	b, found, err := s.store.Find(sessionId)
	if err != nil {
		return false, err
	}
	if found != true {
		return false, nil
	}
	if bytes.Equal(b, []byte(sessionId)) == false {
		return false, ErrInvalidSessionId
	}
	return true, nil
}

//List returns a list of all the sessions that the service is currently tracking
func (s *sessionMgmntRepository) List() (*Sessions, error) {
	session := &Sessions{}
	sessionMap, err := s.store.List()
	if err != nil {
		return nil, err
	}
	// put session map values into sessions list
	for _, value := range sessionMap {
		session.List = append(session.List, string(value.Oject))
	}
	return session, nil
}
