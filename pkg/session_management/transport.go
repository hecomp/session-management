package session_management

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	kitlog "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	. "github.com/hecomp/session-management/internal/models"
	. "github.com/hecomp/session-management/pkg/repository"
)

var (
	// ErrBadRouting is used when a client sends a bad routing.
	ErrBadRouting = errors.New("Error bad routing")
	// ErrNotFound is used when a client could not be found.
	ErrNotFound   = errors.New("Asset not found\n")
	// ErrBadRequest is used when a client send a bad request.
	ErrBadRequest = errors.New("Bad Request")
	// ErrUnknown is used when a client is unknown.
	ErrUnknown = errors.New("unknown session")
)

// MakeHandler
func MakeHandler(svc SessionMgmntService, repo SessionMgmntRepository, logger kitlog.Logger) http.Handler {

	mux := http.NewServeMux()

	createHandler := httptransport.NewServer(
			makeCreateEndpoint(svc),
		decodeHTTPCreateRequest,
			encodeResponse)
	destroyHandler := httptransport.NewServer(
		makeDestroyEndpoint(svc),
		decodeHTTPCreateRequest,
		encodeResponse)
	extendHandler := httptransport.NewServer(
		makeExtendEndpoint(svc),
		decodeHTTPCreateRequest,
		encodeResponse)
	listHandler := httptransport.NewServer(
		makeListEndpoint(svc),
		decodeHTTPCreateRequest,
		encodeResponse)

	mux.Handle("/create", createHandler)
	mux.Handle("/destroy", destroyHandler)
	mux.Handle("/extend", extendHandler)
	mux.Handle("/list", listHandler)

	http.Handle("/", accessControl(mux))

	return mux
}

// decodeHTTPCreateRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded signup request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var user Session

	if r.Body == nil {
		return nil, ErrBadRequest
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, errors.New(err.Error())
	} else {
		return user, nil
	}
}

// decodeHTTPDestroyRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded signup request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPDestroyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var user Session

	if r.Body == nil {
		return nil, ErrBadRequest
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, errors.New(err.Error())
	} else {
		return user, nil
	}
}

// decodeHTTPExtendRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded signup request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPExtendRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var user Session

	if r.Body == nil {
		return nil, ErrBadRequest
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, errors.New(err.Error())
	} else {
		return user, nil
	}
}

// decodeHTTPListRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded signup request from the HTTP request body. Primarily useful in a
// server.
func decodeHTTPListRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var user Session

	if r.Body == nil {
		return nil, ErrBadRequest
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return nil, errors.New(err.Error())
	} else {
		return user, nil
	}
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}
