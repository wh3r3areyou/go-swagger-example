package models

// APIResponse Api response
//
// swagger:model ApiResponse
type APIResponse struct {

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}
