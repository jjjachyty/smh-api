package controlers

import (
	"smh-api/base"
	"smh-api/models"

	"github.com/gin-gonic/gin"
)

//PlayerController
type PlayerController struct{}

func (PlayerController) Get(c *gin.Context) {
	id, has := c.GetQuery("id")
	var err error
	var player = new(models.Player)
	if has {
		player.ID = id
		err = player.Get()
	}
	base.Response(c, err, player)
}
