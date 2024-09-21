package handlers

import (
	"encoding/json"
	"net/http"
	"server/internal/database"

	"github.com/google/uuid"
)

type userRequestFile struct {
	Fileurl  string `json:"fileUrl"`
	Filetype string `json:"fileType"`
}

func (DBConfig *DBConfig) Addfile(w http.ResponseWriter, r *http.Request) {
	var user database.User = r.Context().Value("user").(database.User)
	print(user.Email)
	fileparm := userRequestFile{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&fileparm)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "invaild request")
		return
	}
	dbParams := database.AddFileUrlParams{
		Userid:   user.ID,
		Fileurl:  fileparm.Fileurl,
		Typefile: fileparm.Filetype,
		ID:       uuid.New(),
	}
	file, errer := DBConfig.DB.AddFileUrl(r.Context(), dbParams)
	if errer != nil {
		RespondWithError(w, http.StatusNotAcceptable, "not vaild")
	}
	RespondWithJson(w, http.StatusOK, file)
}
