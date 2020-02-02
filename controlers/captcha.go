package controlers

import (
	"errors"
	"fmt"
	"smh-api/base"
	"smh-api/models"
	"smh-api/service"

	"github.com/mojocn/base64Captcha"

	"github.com/gin-gonic/gin"
)

var configD = base64Captcha.ConfigDigit{
	Height:     100,
	Width:      300,
	MaxSkew:    0.7,
	DotCount:   80,
	CaptchaLen: 6,
}

//TermController 期限结构控制器
type CaptchaController struct{}

func (CaptchaController) GetCaption(c *gin.Context) {
	phone, hasPhone := c.GetQuery("phone")
	// device, hasDevice := c.GetQuery("device")
	// ip := c.Request.RemoteAddr
	var base64blob, captchaId string
	var captcaInterfaceInstance base64Captcha.CaptchaInterface
	if hasPhone {
		captchaId, captcaInterfaceInstance = base64Captcha.GenerateCaptcha(phone, configD)
		base64blob = base64Captcha.CaptchaWriteToBase64Encoding(captcaInterfaceInstance)
		base.Response(c, nil, map[string]interface{}{"img": base64blob, "captchaId": captchaId})
		return
	}
	base.Response(c, errors.New("手机号不能为空"), nil)

}

func (CaptchaController) VerificationCaption(c *gin.Context) {
	var err error
	// ip := c.Request.RemoteAddr
	sms := &models.SMS{}
	if err = c.BindJSON(sms); err != nil {
		base.Response(c, errors.New("参数错误"), nil)
		return
	}
	fmt.Println(sms)
	verifyResult := base64Captcha.VerifyCaptchaAndIsClear(sms.Phone, sms.Code, true)
	if verifyResult {
		err = service.SMSService{}.Send(sms.Phone)
	} else {
		err = errors.New("校验验证码错误")
	}

	base.Response(c, err, nil)

}
