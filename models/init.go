package models

import (
	. "Loango/database"
)

func init() {
	db := InitDB()
	defer db.Close()

	db.AutoMigrate(&Product{}, &User{},&Banner{})
}
