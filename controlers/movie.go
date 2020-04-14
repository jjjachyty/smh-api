package controlers

import (
	"encoding/hex"
	"errors"
	"fmt"
	"smh-api/base"
	"smh-api/middlewares/jwt"
	"smh-api/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

func Newest(c *gin.Context) {
	var err error
	var result []*models.Movie
	fmt.Println("Header", c.Request.Header)
	result, err = models.FindMovie(bson.M{}, 0, 5, bson.M{"createAt": -1})
	base.Response(c, err, result)
}
func Add(c *gin.Context) {
	var err error
	var result = &models.Movie{}
	var has bool
	if err = c.BindJSON(result); err == nil {
		if _, has = c.Get("claims"); has {
			cla := jwt.GetClaims(c)
			result.CreateBy = cla.UserID
		}

		result.ID = xid.New().String()
		result.CreateAt = time.Now()
		result.UpdateAt = result.CreateAt
		if err = result.Insert(); err != nil {
			base.Response(c, err, result)
			return
		}
	}
	base.Response(c, err, result)
}

func AddResources(c *gin.Context) {
	var err error
	var result = &models.Resources{}
	var has bool

	if err = c.BindJSON(result); err == nil {

		if _, has = c.Get("claims"); has {
			cla := jwt.GetClaims(c)
			result.CreateBy = cla.UserID
		}
		result.ID = xid.New().String()
		result.CreateAt = time.Now()
		result.UpdateAt = time.Now()
		if err = result.Insert(); err != nil {
			base.Response(c, err, result)
			return
		}
	}
	base.Response(c, err, result)
}

func All(c *gin.Context) {
	offsetQuery, _ := c.GetQuery("offset")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 64)

	var err error
	var result []*models.Movie
	result, err = models.FindMovie(bson.M{}, offset*12, 12, bson.M{"createAt": -1})
	base.Response(c, err, result)

}

func GetMovie(c *gin.Context) {
	id, has := c.GetQuery("id")
	var err error
	movie := models.Movie{}
	if has {
		err = movie.Get(bson.M{"_id": id})
	}

	base.Response(c, err, movie)
}

func MovieDelete(c *gin.Context) {
	id := c.Param("id")
	var err error
	movie := models.Movie{}
	if id != "" {
		cla := jwt.GetClaims(c)
		err = movie.Remove(bson.M{"_id": id, "createBy": cla.UserID})
	}

	base.Response(c, err, movie)
}

//Recommend 搜索推荐
func Recommend(c *gin.Context) {
	var err error
	var result []*models.Movie
	result, err = models.FindMovie(bson.M{}, 0, 5, bson.M{"createAt": -1})
	base.Response(c, err, result)
}

func Serach(c *gin.Context) {
	var err error
	var result []*models.Movie
	key, has := c.GetQuery("key")
	if has {
		where := bson.M{"$or": []bson.M{bson.M{"name": bson.M{"$regex": key, "$options": "$i"}}, bson.M{"actor": bson.M{"$regex": key, "$options": "$i"}}}}
		result, err = models.FindMovie(where, 0, 10, bson.M{"createAt": -1})
	}
	base.Response(c, err, result)
}

func MyCreateMovies(c *gin.Context) {
	var err error
	var result []*models.Movie
	offsetQuery, hasOffset := c.GetQuery("offset")
	userID, hasUserId := c.GetQuery("userid")

	if hasOffset && hasUserId {
		offset, _ := strconv.ParseInt(offsetQuery, 10, 64)
		where := bson.M{"createBy": userID}
		result, err = models.FindMovie(where, offset, 10, bson.M{"createAt": -1})

	}

	base.Response(c, err, result)

}

func Resources(c *gin.Context) {
	id := c.Param("id")
	var err error
	var result []*models.Resources

	if id != "" {
		result, err = models.FindMovieResources(bson.M{"movieid": id}, 0, 100, bson.M{})
	}
	base.Response(c, err, result)

}

func Apply(c *gin.Context) {
	var err error
	apply := new(models.Apply)
	if err = c.BindJSON(apply); err != nil {
		base.Response(c, errors.New("参数错误"), nil)
		return
	}
	apply.CreateAt = time.Now()
	apply.UpdateAt = apply.CreateAt
	apply.ID = hex.EncodeToString([]byte(strings.TrimSpace(apply.Name)))
	err = apply.Insert()
	base.Response(c, err, nil)
}

func Applys(c *gin.Context) {
	var err error
	var applys []*models.Apply
	offset := c.GetInt64("offset")
	limit := c.GetInt64("limit")
	applys, err = models.FindMovieApply(bson.M{}, offset, limit, bson.M{"createAt": -1})
	base.Response(c, err, applys)
}

func MovieGenre(c *gin.Context) {
	data, err := models.FindMovieGenre()
	base.Response(c, err, data)
}
