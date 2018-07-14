package models

import (
	"time"
	."Loango/database"
)

type Banner struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	ImgName string `gorm:"size:255"`
	Url     string `gorm:"size:255"`
	Active  bool
	Weight  int64 `gorm:"size:255"`
}

//CURD

func (b *Banner)Create() error{
	db := InitDB()
	defer db.Close()

	return db.Create(&b).Error

}


func ReadBanners() (*[]Banner,error){
	db := InitDB()
	defer db.Close()

	var banners []Banner
	err := db.Find(&banners).Error
	return &banners,err
}


func DeleteBannerById(id string) error {

	db := InitDB()
	defer db.Close()
	return db.Where("ID = ?", id).Delete(&Banner{}).Error

}

