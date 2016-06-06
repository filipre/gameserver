package main

import (
	"encoding/json"
	"net/http"

	"./player"

	"github.com/gocraft/web"
)

type responseJSON struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Result  interface{} `json:"result,omitempty"`
}

func getErrCode(err error) int {
	errCodes := map[error]int{
		errGetID:   http.StatusBadRequest,
		errGetJSON: http.StatusBadRequest,

		player.ErrDBExistName:  http.StatusInternalServerError,
		player.ErrDBInsert:     http.StatusInternalServerError,
		player.ErrDBUpdate:     http.StatusInternalServerError,
		player.ErrDBFind:       http.StatusInternalServerError,
		player.ErrPlayerExists: http.StatusBadRequest,
		player.ErrNotFound:     http.StatusNotFound,
		player.ErrBadName:      http.StatusBadRequest,
	}
	code, ok := errCodes[err]
	if !ok {
		//log...
		return http.StatusNotImplemented
	}
	return code
}

func responseError(rw web.ResponseWriter, err error) {
	response(rw, err.Error(), getErrCode(err), nil)
}

func response(rw web.ResponseWriter, message string, code int, result interface{}) {
	responseJSON, err := json.Marshal(&responseJSON{
		Message: message,
		Code:    code,
		Result:  result,
	})
	if err != nil {
		http.Error(rw, "Can't marshal Object. Please contact Admin.", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(code)
	rw.Write(responseJSON)
}
