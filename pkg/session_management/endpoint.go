package session_management

import (
	"context"

	"github.com/go-kit/kit/endpoint"

	. "github.com/hecomp/session-management/internal/models"
	. "github.com/hecomp/session-management/pkg/http_response"
)

//type Set struct {
//	CreateEndpoint endpoint.Endpoint
//	DestroyEndpoint endpoint.Endpoint
//	ExtendEndpoint endpoint.Endpoint
//	ListEndpoint endpoint.Endpoint
//}

//// New returns a Set that wraps the provided server, and wires in all of the
//// expected endpoint middlewares via the various parameters.
//func New(svc SessionMgmntService, logger log.Logger, repository Repository) Set {
//	var createEndpoint endpoint.Endpoint
//	{
//		createEndpoint = makeCreateEndpoint(svc)
//	}
//	var destroyEndpoint endpoint.Endpoint
//	{
//		destroyEndpoint = makeDestroyEndpoint(svc)
//	}
//	var extendEndpoint endpoint.Endpoint
//	{
//		extendEndpoint = makeExtendEndpoint(svc)
//	}
//	var listEndpoint endpoint.Endpoint
//	{
//		listEndpoint = makeListEndpoint(svc)
//	}
//	return Set{
//		CreateEndpoint: createEndpoint,
//		DestroyEndpoint: destroyEndpoint,
//		ExtendEndpoint: extendEndpoint,
//		ListEndpoint: listEndpoint,
//	}
//}

//func (s Set) Create(ctx context.Context, session Session) (*SessionMgmntResponse, error) {
//	resp, err := s.CreateEndpoint(ctx, session)
//	if err != nil {
//		return nil, err
//	}
//	response := resp.(SessionMgmntResponse)
//	return &response, response.Err
//}
//
//func (s Set) Destroy(ctx context.Context, session DestroyRequest) (*SessionMgmntResponse, error) {
//	resp, err := s.DestroyEndpoint(ctx, session)
//	if err != nil {
//		return nil, err
//	}
//	response := resp.(SessionMgmntResponse)
//	return &response, response.Err
//}
//
//func (s Set) Extend(ctx context.Context, request ExtendRequest) (*SessionMgmntResponse, error) {
//	resp, err := s.ExtendEndpoint(ctx, request)
//	if err != nil {
//		return nil, err
//	}
//	response := resp.(SessionMgmntResponse)
//	return &response, response.Err
//}
//
//func (s Set) List(ctx context.Context) (*SessionMgmntResponse, error) {
//	type obj struct {}
//	resp, err := s.ListEndpoint(ctx, obj{})
//	if err != nil {
//		return nil, err
//	}
//	response := resp.(SessionMgmntResponse)
//	return &response, response.Err
//}

// makeCreateEndpoint create a session and return a unique session-id
func makeCreateEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(interface{}, error) {
		session := request.(Session)

		resp, erro := service.Create(&session)
		if erro != nil {
			return SessionMgmntResponse{ Message: resp.Message, Err: resp.Err }, nil
		}
		return SessionMgmntResponse{Message: resp.Message, Data: resp.Data }, nil
	}
}

// makeDestroyEndpoint remove the session from its cache
func makeDestroyEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(interface{}, error) {
		session := request.(DestroyRequest)

		resp, err := service.Destroy(&session)
		if err != nil {
			return SessionMgmntResponse{ Message: resp.Message, Err: err.Error() }, nil
		}
		return SessionMgmntResponse{Message: resp.Message, Data: resp.Data }, nil
	}
}

// makeExtendEndpoint extend TTL
func makeExtendEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(interface{}, error) {
		session := request.(ExtendRequest)

		resp, err := service.Extend(&session)
		if err != nil {
			return SessionMgmntResponse{ Message: resp.Message, Err: err.Error() }, nil
		}
		return SessionMgmntResponse{Message: resp.Message, Data: resp.Data }, nil
	}
}

// makeListEndpoint return a list of all the sessions that the sessionMgmntService is currently tracking
func makeListEndpoint(service SessionMgmntService) endpoint.Endpoint {
	return func(_ context.Context, request interface{})(response interface{}, err error) {
		resp, err := service.List()
		if err != nil {
			return SessionMgmntResponse{ Message: resp.Message, Err: err.Error() }, nil
		}
		return SessionMgmntResponse{Message: resp.Message, Data: resp.Data }, nil
	}
}


