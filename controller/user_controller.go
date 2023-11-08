package controller

import (
	"fmt"
	"gameComp-Backend/models"
	attach_service "gameComp-Backend/services/attach"
	user_service "gameComp-Backend/services/user"
	"gameComp-Backend/utils"

	"github.com/gofiber/fiber/v2"
)

type RegisterSuccessRes struct {
	Token string `json:"token"`
}

func GetUsers(c *fiber.Ctx) error {

	users := user_service.GetUsers()

	return utils.RespOK(c, users, "Get Users Success")

}

func UploadGameFile(c *fiber.Ctx) error {
	user := c.Locals("Auth").(*models.User)
	user.FindOne()

	videoLink := c.FormValue("videoLink", "")
	if videoLink == "" {
		return utils.RespFail(c, "Upload VideoLink Fail")
	}

	form, err := c.MultipartForm()
	if err != nil {
		return utils.RespFail(c, "Upload PDF Fail")
	}

	files := form.File["pdf"]
	if len(files) != 1 {
		return utils.RespFail(c, "files count error")
	}

	allowsuffix := []string{".pdf"}

	attach := models.Attach{
		Owner: user,
	}

	attach_service.UploadFile(&attach, files, allowsuffix)

	err = c.SaveFile(attach.FileSrc, fmt.Sprintf("%s/%s", attach.SaveLocation, attach.Filename))
	if err != nil {
		return utils.RespFail(c, "Upload PDF Fail")
	}

	//生成檔案連結
	url := utils.GetURL()
	url = url + attach.Filename

	user.PdfPath = url
	user.VideoLink = videoLink
	user.Save()

	return utils.RespOK(c, user, "Upload PDF Success")
}

func UploadIDPhoto(c *fiber.Ctx) error {
	user := c.Locals("Auth").(*models.User)
	user.FindOne()

	form, err := c.MultipartForm()
	if err != nil {
		return utils.RespFail(c, "Upload ID Photo Fail")
	}

	files := form.File["IdPhoto"]
	if len(files) != 1 {
		return utils.RespFail(c, "files count error")
	}

	allowsuffix := []string{".png", ".jpg", ".jpeg"}

	attach := models.Attach{
		Owner: user,
	}

	attach_service.UploadFile(&attach, files, allowsuffix)

	err = c.SaveFile(attach.FileSrc, fmt.Sprintf("%s/%s", attach.SaveLocation, attach.Filename))
	if err != nil {
		return utils.RespFail(c, "Upload ID Photo Fail")
	}

	//生成檔案連結
	url := utils.GetURL()
	fmt.Println(">>>>>>>>>>", url)
	url = url + attach.Filename

	user.Save()

	return utils.RespOK(c, user, "Upload ID Photo Success")
}
