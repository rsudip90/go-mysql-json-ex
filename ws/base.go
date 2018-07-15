package ws

import (
	"encoding/json"
	"net/http"
)

// ErrResp for error message
type ErrResp struct {
	Error string `json:"error"`
}

// SuccessResponse returns success response in json
func SuccessResponse(w http.ResponseWriter, g interface{}) {
	var (
		b   []byte
		err error
	)

	b, err = json.Marshal(g)
	if err != nil {
		ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

// ErrorResponse returns error response in json
func ErrorResponse(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	e := ErrResp{Error: errMsg}
	b, _ := json.Marshal(&e)
	w.Write(b)
}
