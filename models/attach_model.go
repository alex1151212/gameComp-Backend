package models

import "mime/multipart"

type Attach struct {
	SaveLocation string
	Filename     string
	FileSrc      *multipart.FileHeader
}
