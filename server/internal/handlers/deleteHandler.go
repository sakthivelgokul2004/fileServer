package handlers

import (
	"fmt"
	"log"
	"net/http"
)

func (DBConfig *DBConfig) deleteFile(w http.Response, r *http.Request) {
	bucket, err := DBConfig.Filestore.DefaultBucket()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(bucket.Attrs(r.Context()))
}
