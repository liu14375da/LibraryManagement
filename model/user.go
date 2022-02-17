package model

type Account struct {
	BaseEntity
	UserId   string `gorm:"type:varchar(36);comment: '用户编号'";json:"userId"`
	PassWord string `gorm:"type:varchar(64);comment: '密码'";json:"passWord"`
	IsAdmin  string `gorm:"type:varchar(1);comment: '是否管理员'";json:"isAdmin"`
	Salt     string `gorm:"type:varchar(36);comment: '密钥'";json:"salt"`
	Email    string `gorm:"type:varchar(36);comment: '邮箱'";json:"email"`
}

func (receiver *Account) TableName() string {
	return "account"
}
