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
 	ErrUserAlreadyExists  = fmt.Sprintf("User already exists with the given email")
	ErrDestroy            = fmt.Sprintf("Error destroying session id")
	ErrExtend             = fmt.Sprintf("Error extending session")
	ErrList              = fmt.Sprintf("Error listing session")

    CreateSessionSuccess  = fmt.Sprintf("Session Created Successfully")
	DestroySessionSuccess = fmt.Sprintf("Session Destroyed Successfully")
	ExtendSessionSuccess = fmt.Sprintf("Session Extended Successfully")
	ListSessionSuccess = fmt.Sprintf("Session Listed Successfully")
)


// SessionMgmntService is the interface that provides session management APIs.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . SessionMgmntService
type SessionMgmntService interface {
	Create(session *SessionRequest) (*SessionMgmntResponse, error)
	Destroy(session *DestroyRequest) (*SessionMgmntResponse, error)
	Extend(request *ExtendRequest) (*SessionMgmntResponse, error)
	List() (*SessionMgmntResponse, error)
}

// sessionMgmntService
type sessionMgmntService struct {
	logger log.Logger
	repo   SessionMgmntRepository
}

// NewService
func NewService(repo SessionMgmntRepository, logger log.Logger) SessionMgmntService {
	return &sessionMgmntService{repo: repo, logger: logger}
}

// Create
func (s sessionMgmntService) Create(session *SessionRequest) (*SessionMgmntResponse, error) {
	if session.TTL == 0 {// default should be 30 seconds
		session.TTL = DefaultTime
	}

	sessionId := s.GenerateSessionId()
	expiration := time.Now().Add(time.Second * time.Duration(session.TTL))
	if err := s.repo.Create(sessionId, expiration); err != nil {
		s.logger.Log("message", "unable to create session to in-memory store", "error", err)
		return &SessionMgmntResponse{ Message: ErrEmpty.Error(), Err: err}, err
	}

	return &SessionMgmntResponse{ Message: CreateSessionSuccess, Data: &Session{ SessionId: sessionId } }, nil
}

// Destroy
func (s sessionMgmntService) Destroy(session *DestroyRequest) (*SessionMgmntResponse, error) {
	if session.SessionId == "" {
		return &SessionMgmntResponse{ Message: ErrEmpty.Error() }, nil
	}

	if err := s.repo.Destroy(session); err != nil {
		s.logger.Log("message", "unable to destroy session in the in-memory store", "error", err)
		return &SessionMgmntResponse{ Message: ErrDestroy, Err: err}, err
	}
	return &SessionMgmntResponse{ Message: DestroySessionSuccess }, nil
}

// Extend with the provided TTL or if the TTL is not provided then by 30 seconds
func (s sessionMgmntService) Extend(request *ExtendRequest) (*SessionMgmntResponse, error) {
	if request.SessionId == "" {
		return &SessionMgmntResponse{ Message: ErrEmpty.Error() }, nil
	}

	if request.TTL == 0 {
		request.TTL = DefaultTime
	}

	if request.TTL > MaxTTL {
		request.TTL = MaxTTL
	}

	found, err := s.repo.Exist(request.SessionId)
	if err != nil {
		s.logger.Log("message", "unable to find session to in-memory store", "error", err)
		return &SessionMgmntResponse{ Message: ErrEmpty.Error(), Err: err }, err
	}
	if !found {
		s.logger.Log("message", "not found session to in-memory store", "error", err)
		return &SessionMgmntResponse{ Message: ErrNotFound.Error() }, ErrNotFound
	}

	if err := s.repo.Extend(request); err != nil {
		s.logger.Log("message", "unable to extend session to in-memory store", "error", err)
		return &SessionMgmntResponse{ Message: ErrExtend }, errors.New(ErrExtend)
	}
	return &SessionMgmntResponse{ Message: ExtendSessionSuccess }, nil
}

// List
func (s sessionMgmntService) List() (*SessionMgmntResponse, error) {
	sessions, err := s.repo.List()
	if err != nil {
		s.logger.Log("message", "unable to list sessions from in-memory store", "error", err)
		return &SessionMgmntResponse{ Message: ErrList }, errors.New(ErrList)
	}
	return &SessionMgmntResponse{ Message: ListSessionSuccess, Data: sessions }, nil
}

// GenerateSessionId
func (s *sessionMgmntService) GenerateSessionId() string {
	return uuid.Must(uuid.NewRandom()).String()
}

