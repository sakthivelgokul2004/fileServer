package handlers

import (
	"fmt"
	"log"
	"net/http"
	"server/internal/auth"
	"server/internal/database"
	"server/types"
)

func (DBConfig *DBConfig) GetFileUrl(w http.ResponseWriter, r *http.Request) {

	var user, ok = r.Context().Value(auth.UserContextKey).(database.User)
	if !ok {
		log.Println("not vaild")
	}

	files, err := DBConfig.DB.GetFileByUserId(r.Context(), user.ID)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "invaild request")
		return
	}
	responseFiles := types.ConvertFileArray(files)
	fmt.Println(responseFiles)
	fmt.Println(files)
	RespondWithJson(w, http.StatusAccepted, responseFiles)

}
