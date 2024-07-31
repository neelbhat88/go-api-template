package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Message   string `json:"message"`
	RequestID string `json:"requestID"`
}

// ReadQueryParam reads the query parameter from the http.Request
func ReadQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

// Respond writes the status and body to the http.ResponseWriter
//
// If the status is >= 400, the response body contains a standard response structure
func Respond(r *http.Request, w http.ResponseWriter, status int, returnObject interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if status >= 400 {
		errResp := ErrorResponse{
			RequestID: middleware.GetReqID(r.Context()),
		}
		if err, ok := returnObject.(error); ok {
			errResp.Message = err.Error()
		} else {
			jsonObj, err := json.Marshal(returnObject)
			if err != nil {
				log.Error().
					Err(err).
					Any("returnObject", returnObject).
					Msg("Failed to marshal returnObject")
			}
			errResp.Message = string(jsonObj)
		}

		// Write returnObject as marshaled JSON string to context
		jsonObj, err := json.Marshal(errResp)
		if err != nil {
			log.Error().
				Err(err).
				Any("returnObject", returnObject).
				Msg("Failed to marshal err response")
		}

		w.Write(jsonObj)
		return
	}

	json, _ := json.Marshal(returnObject)
	w.Write(json)
}

// DecodeBody reads the request Body and decodes it into the Object passed in
func DecodeBody(r *http.Request, v interface{}) error {
	err := decodeJSON(r.Body, v)
	if err != nil {
		log.Error().Msg("error decoding request body")
		return err
	}

	return nil
}

func decodeJSON(r io.Reader, v interface{}) error {
	defer io.Copy(io.Discard, r)
	return json.NewDecoder(r).Decode(v)
}
