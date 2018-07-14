package main

import (
	. "Loango/models"
	. "Loango/routers"
	"fmt"
)

func main() {

	router := InitRouter()

	router.Run(":9000")
	users, _ := GetAllProduct()
	fmt.Print(users)

}
