package home

import (
	."Loango/models"
	"github.com/gin-gonic/gin"
	"net/http"

)

func Home(c *gin.Context) {
	loans, _ := GetAllProduct()

	banners,_ := ReadBanners()


	c.HTML(http.StatusOK, "hom_home.html", gin.H{
		"title":"首页",
		"products": loans,
		"banners":banners,
	})
}
