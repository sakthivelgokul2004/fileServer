package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/internal/auth"
	"server/internal/database"

	"github.com/google/uuid"
)

type userRequestFile struct {
	Fileurl  string `json:"fileurl"`
	Filetype string `json:"filetype"`
	Filename string `json:"filename"`
}

func (DBConfig *DBConfig) Addfile(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(auth.UserContextKey).(database.User)
	if !ok {
		log.Println("not vaild")
	}
	print(user.Email)
	fileparm := userRequestFile{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&fileparm)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "invaild request")
		return
	}
	fmt.Println(fileparm)
	dbParams := database.AddFileUrlParams{
		Userid:   user.ID,
		Filename: fileparm.Filename,
		Fileurl:  fileparm.Fileurl,
		Typefile: fileparm.Filetype,
		ID:       uuid.New(),
	}
	fmt.Println(dbParams)
	file, errer := DBConfig.DB.AddFileUrl(r.Context(), dbParams)
	fmt.Println(file)
	if errer != nil {
		RespondWithError(w, http.StatusNotAcceptable, "not vaild")
	}
	RespondWithJson(w, http.StatusOK, file)
}
