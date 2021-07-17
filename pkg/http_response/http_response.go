package http_response

// SessionMgmntResponse collects the response values for the Create API.
type SessionMgmntResponse struct {
	Message   string   `json:",omitempty"`
	Data      interface{} `json:"data"`
	Err       string `json:"err,omitempty"` // should be intercepted by Failed/errorEncoder
}