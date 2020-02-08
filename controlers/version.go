package controlers

import (
	"smh-api/base"
	"smh-api/models"

	"github.com/gin-gonic/gin"
)

//TermController 期限结构控制器
type VersionController struct{}

func (VersionController) Get(c *gin.Context) {
	version, err := models.GetVsersion()
	base.Response(c, err, version)
}
