package api

import "github.com/gin-gonic/gin"

type Book struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Jsonresponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseJSON(c *gin.Context, status int, message string, data any) {
	response := Jsonresponse{
		Status:  status,
		Message: message,
		Data:    data,
	}

	c.JSON(status, response)
}
