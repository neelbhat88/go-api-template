package apimiddleware

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"runtime/debug"
)

func Recoverer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// try to cast err to error
				errStr, ok := err.(error)
				if !ok {
					errStr = fmt.Errorf("%v", err)
				}
				log.Error().
					Str("stack", string(debug.Stack())).
					Err(errStr).
					Msg("Recovered from panic")

				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
