package service

import (
	"crypto/sha256"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"smh-api/base"
	"smh-api/models"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

const (
	REGISTER_TPID = "a061eb95f24441f99ea611d0bb3ca4ca" //注册验证码模板ID
)

//必填,请参考"开发准备"获取如下数据,替换为实际值
var realURL = "https://api.rtc.huaweicloud.com:10443/sms/batchSendSms/v1" //APP接入地址+接口访问URI
var appKey = "b9JGuVoNnD529tLfu7Ui3x48LsHq"                               //APP_Key
var appSecret = "HBkCrHvT71sguLt2nLtF9GQZ73F0"                            //APP_Secret
var sender = "8819081219913"                                              //国内短信签名通道号或国际/港澳台短信通道号
var signature = "register"                                                //签名名称

type SMSService struct{}

func (SMSService) Send(phone string) error {
	var err error
	sms := &models.SMS{Phone: phone}
	if err = sms.Get(); err == nil {
		if sms.Code == "" {
			sms.Code = base.CreateRandomNumber(4)

			if err = sms.Insert(); err == nil {
				return send(phone, REGISTER_TPID, []string{sms.Code})
			}
		}

		if time.Now().Sub(sms.CreateAt) > 60*time.Second {
			sms.Code = base.CreateRandomNumber(4)

			if err = sms.Update(); err == nil {
				return send(phone, REGISTER_TPID, []string{sms.Code})
			}
		}
	}
	err = errors.New("一分钟只能发送一次")

	return err

}

func (SMSService) VerificationSMS(phone, code string) error {
	var err error
	sms := &models.SMS{Phone: phone}
	if err = sms.Get(); err == nil {
		if sms.Code == code {
			err = sms.Delete()
			return err
		}
	}
	return errors.New("短信校验失败")
}

//发送短息
func send(to string, templateID string, templateParas []string) error {
	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	},
	}
	v := url.Values{}
	templateParasByts, _ := json.Marshal(templateParas)
	v.Set("from", sender)
	v.Add("to", to)
	v.Add("templateId", templateID)
	v.Add("templateParas", string(templateParasByts))
	v.Add("statusCallback", "")
	v.Add("signature", signature)

	var err error
	var resp *http.Response
	var req *http.Request
	var body []byte
	if req, err = http.NewRequest("POST", realURL, strings.NewReader(v.Encode())); err == nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Authorization", "WSSE realm=\"SDP\",profile=\"UsernameToken\",type=\"Appkey\"")
		xwsse := buildWsseHeader(appKey, appSecret)
		fmt.Println("xwsse", xwsse)
		req.Header.Set("X-WSSE", xwsse)

		if resp, err = client.Do(req); err == nil {
			defer resp.Body.Close()
			if body, err = ioutil.ReadAll(resp.Body); err == nil {
				fmt.Println(string(body))
				code := gjson.Get(string(body), "code").String()
				if code == "000000" {
					return nil
				}
				return fmt.Errorf("发送短信消息失败 code=%s", code)
			}

		}
	}
	return err

}

func buildWsseHeader(appKey, appSecret string) string {

	var created = time.Now().Local().UTC().Format("2006-01-02T15:04:05Z")
	var nonce = strings.ReplaceAll(base.GetUUID(), "-", "")
	hash := sha256.New()
	hash.Write([]byte(nonce + created + appSecret))

	return fmt.Sprintf("UsernameToken Username=%q,PasswordDigest=%q,Nonce=%q,Created=%q", appKey, base64.StdEncoding.EncodeToString(hash.Sum(nil)), nonce, created)
}
