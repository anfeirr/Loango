package api

import (
	"github.com/gin-gonic/gin"
	."Loango/models"
)

func Authentication(userName string, password string) bool {

	user,err := GetUserByName(userName)
	if err != nil{
		panic(err)
	}

	if password == user.Password {
		return true
	}
	return false
}

func Login(c *gin.Context) {
	user := c.PostForm("user")
	password := c.PostForm("password")

	au := Authentication(user, password)

	if au {
		c.JSON(200, gin.H{"Success:": "login successfully "})
	} else {
		c.JSON(200, gin.H{"error:": "Name or password is wrong "})
	}

}
