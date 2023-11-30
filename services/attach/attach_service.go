package attach_service

import (
	"errors"
	"fmt"
	"gameComp-Backend/models"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gorm.io/gorm"
)

type AttachError struct {
	Message string
}

func (c AttachError) Error() string {
	return fmt.Sprintf("AttachError : %s", c.Message)
}

func GenerateFileName(teamUUID string, filename string, allowMulti bool) string {
	extension := filepath.Ext(filename)
	filename = strings.TrimSuffix(filename, extension)
	filename = strings.Replace(filename, " ", "_", -1)
	filename = strings.ToLower(filename)
	timestamp := time.Now().Unix()
	if allowMulti {
		filename = fmt.Sprintf("%s_%s_%d%s", teamUUID, filename, timestamp, extension)
	} else {
		filename = fmt.Sprintf("%s_Document%s", teamUUID, extension)
	}

	return filename
}

func GenerateDir(teamUUID string) string {
	return fmt.Sprintf("%s/%s", os.Getenv("FILE_STORAGE_PATH"), teamUUID)
}

func UploadFile(team *models.Team, attach *models.Attach, file *multipart.FileHeader, attachType string, allowMulti bool, allowsuffix []string) error {
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
	filename := GenerateFileName(team.UUID, file.Filename, allowMulti)

	// Save file
	saveLocation := os.Getenv("FILE_STORAGE_PATH")
	if saveLocation == "" {
		saveLocation = "."
	}
	if saveLocation[len(saveLocation)-1] == '/' {
		saveLocation = saveLocation[0 : len(saveLocation)-2]
	}
	saveLocation = fmt.Sprintf("%s/%s", saveLocation, team.UUID)
	attach.Path = fmt.Sprintf("%s/", team.UUID)
	attach.Filename = filename
	attach.SaveLocation = saveLocation
	attach.Type = attachType

	attach.Save()

	return nil
}

func GetAttachesWithConds(condi models.Attach) (*[]models.Attach, error) {
	dest, err := condi.FindMany()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return dest, nil
	}
	return dest, err
}
