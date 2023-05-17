package router

import (
	"Easy_IM/docs"
	"Easy_IM/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {

	r := gin.Default()

	// swagger
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 静态资源
	r.Static("/asset", "asset/")
	r.LoadHTMLGlob("views/**/*")

	// 首页
	r.GET("/", service.GetIndex)
	r.GET("/index", service.GetIndex)
	r.GET("/toRegister", service.ToRegister)

	// 用户模块
	r.POST("/user/getUserList", service.GetUserList)
	r.POST("/user/createUser", service.CreateUser)
	r.POST("/user/deleteUser", service.DeleteUser)
	r.POST("/user/updateUser", service.UpdateUser)
	r.POST("/user/findUserByNameAndPwd", service.FindUserByNameAndPwd)

	//发送消息
	r.GET("/user/sendMsg", service.SendMsg)
	r.GET("/user/sendUserMsg", service.SendUserMsg)
	return r
}
