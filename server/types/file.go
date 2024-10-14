package types

import (
	"server/internal/database"

	"github.com/google/uuid"
)

type FileResponse struct {
	Fileurl  string    `json:"fileUrl"`
	Typefile string    `json:"fileType"`
	Filename string    `json:"fileName"`
	ID       uuid.UUID `json:"id"`
}

func ConvertFile(file database.File) FileResponse {

	return FileResponse{
		file.Fileurl,
		file.Typefile,
		file.Filename,
		file.ID,
	}
}

func ConvertFileArray(files []database.File) []FileResponse {
	var Responsefiles []FileResponse

	for i := 0; i < len(files); i++ {
		file := FileResponse{
			files[i].Fileurl,
			files[i].Typefile,
			files[i].Filename,
			files[i].ID,
		}
		Responsefiles = append(Responsefiles, file)
	}
	return Responsefiles
}
