package session_management

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/hecomp/session-management/internal/models"
)

//loggingService
type loggingService struct {
	logger log.Logger
	SessionMgmntService
}

//NewLoggingService
func NewLoggingService(logger log.Logger, s SessionMgmntService) SessionMgmntService {
	return &loggingService{logger: logger, SessionMgmntService: s}
}

// Create
func (s *loggingService) Create(session *models.SessionRequest) (sig *SessionMgmntResponse, err error)  {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "create",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SessionMgmntService.Create(session)
}

//Destroy
func (s *loggingService) Destroy(session *models.DestroyRequest) (sig *SessionMgmntResponse, err error)  {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "destroy",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SessionMgmntService.Destroy(session)
}

//Extend
func (s *loggingService) Extend(session *models.ExtendRequest) (sig *SessionMgmntResponse, err error)  {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "extend",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SessionMgmntService.Extend(session)
}

//List
func (s *loggingService) List() (sig *SessionMgmntResponse, err error)  {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "list",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SessionMgmntService.List()
}

