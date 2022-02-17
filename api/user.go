package api

import (
	"LibraryManagement/dao"
	"LibraryManagement/db"
	"LibraryManagement/middleware"
	"LibraryManagement/model"
	"LibraryManagement/model/request"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var User *user = &user{}

type user struct{}

func (u user) Register(context *gin.Context) {
	var user request.AccountRequest
	context.ShouldBindJSON(&user)

	sess := db.DB
	selectEmail, err := dao.EmailDao.SelectEmail(sess, user.Id)
	if err != nil {
		context.JSON(400, err.Error())
	}

	if selectEmail.Code != user.Code {
		context.JSON(300, errors.New("验证码不正确"))
	}

	passWord, err := middleware.BuildPassSalt(user.Pass, middleware.PASS_WORD_SALT)

	account := model.Account{
		UserId:   strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int()),
		PassWord: passWord,
		IsAdmin:  "0",
		Salt:     middleware.PASS_WORD_SALT,
		Email:    user.Email,
	}
	account.InitEntity(strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Int()))

	if err := sess.Create(&account).Error; err != nil {
		context.JSON(400, err.Error())
	}

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (u user) Login(context *gin.Context) {
	email := context.Query("email")
	pass := context.Query("pass")

	sess := db.DB
	user, err := dao.AccountDao.SelectUser(sess, email)
	if err != nil {
		context.JSON(400, err.Error())
	}

	if user == nil || user.Id == "" {
		context.JSON(205, "用户不存在")
	}

	info := middleware.ValidPassword(user.PassWord, pass, user.Salt)
	if !info {
		context.JSON(205, "用户密码错误")
	}

	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (u user) List(context *gin.Context) {
	sess := db.DB
	user, err := dao.AccountDao.UserList(sess)
	if err != nil {
		context.JSON(400, err.Error())
	}

	if user == nil {
		context.JSON(205, "用户不存在")
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": user,
	})
}

func (u user) Update(context *gin.Context) {
	var user request.AccountRequest
	context.ShouldBindJSON(&user)
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if user.Pass != "" {
			passWord, err := middleware.BuildPassSalt(user.Pass, middleware.PASS_WORD_SALT)
			if err != nil {
				return err
			}

			updateErr := tx.Model(model.Account{}).Where("id = ?", user.Id).Updates(map[string]interface{}{
				"email": user.Email, "pass_word": passWord,
			})
			if updateErr.Error != nil {
				return updateErr.Error
			}
		} else {
			updateErr := tx.Model(model.Account{}).Where("id = ?", user.Id).Updates(map[string]interface{}{
				"email": user.Email,
			})
			if updateErr.Error != nil {
				return updateErr.Error
			}
		}
		return nil
	})
	if err != nil {
		context.JSON(400, err.Error())
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}

func (u user) SearchEmail(context *gin.Context) {
	email := context.Query("email")
	sess := db.DB

	user,err := dao.AccountDao.SelectUser(sess,email)
	if err != nil {
		context.JSON(400,err.Error())
	}
	context.JSON(http.StatusOK,gin.H{
		"code":http.StatusOK,
		"data":user,
	})
}
