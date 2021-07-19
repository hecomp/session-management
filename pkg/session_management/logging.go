package session_management

import (
	"time"

	"github.com/go-kit/kit/log"

	"github.com/hecomp/session-management/internal/models"
)

//loggingService has the implementation of the logging middleware methods.
type loggingService struct {
	logger log.Logger
	SessionMgmntService
}

//NewLoggingService create a instance of logging service
func NewLoggingService(logger log.Logger, s SessionMgmntService) SessionMgmntService {
	return &loggingService{logger: logger, SessionMgmntService: s}
}

// Create session is stored in-memory
func (s *loggingService) Create(session *models.SessionRequest) (sig string, err error)  {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "create",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SessionMgmntService.Create(session)
}

//Destroy remove the session from its cache
func (s *loggingService) Destroy(session *models.DestroyRequest) (err error)  {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "destroy",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SessionMgmntService.Destroy(session)
}

//Extend session id with the provided TTL
func (s *loggingService) Extend(session *models.ExtendRequest) (err error)  {
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
func (s *loggingService) List() (sig *models.Sessions, err error)  {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "list",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.SessionMgmntService.List()
}

