package session_management

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	. "github.com/hecomp/session-management/internal/models"
)

// SessionMgmntResponse collects the response values for the Create API.
type SessionMgmntResponse struct {
	Message   string   `json:",omitempty"`
	Data      interface{} `json:"data"`
	Err       error `json:"err,omitempty"` // should be intercepted by Failed/errorEncoder
}

// Failed implements endpoint.Failer.
func (r SessionMgmntResponse) Failed() error { return r.Err }

// MakeCreateEndpoint create a session and return a unique session-id
func MakeCreateEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(interface{}, error) {
		session := request.(SessionRequest)

		resp, erro := service.Create(&session)
		if erro != nil {
			return &SessionMgmntResponse{ Message: resp.Message, Err: erro }, nil
		}
		return &SessionMgmntResponse{Message: resp.Message, Data: resp.Data }, nil
	}
}

// MakeDestroyEndpoint remove the session from its cache
func MakeDestroyEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(interface{}, error) {
		session := request.(DestroyRequest)

		resp, err := service.Destroy(&session)
		if err != nil {
			return SessionMgmntResponse{ Message: resp.Message, Err: resp.Err }, nil
		}
		return SessionMgmntResponse{Message: resp.Message, Data: resp.Data }, nil
	}
}

// MakeExtendEndpoint extend TTL
func MakeExtendEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(interface{}, error) {
		session := request.(ExtendRequest)

		resp, err := service.Extend(&session)
		if err != nil {
			return SessionMgmntResponse{ Message: resp.Message, Err: resp.Err }, nil
		}
		return SessionMgmntResponse{Message: resp.Message, Data: resp.Data }, nil
	}
}

// MakeListEndpoint return a list of all the sessions that the sessionMgmntService is currently tracking
func MakeListEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(interface{},  error) {
		resp, err := service.List()
		if err != nil {
			return SessionMgmntResponse{ Message: resp.Message, Err: resp.Err }, nil
		}
		return SessionMgmntResponse{Message: resp.Message, Data: resp.Data }, nil
	}
}


