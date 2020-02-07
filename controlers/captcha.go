package controlers

import (
	"errors"
	"fmt"
	"smh-api/base"
	"smh-api/models"
	"smh-api/service"
	"time"

	"github.com/mojocn/base64Captcha"

	"github.com/gin-gonic/gin"
)

// var configD = base64Captcha.ConfigDigit{
// 	Height:     100,
// 	Width:      300,
// 	MaxSkew:    0.7,
// 	DotCount:   80,
// 	CaptchaLen: 6,
// }

//TermController 期限结构控制器
type CaptchaController struct{}

var captcha = base64Captcha.NewCaptcha(&base64Captcha.DriverDigit{
	Height:   100,
	Width:    300,
	MaxSkew:  0.7,
	DotCount: 80,
	Length:   6,
}, base64Captcha.NewMemoryStore(1024, time.Minute*1))

func (CaptchaController) GetCaption(c *gin.Context) {
	// phone, hasPhone := c.GetQuery("phone")
	// device, hasDevice := c.GetQuery("device")
	// ip := c.Request.RemoteAddr
	// var base64blob, captchaId string
	// var captcaInterfaceInstance base64Captcha.CaptchaInterface
	// if hasPhone {
	// captchaId, captcaInterfaceInstance = base64Captcha.GenerateCaptcha(phone, configD)
	// base64blob = base64Captcha.CaptchaWriteToBase64Encoding(captcaInterfaceInstance)

	id, b64s, err := captcha.Generate()

	base.Response(c, err, map[string]interface{}{"img": b64s, "id": id})
	return
	// }
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
	verifyResult := captcha.Verify(sms.ID, sms.Code, true)
	if verifyResult {
		err = service.SMSService{}.Send(sms.Phone)
	} else {
		err = errors.New("校验验证码错误")
	}

	base.Response(c, err, nil)

}
