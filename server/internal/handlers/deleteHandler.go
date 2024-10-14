package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type fileparams struct {
	FileId uuid.UUID `json:"fileid"`
}

func (DBConfig *DBConfig) DeleteFile(w http.ResponseWriter, r *http.Request) {
	fileparams := fileparams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&fileparams)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "invaild request")
		return
	}
	bucket, err := DBConfig.Filestore.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(bucket.Attrs(r.Context()))
	tx, err := DBConfig.Db.Begin()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "files does not exsit")
	}
	defer tx.Rollback()
	queries := DBConfig.DB.WithTx(tx)
	file, err := queries.GetFileByFileId(r.Context(), fileparams.FileId)
	if err != nil {
		log.Println("can't find file")
		RespondWithError(w, 401, "files does not exsit")
		// return
	}
	fmt.Println(file.Filename, fileparams.FileId)
	errs := bucket.Object(file.Filename).Delete(r.Context())

	if errs != nil {
		log.Println("can't find ")
		RespondWithError(w, 401, "files does not exsit")
		return
	}
	err = queries.DeleteByFileID(r.Context(), file.ID)
	// err = DBConfig.DB.DeleteByFileID(r.Context(), file.ID)
	if err != nil {

		log.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "files does not exsit")
		tx.Rollback()
		return
	}
	err = tx.Commit()
	if err != nil {
		log.Println(err)
		RespondWithError(w, http.StatusInternalServerError, "files does not exsit")
		tx.Rollback()
		return
	}
	RespondWithJson(w, 200, "file deleted")

}
