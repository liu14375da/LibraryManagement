package dao

import (
	"LibraryManagement/model"
	"gorm.io/gorm"
)

var EmailDao *emailDao = &emailDao{}

type emailDao struct{}

func (d emailDao) SelectEmail(tx *gorm.DB, id string) (email model.Email, err error) {
	err = tx.Where("id = ?",id).Find(&email).Error
	return
}


