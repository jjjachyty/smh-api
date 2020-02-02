package controlers

import (
	"errors"
	"smh-api/base"
	"smh-api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

//AddWatchingHistory 新增观看历史
func AddWatchingHistory(c *gin.Context) {
	var err error
	history := new(models.WatchingHistory)
	if err = c.BindJSON(history); err != nil {
		base.Response(c, errors.New("参数错误"), nil)
		return
	}
	history.CreateAt = time.Now()
	history.UpdateAt = history.CreateAt
	err = history.Insert()
	base.Response(c, err, nil)
}

//UpdateWatchingHistory 更新观看历史
func UpdateWatchingHistory(c *gin.Context) {
	var err error
	history := new(models.WatchingHistory)
	if err = c.BindJSON(history); err != nil {
		base.Response(c, errors.New("参数错误"), nil)
		return
	}

	err = history.Update()
	if err == nil {
		watching := models.WatchingHistory{MovieID: history.MovieID, MovieThumbnail: history.MovieThumbnail, UserID: history.UserID}
		watching.Insert()
	}
	base.Response(c, err, nil)

}

func WatchingHistory(c *gin.Context) {
	var err error
	var history = &models.WatchingHistory{}
	if err = c.BindJSON(history); err != nil {
		base.Response(c, errors.New("参数错误"), nil)
		return
	}
	err = history.Get(bson.M{"userid": history.UserID, "movieid": history.MovieID, "resourcesid": history.ResourcesID})
	base.Response(c, err, history)
}

func WatchingHistorys(c *gin.Context) {
	var err error
	var historys []*models.WatchingHistory

	offsetQuery, _ := c.GetQuery("offset")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 64)
	userid, _ := c.GetQuery("userid")

	historys, err = models.FindWatchHistorys(bson.M{"userid": userid}, offset, 15, bson.M{"createAt": -1})
	base.Response(c, err, historys)
}

//Watchings 查询正在观看的电影
func Watchings(c *gin.Context) {
	var err error
	var historys []*models.Watching

	offsetQuery, _ := c.GetQuery("offset")
	offset, _ := strconv.ParseInt(offsetQuery, 10, 64)

	historys, err = models.FindWatching(offset, 15)
	base.Response(c, err, historys)
}
