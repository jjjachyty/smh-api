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
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	user.ID = xid.New().String()
	user.CreateAt = time.Now()
	user.State = true
	user.IP = c.ClientIP()
	user.NickName = user.Phone[8:]
	user.PassWord = base.GetMD5(user.PassWord)
	user.VIPEndTime = time.Now().Add(time.Hour * 24)
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
		if user.ID != "" {
			token, err = jwt.GenerateToken(*user)
		} else {
			err = errors.New("用户名或密码错误")
		}

	}
	base.Response(c, err, map[string]interface{}{"User": user, "Token": token})
}

//UserVIP 续VIP
func UserVIP(c *gin.Context) {
	var err error
	var user = new(models.User)
	cla := jwt.GetClaims(c)
	user.ID = cla.UserID
	var expTime = time.Now().Add(time.Hour * 24)
	err = user.Update(bson.M{"$set": bson.M{"vipendtime": expTime}})
	base.Response(c, err, expTime)
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
		err = user.Update(bson.M{"$set": bson.M{"nickname": user.NickName, "introduce": user.Introduce, "sex": user.Sex, "avatar": user.Avatar}})
	}
	base.Response(c, err, nil)
}

func UserInfoByID(c *gin.Context) {
	var err error
	var userID = c.Param("id")
	var user = &models.User{}
	if userID != "" {
		err = user.Get(bson.M{"_id": userID})
	}
	base.Response(c, err, models.User{ID: userID, NickName: user.NickName, Avatar: user.Avatar})
}
