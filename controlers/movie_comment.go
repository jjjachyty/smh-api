package controlers

import (
	"errors"
	"smh-api/base"
	"smh-api/middlewares/jwt"
	"smh-api/models"
	"strconv"
	"time"

	"github.com/rs/xid"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

//CommentAdd 新增观看历史
func CommentAdd(c *gin.Context) {
	var err error
	comment := new(models.Comment)
	if err = c.BindJSON(comment); err != nil {
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	comment.CreateAt = time.Now()
	comment.ID = xid.New().String()
	comment.Likes = []string{}
	comment.UnLikes = []string{}
	comment.At = []string{}
	err = comment.Insert()
	base.Response(c, err, comment)
}

//CommentAddLike 点赞
func CommentAddLike(c *gin.Context) {
	var err error
	var params = make(map[string]string, 0)
	if err = c.BindJSON(&params); err != nil {
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	err = models.Comment{ID: params["ID"]}.Update(bson.D{bson.E{"$inc", bson.M{"likecount": 1}}, bson.E{"$addToSet", bson.M{"likes": jwt.GetClaims(c).UserID}}})
	base.Response(c, err, nil)
}

//CommentAddLikeCancel 点赞取消
func CommentAddLikeCancel(c *gin.Context) {
	var err error
	var params = make(map[string]string, 0)
	if err = c.BindJSON(&params); err != nil {
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	err = models.Comment{ID: params["ID"]}.Update(bson.D{bson.E{"$inc", bson.M{"likecount": -1}}, bson.E{"$pull", bson.M{"likes": jwt.GetClaims(c).UserID}}})
	base.Response(c, err, nil)
}

//CommentAddUnLike 踩
func CommentAddUnLike(c *gin.Context) {
	var err error
	var params = make(map[string]string, 0)
	if err = c.BindJSON(&params); err != nil {
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	err = models.Comment{ID: params["ID"]}.Update(bson.D{bson.E{"$inc", bson.M{"unlikecount": 1}}, bson.E{"$addToSet", bson.M{"unlikes": jwt.GetClaims(c).UserID}}})
	base.Response(c, err, nil)
}

//CommentAddUnLike 踩
func CommentAddUnLikeCancel(c *gin.Context) {
	var err error
	var params = make(map[string]string, 0)
	if err = c.BindJSON(&params); err != nil {
		base.Response(c, errors.New("参数错误"), err.Error())
		return
	}
	err = models.Comment{ID: params["ID"]}.Update(bson.D{bson.E{"$inc", bson.M{"unlikecount": -1}}, bson.E{"$pull", bson.M{"unlikes": jwt.GetClaims(c).UserID}}})
	base.Response(c, err, nil)
}

func Comments(c *gin.Context) {
	var err error
	var comments []*models.Comment

	offsetQuery, _ := c.GetQuery("offset")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 64)
	movieid, hasMovieID := c.GetQuery("movieid")

	if hasMovieID {
		comments, err = models.FindComments(bson.M{"movieid": movieid}, offset, 15, bson.M{"createat": -1})

	}
	base.Response(c, err, comments)
}

func UserComments(c *gin.Context) {
	var err error
	var comments []*models.Comment

	offsetQuery, _ := c.GetQuery("offset")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 64)
	userid, hasUserID := c.GetQuery("userid")
	if hasUserID {
		comments, err = models.FindComments(bson.M{"sender": userid}, offset, 15, bson.M{"createat": -1})

	}
	base.Response(c, err, comments)
}
