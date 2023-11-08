package models

import "mime/multipart"

type Attach struct {
	Owner        *User
	SaveLocation string
	Filename     string
	FileSrc      *multipart.FileHeader
}
