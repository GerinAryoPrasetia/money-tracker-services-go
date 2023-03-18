package routes

import (
	"github.com/GerinAryoPrasetia/go-money-tracking-services/controllers"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userController controllers.UserController
}

func NewRouteUser(userController controllers.UserController) UserRoutes {
	return UserRoutes{userController}
}

func (ur *UserRoutes) UserRoutes(rg *gin.RouterGroup) {
	router := rg.Group("users")
	router.POST("/register", ur.userController.CreateUser)
}
