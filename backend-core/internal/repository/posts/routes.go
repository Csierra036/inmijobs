package posts

import "github.com/gin-gonic/gin"

func RegisterPostRoutes(r *gin.RouterGroup, postHandler *PostHandler) {
	r.POST("/")
	r.PUT("/:id", postHandler.EditPost)
	r.DELETE("/:id")
}
