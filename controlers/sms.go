package controlers

import (
	"errors"
	"fmt"
	"smh-api/base"
	"smh-api/models"
	"smh-api/service"

	"github.com/gin-gonic/gin"
)

//TermController 期限结构控制器
type SMSController struct{}

func (SMSController) VerificationSMS(c *gin.Context) {
	var err error
	sms := &models.SMS{}
	if err = c.BindJSON(sms); err != nil {
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	fmt.Println(sms)
	err = service.SMSService{}.VerificationSMS(sms.Phone, sms.Code)
	base.Response(c, err, nil)
}
