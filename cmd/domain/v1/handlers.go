package v1

import (
	"fmt"
	"github.com/neelbhat88/go-api-template/internal/api"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	name := api.ReadQueryParam(r, "name")

	log.Info().Str("name", name).Msg("saying hi to someone")

	if name == "world" {
		api.Respond(r, w, http.StatusBadRequest, "Name cannot be 'world'")
		return
	}

	api.Respond(r, w, http.StatusOK, fmt.Sprintf("Hello, %s!", name))
}
