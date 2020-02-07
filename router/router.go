package router

import (
	"smh-api/controlers"
	"smh-api/middlewares/jwt"

	"github.com/gin-gonic/gin"
)

func Init(e *gin.Engine) {
	v1 := e.Group("/api/v1")
	{
		base := v1.Group("/base")
		{
			base.GET("/captcha", controlers.CaptchaController{}.GetCaption)
			base.POST("/captcha", controlers.CaptchaController{}.VerificationCaption)
			base.POST("/sms", controlers.SMSController{}.VerificationSMS)

		}

		movie := v1.Group("/movie")
		{
			movie.GET("/newest", controlers.Newest)
			movie.GET("/get", controlers.GetMovie)
			movie.GET("/resources/:id", controlers.Resources)
			movie.GET("/serach", controlers.Serach)
			movie.GET("/applys", controlers.Applys)
			movie.GET("/all", controlers.All)
			movie.GET("/recommend", controlers.Recommend)
			movie.GET("/watchs", controlers.WatchingHistorys)
			movie.GET("/watching", controlers.Watchings)
			movie.POST("/add", controlers.Add)
			movie.POST("/addresources", controlers.AddResources)

			movie.Use(jwt.JWTAuth())
			movie.POST("/apply", controlers.Apply)
			movie.POST("/addwatch", controlers.AddWatchingHistory)
			movie.POST("/updatewatch", controlers.UpdateWatchingHistory)
			movie.POST("/watch", controlers.WatchingHistory)

			// v1.Use(jwt.JWTAuth())

		}
		user := v1.Group("/user")
		{
			user.GET("/checkphone", controlers.UserCheckPhone)
			user.POST("/register", controlers.UserRegister)
			user.POST("/login", controlers.UserLoginWithPW)
			user.Use(jwt.JWTAuth())
			user.POST("/updateinfo", controlers.UserUpdateInfo)

		}
		comment := v1.Group("/comment")
		{
			comment.GET("/list", controlers.Comments)
			comment.Use(jwt.JWTAuth())
			comment.POST("/add", controlers.CommentAdd)
			comment.POST("/like", controlers.CommentAddLike)
			comment.POST("/likecancel", controlers.CommentAddLikeCancel)
			comment.POST("/unlike", controlers.CommentAddUnLike)
			comment.POST("/unlikecancel", controlers.CommentAddUnLikeCancel)

		}

	}
}
