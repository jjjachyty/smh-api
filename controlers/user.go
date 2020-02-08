package controlers

import (
	"errors"
	"fmt"
	"smh-api/base"
	"smh-api/middlewares/jwt"
	"smh-api/models"
	"smh-api/service"
	"time"

	"github.com/rs/xid"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UserRegister(c *gin.Context) {
	var err error
	user := new(models.User)

	if err = c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err.Error())
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	user.ID = xid.New().String()
	user.CreateAt = time.Now()
	user.State = true
	user.IP = c.ClientIP()
	user.NickName = user.Phone[8:]
	user.PassWord = base.GetMD5(user.PassWord)
	err = user.Insert()
	base.Response(c, err, nil)

}

func UserCheckPhone(c *gin.Context) {
	phone, has := c.GetQuery("phone")
	var err error
	if has {
		user := models.User{Phone: phone}
		if err = user.Get(bson.M{"phone": user.Phone}); err == nil {
			if user.ID != "" {
				err = errors.New("用户已存在,请直接登录")
			}
		}
	}
	base.Response(c, err, nil)
}

//UserLoginWithPW 用户名密码登录
func UserLoginWithPW(c *gin.Context) {
	var err error
	var token string
	var user = new(models.User)
	if err = c.BindJSON(user); err == nil {
		err = user.Get(bson.M{"phone": user.Phone, "password": base.GetMD5(user.PassWord)})
		fmt.Println(err, "ID", user.ID)
		if user.ID != "" {
			token, err = jwt.GenerateToken(*user)
		} else {
			err = errors.New("用户名或密码错误")
		}

	}
	fmt.Println(err, user)
	base.Response(c, err, map[string]interface{}{"User": user, "Token": token})
}

//LoginWithSMS 短信验证码登录
func UserLoginWithSMS(c *gin.Context) {
	var err error
	var user = new(models.User)

	if err = c.BindJSON(user); err != nil {
		base.Response(c, err, nil)
		return
	}

	if err = (service.SMSService{}).VerificationSMS(user.Phone, user.PassWord); err != nil {
		base.Response(c, err, nil)
		return
	}
	var token string

	if err = user.Get(bson.M{"phone": user.Phone}); err != nil {
		base.Response(c, err, nil)
		return
	}

	if token, err = jwt.GenerateToken(*user); err != nil {
		base.Response(c, err, nil)
	}

	fmt.Println(user)
	base.Response(c, err, map[string]interface{}{"User": user, "Token": token})
}

func UserUpdateInfo(c *gin.Context) {
	var err error
	var user = new(models.User)
	if err = c.BindJSON(user); err == nil {
		err = user.Update(bson.M{"$set": bson.M{"nickname": user.NickName, "introduce": user.Introduce}})
	}
	base.Response(c, err, nil)
}
