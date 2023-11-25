package attach_service

import (
	"fmt"
	"gameComp-Backend/models"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type AttachError struct {
	Message string
}

func (c AttachError) Error() string {
	return fmt.Sprintf("AttachError : %s", c.Message)
}

func GenerateFileName(owner string, filename string, allowMulti bool) string {
	extension := filepath.Ext(filename)
	filename = strings.TrimSuffix(filename, extension)
	filename = strings.Replace(filename, " ", "_", -1)
	filename = strings.ToLower(filename)
	timestamp := time.Now().Unix()
	if allowMulti {
		filename = fmt.Sprintf("%s_%d%s", filename, timestamp, extension)
	} else {
		filename = fmt.Sprintf("%s_%s", owner, extension)
	}

	return filename
}

func UploadFile(owner *models.User, attach *models.Attach, file *multipart.FileHeader, attachType string, allowMulti bool, allowsuffix []string) error {
	// Check if the file type is allowed
	isAllowed := false
	for _, s := range allowsuffix {
		if filepath.Ext(file.Filename) == s {
			isAllowed = true
			break
		}
	}
	if !isAllowed {
		return AttachError{
			Message: "file type error",
		}
	}

	// Generate file name
	filename := GenerateFileName(owner.TeamName, file.Filename, allowMulti)

	// Save file
	saveLocation := os.Getenv("FILE_STORAGE_PATH")
	if saveLocation == "" {
		saveLocation = "."
	}
	if saveLocation[len(saveLocation)-1] == '/' {
		saveLocation = saveLocation[0 : len(saveLocation)-2]
	}

	attach.Filename = filename
	attach.SaveLocation = saveLocation
	attach.Type = attachType

	attach.Save()

	return nil
}

func GetAttachesWithConds(condi models.Attach) *[]models.Attach {
	dest := condi.FindMany()
	return dest
}
