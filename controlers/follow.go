package controlers

import (
	"errors"
	"smh-api/base"
	"smh-api/middlewares/jwt"
	"smh-api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

//FollowAdd 新增观看历史
func FollowAdd(c *gin.Context) {
	var err error
	follow := new(models.Follow)
	if err = c.BindJSON(follow); err != nil {
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	follow.CreateAt = time.Now()

	err = follow.Insert()
	base.Response(c, err, nil)
}

func FollowRemove(c *gin.Context) {
	var err error
	follow := new(models.Follow)
	if err = c.BindJSON(follow); err != nil {
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	follow.UserID = jwt.GetClaims(c).UserID
	follow.CreateAt = time.Now()
	err = follow.Delete()
	base.Response(c, err, nil)
}

func FollowCheck(c *gin.Context) {
	var err error
	follow := new(models.Follow)
	if err = c.BindQuery(follow); err != nil {
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	err = follow.Get(bson.M{"userid": jwt.GetClaims(c).UserID, "followid": follow.FollowID})
	base.Response(c, err, follow)
}

func UserFollows(c *gin.Context) {
	var err error
	var follow []*models.Follow

	offsetQuery, hasOffset := c.GetQuery("offset")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 64)

	if hasOffset {
		follow, err = models.Follows(offset, 15, jwt.GetClaims(c).UserID)

	}
	base.Response(c, err, follow)
}
