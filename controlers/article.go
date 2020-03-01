package controlers

import (
	"errors"
	"smh-api/base"
	"smh-api/middlewares/jwt"
	"smh-api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

//ArticleAdd
func ArticleAdd(c *gin.Context) {
	var err error
	article := new(models.Article)
	if err = c.BindJSON(article); err != nil {
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	claims := jwt.GetClaims(c)
	article.ID = xid.New().String()
	article.CreateBy = claims.UserID
	article.CreateAt = time.Now()
	err = article.Insert()

	base.Response(c, err, nil)
}

//ArticleRemove
func ArticleRemove(c *gin.Context) {
	var err error
	article := new(models.Article)
	if err = c.BindJSON(article); err != nil {
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	err = article.Remove(bson.M{"_id": article.ID, "createby": jwt.GetClaims(c).UserID})

	base.Response(c, err, nil)
}

func MyArticles(c *gin.Context) {
	var err error
	var articles []*models.Article

	offsetQuery, hasOffset := c.GetQuery("offset")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 64)
	if hasOffset {
		articles, err = models.FindArticles(bson.M{"createby": jwt.GetClaims(c).UserID}, offset, 15, bson.M{"createAt": -1})
	}
	base.Response(c, err, articles)
}

func Articles(c *gin.Context) {
	var err error
	var articles []*models.Article

	offsetQuery, hasOffset := c.GetQuery("offset")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 64)
	if hasOffset {
		articles, err = models.FindArticles(bson.M{}, offset, 15, bson.M{"createAt": -1})
	}
	base.Response(c, err, articles)
}
