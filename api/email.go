package api

import (
	"LibraryManagement/db"
	"LibraryManagement/model"
	"LibraryManagement/model/request"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var EmailApi *emailApi = &emailApi{}

type emailApi struct {
}

func (e emailApi) SendEmail(context *gin.Context) {
	var email request.EmailRequest
	context.ShouldBindJSON(&email)
	sess := db.DB

	mailConn := map[string]string{
		"user": "liuda2498@163.com",
		"pass": "HUTWUYQNCHLPRJAJ",
		"host": "smtp.163.com",
		"port": "25",
	}

	//产生六位数验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))

	//发送的内容
	html := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%s,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, vcode)


	port, _ := strconv.Atoi(mailConn["port"])
	mailInstance := gomail.NewMessage()
	mailInstance.SetHeader("From", mailInstance.FormatAddress(mailConn["user"], "Trustpass官方"))
	mailInstance.SetHeader("To", email.Email)            //发送给多个用户
	mailInstance.SetHeader("Subject", "测试")    //设置邮件主题
	mailInstance.SetBody("text/html", html) //设置邮件正文
	go func() {
		d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
		err := d.DialAndSend(mailInstance)
		if err != nil {
			log.Fatalln("To:", email.Email, "##", "Send Email Failed!Err:", err)
		} else {
			log.Println("To:", email.Email, "##", "Send Email Successfully!")
		}
	}()



	emailSend := model.Email{
		Form:  mailConn["user"],
		To:    email.Email,
		Title: "测试",
		Code:  vcode,
	}
	emailSend.Content.ToJsonData(html)
	emailSend.InitEntity(strconv.Itoa(rnd.Int()))
	if err := sess.Create(&emailSend).Error; err != nil {
		context.JSON(500, err.Error())
	}
	context.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"emailId": emailSend.Id,
	})
}