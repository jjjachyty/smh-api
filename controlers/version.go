package controlers

import (
	"smh-api/base"
	"smh-api/models"

	"github.com/gin-gonic/gin"
)

//TermController 期限结构控制器
type VersionController struct{}

func (VersionController) Get(c *gin.Context) {
	platform, has := c.GetQuery("platform")
	var err error
	var version *models.Version
	if has {
		version, err = models.GetVsersion(platform)
	}
	base.Response(c, err, version)
}
