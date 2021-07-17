package session_management

import (
	"errors"
	"github.com/go-kit/kit/log"

	. "github.com/hecomp/session-management/internal/models"
	. "github.com/hecomp/session-management/pkg/http_response"
	. "github.com/hecomp/session-management/pkg/repository"
)

// ErrInvalidArgument is returned when one or more arguments are invalid.
var ErrInvalidArgument = errors.New("invalid argument")

// SessionMgmntService is the interface that provides session management APIs.
type SessionMgmntService interface {
	Create(session *Session) (*SessionMgmntResponse, error)
	Destroy(session *DestroyRequest) (*SessionMgmntResponse, error)
	Extend(request *ExtendRequest) (*SessionMgmntResponse, error)
	List() (*SessionMgmntResponse, error)
}

//sessionMgmntService
type sessionMgmntService struct {
	logger log.Logger
	repo   SessionMgmntRepository
}

//NewService
func NewService(repo SessionMgmntRepository, logger log.Logger) SessionMgmntService {
	return &sessionMgmntService{repo: repo, logger: logger}
}

func (s sessionMgmntService) Create(session *Session) (*SessionMgmntResponse, error) {
	panic("implement me")
}

func (s sessionMgmntService) Destroy(session *DestroyRequest) (*SessionMgmntResponse, error) {
	panic("implement me")
}

func (s sessionMgmntService) Extend(request *ExtendRequest) (*SessionMgmntResponse, error) {
	panic("implement me")
}

func (s sessionMgmntService) List() (*SessionMgmntResponse, error) {
	panic("implement me")
}

