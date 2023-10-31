package utils

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type H struct {
	Code  int
	Msg   string
	Data  interface{}
	Total interface{}
}

func Resp(c *fiber.Ctx, statusCode int, code int, data interface{}, msg string) error {
	c.Set("content-Type", "application/json")
	c.Status(statusCode)

	h := H{
		Code: code,
		Data: data,
		Msg:  msg,
	}
	return c.JSON(h)
}

func RespFail(c *fiber.Ctx, msg string) error {
	return Resp(c, http.StatusBadRequest, -1, nil, msg)
}
func RespInternalServerError(c *fiber.Ctx, msg string) error {
	return Resp(c, http.StatusInternalServerError, -1, nil, msg)
}
func RespUnauthorized(c *fiber.Ctx, msg string) error {
	return Resp(c, http.StatusUnauthorized, -1, nil, msg)
}

func RespOK(c *fiber.Ctx, data interface{}, msg string) error {
	return Resp(c, http.StatusOK, 0, data, msg)
}

func RespOKList(c *fiber.Ctx, data interface{}, total interface{}) error {
	return RespList(c, http.StatusOK, 0, data, total)
}

func RespList(c *fiber.Ctx, statusCode int, code int, data interface{}, total interface{}) error {
	c.Set("content-Type", "application/json")
	c.Status(statusCode)

	h := H{
		Code:  code,
		Data:  data,
		Total: total,
	}
	return c.JSON(h)
}
