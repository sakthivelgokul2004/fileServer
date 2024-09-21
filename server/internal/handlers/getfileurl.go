package handlers

import (
	"fmt"
	"net/http"
)

func (DBConfig *DBConfig) GetFileUrl(w http.ResponseWriter, r *http.Request) {

	// // // var user database.User = r.Context().Value("user").(database.User)

	// // files, err := DBConfig.DB.GetFileByUserId(r.Context(), user.ID)
	// if err != nil {
	// 	RespondWithError(w, http.StatusBadRequest, "invaild request")
	// 	return
	// }
	// fmt.Println(len(files))
	// RespondWithJson(w, http.StatusAccepted, len(files))
	fmt.Println(r.Cookies())

}
