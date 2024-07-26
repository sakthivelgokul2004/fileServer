package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func ensureTmpDirExists() error {
	_, err := os.Stat("./images")
	if os.IsNotExist(err) {
		err := os.Mkdir("./images", 0755) // Create 'tmp' directory with read/write/execute permissions for owner and read/execute permissions for group and others
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}
func Upload(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data with a max limit of 200 MB
	err := r.ParseMultipartForm(200 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		log.Error(err)
		return
	}
	ensureTmpDirExists()
	// Access the uploaded file
	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving file from form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Print out details for debugging
   fmt.Println("Handler:", handler.Filename)
   fmt.Println("Size:", handler.Size)
   fmt.Println("MIME:", handler.Header.Get("Content-Type"))

	// You can save the file to disk or process it as needed
	// Example: Saving the file to a local directory
	destination, err := os.Create("./images/" + handler.Filename)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		log.Error(err)
		return
	}
	defer destination.Close()
	_, err = io.Copy(destination, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		log.Error(err)
		return
	}

	// Respond with OK status
	w.WriteHeader(http.StatusOK)
}
