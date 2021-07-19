package session_management

import (
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/google/uuid"
	"time"

	. "github.com/hecomp/session-management/internal/models"
	. "github.com/hecomp/session-management/pkg/repository"
)


const (
	DefaultTime = 30
	MaxTTL      = 300
)

var (
	// ErrInvalidArgument is returned when one or more arguments are invalid.
    ErrInvalidArgument    = errors.New("invalid argument")
 	ErrUserAlreadyExists  = fmt.Sprintf("user already exists with the given email")
	ErrDestroy            = errors.New("error destroying session id")
	ErrExtend             = errors.New("error extending session")
	ErrList               = errors.New("error listing session")
)


// SessionMgmntService is the interface that provides session management APIs.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SessionMgmntService
type SessionMgmntService interface {
	Create(session *SessionRequest) (string, error)
	Destroy(session *DestroyRequest) error
	Extend(request *ExtendRequest) error
	List() (*Sessions, error)
}

// sessionMgmntService has the implementation of the service methods
type sessionMgmntService struct {
	logger log.Logger
	repo   SessionMgmntRepository
}

// NewService create a instance of session management service
func NewService(repo SessionMgmntRepository, logger log.Logger) SessionMgmntService {
	return &sessionMgmntService{repo: repo, logger: logger}
}

// Create session is stored in-memory
func (s sessionMgmntService) Create(session *SessionRequest) (string, error) {
	if session.TTL == 0 {// default should be 30 seconds
		session.TTL = DefaultTime
	}

	sessionId := s.GenerateSessionId()
	expiration := time.Now().Add(time.Second * time.Duration(session.TTL))
	if err := s.repo.Create(sessionId, expiration); err != nil {
		s.logger.Log("message", "unable to create session to in-memory store", "error", err)
		return "", ErrEmpty
	}

	return sessionId, nil
}

// Destroy remove the session from its cache
func (s sessionMgmntService) Destroy(session *DestroyRequest) error {
	if session.SessionId == "" {
		return ErrEmpty
	}

	found, err := s.repo.Exist(session.SessionId)
	if err != nil {
		s.logger.Log("message", "unable to find session to in-memory store", "error", err)
		return ErrExist
	}
	if !found {
		s.logger.Log("message", "not found session to in-memory store", "error", ErrNotFound.Error())
		return ErrNotFound
	}

	if err := s.repo.Destroy(session); err != nil {
		s.logger.Log("message", "unable to destroy session in the in-memory store", "error", err)
		return ErrDestroy
	}
	return nil
}

// Extend with the provided TTL or if the TTL is not provided then by 30 seconds
func (s sessionMgmntService) Extend(request *ExtendRequest) error {
	if request.SessionId == "" {
		return ErrEmpty
	}

	if request.TTL == 0 {
		request.TTL = DefaultTime
	}

	if request.TTL > MaxTTL {
		request.TTL = MaxTTL
	}

	found, err := s.repo.Extend(request)
	if err != nil {
		s.logger.Log("message", "unable to extend session to in-memory store", "error", err)
		return ErrExtend
	}
	if !found {
		s.logger.Log("message", "not found session to in-memory store", "error", ErrNotFound.Error())
		return ErrNotFound
	}
	return nil
}

// List return a list of all the sessions that the service is currently tracking
func (s sessionMgmntService) List() (*Sessions, error) {
	sessions, err := s.repo.List()
	if err != nil {
		s.logger.Log("message", "unable to list sessions from in-memory store", "error", err)
		return nil, ErrList
	}

	if len(sessions.List) == 0 {
		return nil, ErrNotFound
	}
	return sessions, nil
}

// GenerateSessionId a unique session-id which should be UUID based
func (s *sessionMgmntService) GenerateSessionId() string {
	return uuid.Must(uuid.NewRandom()).String()
}

