package resthandler

import "github.com/gin-gonic/gin"

type TodoResthandler interface {
	Create(c *gin.Context)
}