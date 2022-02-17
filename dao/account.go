package dao

import (
	"LibraryManagement/model"
	"gorm.io/gorm"
)

var AccountDao *accountDao = &accountDao{}

type accountDao struct{}

func (d accountDao) SelectUser(sess *gorm.DB, email string) (account *model.Account, err error) {
	err = sess.Where("email = ?", email).Find(&account).Error
	return
}

func (d accountDao) UserList(sess *gorm.DB)(account []model.Account, err error) {
	err = sess.Find(&account).Error
	return
}
