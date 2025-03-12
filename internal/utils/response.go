package utils

import "github.com/labstack/echo/v4"

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c echo.Context, message string, data interface{}) error {
	return c.JSON(200, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c echo.Context, message string) error {
	return c.JSON(400, APIResponse{
		Status:  "error",
		Message: message,
		Data:    nil,
	})
}
