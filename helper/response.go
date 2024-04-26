package helper

import (
	"github.com/gin-gonic/gin"
)

func ResponseJSON(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, payload)
}
