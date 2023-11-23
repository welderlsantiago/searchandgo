package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(c *gin.Context) {
	request := struct {
		role string
	}{}

	c.BindJSON(&request)
}
