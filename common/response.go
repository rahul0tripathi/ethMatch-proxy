package common

import "encoding/json"

const (
	INTERNAL_SERVER_ERROR = "INTERNAL_SERVER_ERROR"
)

type RequestResponse struct {
	StatusCode uint        `json:"statusCode"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func NewResponse(code uint, message string, data interface{}) RequestResponse {
	return RequestResponse{
		StatusCode: code,
		Message:    message,
		Data:       data,
	}
}
func (r RequestResponse) Marshal() (response []byte, err error) {
	response, err = json.Marshal(r)
	return
}
func (r RequestResponse) MarshalMust() []byte {
	response, err := json.Marshal(r)
	if err != nil {
		r.Data = struct {
		}{}
		r.StatusCode = 500
		r.Message = INTERNAL_SERVER_ERROR
	}
	response, _ = json.Marshal(r)
	return response
}
