package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"server/internal/auth"
)

type userResponse struct {
	UserEmail string `json:"Email"`
	Time      int64  `json:"time"`
}

func (DBConfig *DBConfig) LoginHandler(w http.ResponseWriter, r *http.Request) {
	usrParam := userRequestParam{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&usrParam)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "invaild request")
		return
	}

	user, err := DBConfig.DB.GetUserByEmail(r.Context(), usrParam.Email)
	if err != nil {
		log.Println("can't find user")
		RespondWithError(w, 401, "does not exsit")
	}
	fmt.Println(user.Password, usrParam.Password)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(usrParam.Password))
	if err != nil {
		RespondWithError(w, 401, "incorrect password")
	}
	jwtToken, err := auth.GenJwt(user.ID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, " can't genrate token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    jwtToken,
		Path:     "/",
		Domain:   "",
		MaxAge:   3600 * 24 * 30,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	reponse := userResponse{
		UserEmail: user.Email,
		Time:      time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	// to create response with userEmail and userId
	RespondWithJson(w, http.StatusCreated, reponse)
}
