package session_management

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/endpoint"

	. "github.com/hecomp/session-management/internal/models"
	. "github.com/hecomp/session-management/pkg/repository"

)

var (
	ErrCreate             = errors.New("error creating session id")
	CreateSessionSuccess  = fmt.Sprintf("session created successfully")
	DestroySessionSuccess = fmt.Sprintf("session destroyed successfully")
	ExtendSessionSuccess  = fmt.Sprintf("session extended successfully")
	ListSessionSuccess    = fmt.Sprintf("session listed successfully")
)

// SessionMgmntResponse collects the response values for the Create API.
type SessionMgmntResponse struct {
	Message   string   `json:",omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Err       error `json:"err,omitempty"` // should be intercepted by Failed/errorEncoder
	StatusCode int `json:"status_code"`
}

// Failed implements endpoint.Failer.
func (r SessionMgmntResponse) Failed() error { return r.Err }

// MakeCreateEndpoint create a session and return a unique session-id
func MakeCreateEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(interface{}, error) {
		session := request.(SessionRequest)

		uuid, err := service.Create(&session)
		if err != nil {
			return &SessionMgmntResponse{ Message: ErrCreate.Error(), Err: ErrCreate, StatusCode: http.StatusInternalServerError  }, nil
		}
		return &SessionMgmntResponse{Message: CreateSessionSuccess, Data: &Session{
			SessionId: uuid,
		}, StatusCode: http.StatusCreated }, nil
	}
}

// MakeDestroyEndpoint remove the session from its cache
func MakeDestroyEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(interface{}, error) {
		session := request.(DestroyRequest)

		err := service.Destroy(&session)
		if err != nil {
			return &SessionMgmntResponse{ Message: err.Error(), Err: err, StatusCode: getStatusCode(err)  }, nil
		}
		return &SessionMgmntResponse{Message: DestroySessionSuccess, StatusCode: http.StatusOK }, nil
	}
}

// MakeExtendEndpoint extend TTL
func MakeExtendEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(interface{}, error) {
		session := request.(ExtendRequest)

		err := service.Extend(&session)
		if err != nil {
			return &SessionMgmntResponse{ Message: err.Error(), Err: err, StatusCode: getStatusCode(err)}, nil
		}
		return &SessionMgmntResponse{Message: ExtendSessionSuccess, StatusCode: http.StatusOK}, nil
	}
}

// MakeListEndpoint return a list of all the sessions that the sessionMgmntService is currently tracking
func MakeListEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(interface{},  error) {
		sessions, err := service.List()
		if err != nil {
			return &SessionMgmntResponse{ Message: err.Error(), Err: err, StatusCode: http.StatusInternalServerError }, nil
		}
		return &SessionMgmntResponse{Message: ListSessionSuccess, Data: sessions, StatusCode: http.StatusOK}, nil
	}
}

func getStatusCode(err error) int {
	var statusCode int

	switch err {
	case ErrNotFound:
		statusCode = http.StatusNotFound
	case ErrInvalidArgument:
		statusCode = http.StatusBadRequest
	default:
		statusCode = http.StatusInternalServerError
	}

	return statusCode
}


