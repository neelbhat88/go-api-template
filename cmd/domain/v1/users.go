package v1

import (
	"fmt"
	"github.com/neelbhat88/go-api-template/internal/api"
	"github.com/neelbhat88/go-api-template/internal/service/usersAdmin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h Handler) Root(w http.ResponseWriter, r *http.Request) {
	name := api.ReadQueryParam(r, "name")

	log.Info().Str("name", name).Msg("saying hi to someone")

	if name == "world" {
		api.Respond(r, w, http.StatusBadRequest, "Name cannot be 'world'")
		return
	}

	api.Respond(r, w, http.StatusOK, fmt.Sprintf("Hello, %s!", name))
}

func (h Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := usersAdmin.LoadAllUsers(h.DB)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get users")
		api.Respond(r, w, http.StatusInternalServerError, "Failed to get users")
		return
	}

	api.Respond(r, w, http.StatusOK, users)
}
