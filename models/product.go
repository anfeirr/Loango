package models

import (
	. "Loango/database"

	"github.com/pkg/errors"
	"strconv"
	"time"
)

type Product struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Name       string `gorm:"size:255" form:"name"`
	Url        string `gorm:"size:255" form:"url"`
	Logo       string `gorm:"size:255" `
	AmountMax  int64  `form:"amount_max"`
	Period     string `gorm:"size:255" form:"period"`
	Advantage  string `gorm:"size:255" form:"advantage"`
	Applicants int64  `form:"applicants"`
	Weight     int64  `form:"weight"`

	Company         string `gorm:"size:255" form:"compnay"`
	CooperateMethod string `gorm:"size:255"	form:"cooperate"`
}

//CURD

func (p *Product) Create() error {
	if p.Name != "" {
		db := InitDB()
		defer db.Close()

		return db.Create(&p).Error

	}
	return errors.New("Product name is nil ")

}

func (p *Product) Update() error {
	db := InitDB()
	defer db.Close()

	return db.Save(&p).Error
}

func GetProductByID(id string) (*Product, error) {

	idInt, _ := strconv.ParseUint(id, 10, 64)
	var product Product

	db := InitDB()
	defer db.Close()

	err := db.Where("id = ?", idInt).First(&product).Error
	return &product, err

}

func GetAllProduct() (*[]Product, error) {
	db := InitDB()
	defer db.Close()

	var users []Product

	err := db.Find(&users).Error
	return &users, err
}

func DeleteProductByID(id string) error {
	db := InitDB()
	defer db.Close()

	return db.Where("ID =?", id).Unscoped().Delete(&Product{}).Error
}
