package models

import "gorm.io/gorm"

type Books struct {
	ID        uint    `gorm:"primary key;autoIncreament" json:"id"`
	Author    *string `json:"author"`
	Title     *string `json:"title"`
	Publisher *string `json:"publisher"`
	NoOfPages *int    `json:"noOfPages"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
