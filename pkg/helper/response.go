package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Code      int         `json:"code"`
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Page      int64       `json:"page,omitempty"`
	PageSize  int64       `json:"page_size,omitempty"`
	Total     int64       `json:"total,omitempty"`
	TotalPage int64       `json:"total_page,omitempty"`
}

func GenerateResponse(ctx *fiber.Ctx, code int, message string, data interface{}, success bool) error {
	return ctx.Status(code).JSON(Response{
		Code:    code,
		Success: success,
		Message: message,
		Data:    data,
	})
}

func GenerateResponsePagination(ctx *fiber.Ctx, res Response) error {
	return ctx.Status(res.Code).JSON(res)
}
