package routers

import (
	"ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter(){
	gin.SetMode(utils.AppMode)
	r:=gin.Default()

	router :=r.Group("api/v1")
	{
		// User模块的路由接口
		router.POST("user/add",v1.AddUser)
		router.PUT("users/:id",v1.EditUser)
		router.GET("users",v1.GetUsers)
		router.DELETE("user/:id",v1.DeleteUser)
		// 分类模块的路由接口
		router.POST("category/add",v1.AddCategory)
		router.PUT("category:id",v1.EditCate)
		router.GET("category",v1.GetCate)
		router.DELETE("category/:id",v1.DeleteCate)

		// 文章模块的路口接口

	}





	r.Run(utils.HttpPort)


}