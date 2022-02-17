package model

import model "LibraryManagement/model/type"

type Email struct {
	BaseEntity
	Form    string     `gorm:"type:varchar(36);comment: '发送邮箱'";json:"form"`
	To      string     `gorm:"type:varchar(36);comment: '接收邮箱'";json:"to"`
	Code    string     `gorm:"type:varchar(35);comment: '验证码'";json:"code"`
	Title   string     `gorm:"type:varchar(255);comment: '发送标题'";json:"title"`
	Content model.JSON `gorm:"type:json;comment: '发送内容'";json:"content"`
}

func (e *Email) TableName() string {
	return "send_email"
}