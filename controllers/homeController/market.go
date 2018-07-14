package home

import (
	md "Loango/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Loans(c *gin.Context) {
	loans, _ := md.GetAllProduct()
	c.HTML(http.StatusOK, "hom_market.html", gin.H{
		"title":    "超市",
		"products": loans,
	})

}
