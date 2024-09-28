package types

import "server/internal/database"

type FileResponse struct {
	Fileurl  string `json:"fileUrl"`
	Typefile string `json:"fileType"`
	Filename string `json:"fileName"`
}

func ConvertFile(file database.File) FileResponse {

	return FileResponse{
		file.Fileurl,
		file.Typefile,
		file.Filename,
	}
}

func ConvertFileArray(files []database.File) []FileResponse {
	var Responsefiles []FileResponse

	for i := 0; i < len(files); i++ {
		file := FileResponse{
			files[i].Fileurl,
			files[i].Typefile,
			files[i].Filename,
		}
		Responsefiles = append(Responsefiles, file)
	}
	return Responsefiles
}
