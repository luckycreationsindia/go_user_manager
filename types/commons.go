package types

type ResponseMessage struct {
	Status   int    `json:"status"`
	Message  string `json:"message,omitempty"`
	Data     any    `json:"data,omitempty"`
	Redirect string `json:"redirect,omitempty"`
}
