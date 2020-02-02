package controlers

import (
	"encoding/hex"
	"errors"
	"smh-api/base"
	"smh-api/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Newest(c *gin.Context) {
	var err error
	var result []*models.Movie
	result, err = models.FindMovie(bson.M{}, 0, 5, bson.M{"createAt": -1})
	base.Response(c, err, result)
}

func All(c *gin.Context) {
	offsetQuery, _ := c.GetQuery("offset")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 64)

	var err error
	var result []*models.Movie
	result, err = models.FindMovie(bson.M{}, offset*10, 10, bson.M{"createAt": -1})
	base.Response(c, err, result)

}

func GetMovie(c *gin.Context) {
	id, has := c.GetQuery("id")
	var err error
	movie := models.Movie{}
	if has {
		movie.Get(bson.M{"_id": id})
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
