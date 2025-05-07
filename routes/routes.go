package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rajesh/personal-skill-tracker/controllers"
	"github.com/rajesh/personal-skill-tracker/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())

	auth.GET("/users", controllers.GetUsers)
	auth.POST("/skills", controllers.CreateSkill)
	auth.GET("/skills", controllers.GetSkills)
	auth.POST("/logs", controllers.CreateLog)
	auth.GET("/logs", controllers.GetLogs)
}
