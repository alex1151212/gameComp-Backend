package controller

import (
	"encoding/json"
	"fmt"
	"gameComp-Backend/models"
	attach_service "gameComp-Backend/services/attach"
	user_service "gameComp-Backend/services/user"
	"gameComp-Backend/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RegisterSuccessRes struct {
	Token string `json:"token"`
}

type ProfileRes struct {
	Email string `json:"email"`
	Phone string `json:"phone"`

	TeamName              string               `json:"teamName"`
	TeamMember            []models.TeamMember  `json:"teamMember"`
	TeamTeacher           []models.TeamTeacher `json:"teamTeacher"`
	TeamSchoolCertificate []string             `json:"teamSchoolCertificate"`

	WorkVideoLink string `json:"workVideoLink"`
	WorkPdf       string `json:"workPdf"`

	IsUpload    bool `json:"isUpload"`
	IsApplyTeam bool `json:"isApplyTeam"`
}

func GetUsers(c *fiber.Ctx) error {

	users := user_service.GetUsers()

	return utils.RespOK(c, users, "Get Users Success")
}

func UploadGameFile(c *fiber.Ctx) error {
	user := c.Locals("Auth").(*models.User)
	user.FindOne()

	team := user_service.GetUserTeam(user)

	videoLink := c.FormValue("workVideoLink", "")
	if videoLink == "" {
		return utils.RespFail(c, "Upload VideoLink Fail")
	}

	form, err := c.MultipartForm()
	if err != nil {
		return utils.RespFail(c, "Upload PDF Fail")
	}

	files := form.File["workPdf"]
	if len(files) != 1 {
		return utils.RespFail(c, "files count error")
	}

	allowsuffix := []string{".pdf"}

	attachPdf := models.Attach{UserID: user.ID}

	if len(files) > 0 {
		attach_service.UploadFile(&team, &attachPdf, files[0], "workPdf", false, allowsuffix)

		err = c.SaveFile(files[0], fmt.Sprintf("%s/%s", attachPdf.SaveLocation, attachPdf.Filename))
		if err != nil {
			return utils.RespFail(c, "Upload PDF Fail")
		}

		//生成檔案連結
		url := utils.GetURL()
		url = url + attachPdf.Filename

		team.WorkPdf = attachPdf
		team.WorkVideoLink = videoLink
		team.IsUpload = true
	}

	err = team.Update().Error
	if err != nil {
		return utils.RespFail(c, "something wrong in server")
	}

	return utils.RespOK(c, team, "Upload PDF Success")
}

func UserApply(ctx *fiber.Ctx) error {
	user := ctx.Locals("Auth").(*models.User)

	//the team is not exist
	if user.Team.ID == 0 {
		teamName := ctx.FormValue("teamName", "")
		if teamName == "" {
			return utils.RespFail(ctx, "TeamName is empty")
		}
		teamTeacher := ctx.FormValue("teamTeacher", "")
		if teamTeacher == "" {
			return utils.RespFail(ctx, "TeamTeacher is empty")
		}
		teamMember := ctx.FormValue("teamMember", "")
		if teamMember == "" {
			return utils.RespFail(ctx, "TeamMember is empty")
		}
		err := json.Unmarshal([]byte(teamTeacher), &user.Team.TeamTeacher)
		if err != nil {
			return utils.RespFail(ctx, "Failed to unmarshal teamTeacher to JSON")
		}
		err = json.Unmarshal([]byte(teamMember), &user.Team.TeamMember)
		if err != nil {
			return utils.RespFail(ctx, "Failed to unmarshal teamMember to JSON")
		}
		user.Team.TeamName = teamName
		user.Team.UUID = uuid.New().String()
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		return utils.RespFail(ctx, "Upload ID Photo Fail")
	}

	files := form.File["teamSchoolCertificate[]"]
	//TODO Remove, it could be upload latter
	// if len(files) < 1 {
	// return utils.RespFail(ctx, "TeamSchoolCertificate is empty")
	// }
	if len(files) != 0 {
		user.Team.IsApplyTeam = true
	}
	allowsuffix := []string{".png", ".jpg", ".jpeg", ".pdf"}

	for _, file := range files {

		attach := models.Attach{UserID: user.ID}
		attach_service.UploadFile(&user.Team, &attach, file, "teamSchoolCertificate", true, allowsuffix)

		err = ctx.SaveFile(file, fmt.Sprintf("%s/%s", attach.SaveLocation, attach.Filename))
		if err != nil {
			return utils.RespFail(ctx, "Upload ID Photo Fail")
		}

		//生成檔案連結
		url := utils.GetURL()
		url = url + attach.Filename
		user.Team.TeamSchoolCertificate = append(user.Team.TeamSchoolCertificate, attach)
	}

	err = user.Update().Error
	if err != nil {
		return utils.RespFail(ctx, "something wrong in server")
	}
	user.FindOne()

	return utils.RespOK(ctx, user, "User Apply Success")
}

func UpdateUser(c *fiber.Ctx) error {
	user := c.Locals("Auth").(*models.User)
	user.FindOne()

	var req models.User
	if err := c.BodyParser(&req); err != nil {
		utils.RespFail(c, err.Error())
	}

	user.Email = req.Email
	user.Phone = req.Phone
	if req.Password != "" {
		user.Password = utils.Md5(req.Password)
	}

	if err := user.Update().Error; err != nil {
		fmt.Println(err)
		return utils.RespFail(c, "Update User Fail")
	}

	return utils.RespOK(c, user, "Update User Success")
}

func GetUserProfile(c *fiber.Ctx) error {
	user := c.Locals("Auth").(*models.User)
	user.FindOne()

	attachSchoolCertificate := attach_service.GetAttachesWithConds(models.Attach{UserID: user.ID, Type: "teamSchoolCertificate"})
	attachWorkPdf := attach_service.GetAttachesWithConds(models.Attach{UserID: user.ID, Type: "workPdf"})
	team := user_service.GetUserTeam(user)

	var schoolCertificateUrl = []string{}
	for _, attach := range *attachSchoolCertificate {
		url := utils.GetURL() + attach.Filename
		schoolCertificateUrl = append(schoolCertificateUrl, url)
	}

	var workPdfUrl string
	for _, attach := range *attachWorkPdf {
		url := utils.GetURL() + attach.Filename
		workPdfUrl = url
	}

	resp := ProfileRes{
		Email: user.Email,
		Phone: user.Phone,

		TeamName:              team.TeamName,
		TeamTeacher:           team.TeamTeacher,
		TeamMember:            team.TeamMember,
		TeamSchoolCertificate: schoolCertificateUrl,
		IsApplyTeam:           team.IsApplyTeam,
		IsUpload:              team.IsUpload,

		WorkVideoLink: team.WorkVideoLink,
		WorkPdf:       workPdfUrl,
	}

	return utils.RespOK(c, resp, "Get User Profile Success")
}
