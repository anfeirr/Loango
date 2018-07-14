package admin

import (
	"github.com/gin-gonic/gin"
	."Loango/models"

	)

func UserManage(c *gin.Context)  {
	users,err := GetAllUsers()
	if err != nil{
		panic(err)
	}
	c.HTML(200,"adm_userBackendManage.html",gin.H{"users":users})
}