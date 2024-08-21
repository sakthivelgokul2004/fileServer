package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"server/internal/database"
)

type DBConfig struct {
	DB *database.Queries
}

type userRequestParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (DBconfig *DBConfig) SignupHandler(w http.ResponseWriter, r *http.Request) {
	usr := userRequestParam{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&usr)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "invaild request")
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 10)
	if err != nil {
		RespondWithError(w, 500, "can't genrate hash")
   }

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Email:     usr.Email,
		Password:  string(hash),
	}

	result, err := DBconfig.DB.CreateUser(r.Context(), params)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "email is already exiest ")
		return
	}
	w.WriteHeader(http.StatusOK)
	RespondWithJson(w, http.StatusOK, result)
}
