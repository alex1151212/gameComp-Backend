package attach_service

import (
	"fmt"
	"gameComp-Backend/models"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type AttachError struct {
	Message string
}

func (c AttachError) Error() string {
	return fmt.Sprintf("AttachError : %s", c.Message)
}

func GenerateFileName(owner string, filename string) string {
	extension := filepath.Ext(filename)
	filename = strings.TrimSuffix(filename, extension)
	filename = strings.Replace(filename, " ", "_", -1)
	filename = strings.ToLower(filename)
	filename = fmt.Sprintf("%s%s", owner, extension)
	return filename
}

func UploadFile(attach *models.Attach, files []*multipart.FileHeader, allowsuffix []string) error {

	for _, file := range files {

		for _, s := range allowsuffix {
			if filepath.Ext(file.Filename) == s {
				break
			}
			return AttachError{
				Message: "file type error",
			}
		}

		//產生檔案名稱
		filename := GenerateFileName(attach.Owner.Username, file.Filename)

		//儲存檔案
		saveLocation := os.Getenv("FILE_STORAGE_PATH")
		if saveLocation == "" {
			saveLocation = "."
		}
		if saveLocation[len(saveLocation)-1] == '/' {
			saveLocation = saveLocation[0 : len(saveLocation)-2]
		}

		attach.Filename = filename
		attach.SaveLocation = saveLocation
		attach.FileSrc = file
	}
	return nil
}

func SaveFile(user *models.User, attach *models.Attach) error {
	user.PdfPath = attach.Filename
	user.Save()
	return nil
}
