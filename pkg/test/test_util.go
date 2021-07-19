package test

import (
	"os"

	"github.com/go-kit/kit/log"

	. "github.com/hecomp/session-management/internal/models"
)

// GetLogger create a single logger, which we'll use and give to other components.
func GetLogger() log.Logger {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	return logger
}

// ConvertMapToList
func ConvertMapToList(sessionMap map[string]Item) *Sessions {
	session := &Sessions{}

	for _, value := range sessionMap {
		session.List = append(session.List, string(value.Oject))
	}
	return session
}
