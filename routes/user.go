package routes

import (
	"d-crud/controller"

	"github.com/gin-gonic/gin"
)

func RouterUser(r *gin.Engine) {
	r.GET("/user", controller.FindUsers)
	r.GET("/user/id/:id", controller.FindUserById)
	r.GET("/user/username/:username", controller.FindUserByUsername)
	r.POST("/user", controller.CreateUser)
	r.PUT("/user/id/update/:id", controller.UpdateUserById)
	r.PUT("/user/username/update/:username", controller.UpdateUserByUsername)
	r.PUT("/user/id/delete/:id", controller.DeleteUserById)
	r.PUT("/user/username/delete/:username", controller.DeleteUserByUsername)
}
