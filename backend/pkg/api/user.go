package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"un-defined/pkg/model"
	"un-defined/pkg/service"
)

// UserAPI ...
type UserAPI struct {
	UserService service.UserService
}

// NewUserAPI ...
func NewUserAPI(u service.UserService) UserAPI {
	return UserAPI{UserService: u}
}

// FindAllUsers ...
func (u UserAPI) FindAllUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users, err := u.UserService.All()
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, users)
	}
}

// FindByID ...
func (u UserAPI) FindByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Check if id is uuid
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}

		// Find login by id from db
		user, err := u.UserService.FindByID(uint(id))
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToUserDTO(user))
	}
}

// Save ...
func (u UserAPI) CreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var userDTO model.UserDTO

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&userDTO); err != nil {
			RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		defer r.Body.Close()

		createdUser, err := u.UserService.Save(model.ToUser(&userDTO))
		if err != nil {
			RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, model.ToUserDTO(createdUser))
	}
}

// Migrate ...
func (u UserAPI) Migrate() {
	err := u.UserService.Migrate()
	if err != nil {
		log.Println(err)
	}

}
