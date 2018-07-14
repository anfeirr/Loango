package api

import (
	. "Loango/database"
	. "Loango/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func UserAdd(c *gin.Context) {
	var user User
	c.Bind(&user)

	if user.Name != "" {
		db := InitDB()
		defer db.Close()

		if user.Password == "" {
			user.Password = "666666"
		}

		db.Create(&user)
		c.JSON(200, gin.H{
			"success ": user,
		})

	} else {
		c.JSON(422, gin.H{"error ": "User name can't be empty "})
	}
}

func UserGetAll(c *gin.Context) {
	users, _ := GetAllUsers()
	c.JSON(200, gin.H{"users :": users})
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")
	err := DeleteUserById(id)
	if err != nil {
		fmt.Print(err)
		c.JSON(200, gin.H{"error: ": "admin can't be deleted "})
	}
	c.JSON(200, gin.H{"success": id})
}


