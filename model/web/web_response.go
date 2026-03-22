package web

type WebResponse struct {
	Success   bool   `json:"success"`
	Code      int    `json:"code,omitempty"`
	Status    string `json:"status,omitempty"`
	Message   string `json:"message,omitempty"`
	RequestID string `json:"requestId,omitempty"`
	Data      any    `json:"data,omitempty"`
	Errors    any    `json:"errors,omitempty"`
}
