package routers

import (
	hom "Loango/controllers/homeController"
	adm "Loango/controllers/admin"
	"github.com/gin-gonic/gin"
	apiv1 "Loango/controllers/api"

)







func InitRouter() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/statics", "./statics")

	api := router.Group("/api/v1")
	{
		api.POST("users",apiv1.UserAdd)
		api.GET("users",apiv1.UserGetAll)
		api.POST("login",apiv1.Login)
		//api.GET("users/:id",apiv1)
	}

	home := router.Group("/")
	{
		home.GET("/", hom.Home)
		home.GET("/market", hom.Loans)
	}

	admin := router.Group("/admin")
	{
		admin.GET("loanAdd",adm.LoanAdd)
		admin.POST("loanAdd",adm.LoanAddPost)
		admin.GET("loanManage",adm.LoanManage)
		admin.GET("loanDelete/:id",adm.LoanDelete)
		admin.GET("loanEdit/:id",adm.LoanEdit)

		//Banner管理
		admin.GET("bannerManage",adm.BannerManage)
		admin.GET("bannerAdd",adm.BannerAdd)
		admin.POST("bannerAdd",adm.BannerAddPost)
		admin.GET("bannerDelete/:id",adm.BannerDelete)

		//后台用户管理
		admin.GET("userManage",adm.UserManage)
	}

	return router
}
