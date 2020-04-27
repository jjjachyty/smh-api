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
			base.GET("/version", controlers.VersionController{}.Get)
			base.POST("/captcha", controlers.CaptchaController{}.VerificationCaption)
			base.POST("/sms", controlers.SMSController{}.VerificationSMS)
			base.POST("/refreshtoken", jwt.RefreshToken)

		}

		v1.GET("/player", controlers.PlayerController{}.Get)

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
			movie.GET("/genre", controlers.MovieGenre)

			movie.Use(jwt.JWTAuth())
			movie.POST("/apply", controlers.Apply)
			movie.POST("/addwatch", controlers.AddWatchingHistory)
			movie.POST("/updatewatch", controlers.UpdateWatchingHistory)
			movie.POST("/watch", controlers.WatchingHistory)
			movie.DELETE("/id/:id", controlers.MovieDelete)

			// v1.Use(jwt.JWTAuth())

		}
		user := v1.Group("/user")
		{
			user.GET("/checkphone", controlers.UserCheckPhone)
			user.GET("/info/:id", controlers.UserInfoByID)

			user.POST("/register", controlers.UserRegister)
			user.POST("/login", controlers.UserLoginWithPW)
			user.POST("/loginsms", controlers.UserLoginWithSMS)
			user.Use(jwt.JWTAuth())
			// user.POST("/vip", controlers.UserVIP)
			user.POST("/updateinfo", controlers.UserUpdateInfo)
			user.GET("/moviecreate", controlers.MyCreateMovies)
			user.GET("/moviecomments", controlers.UserComments)
			user.GET("/follows", controlers.UserFollows)
			user.GET("/followcheck", controlers.FollowCheck)

			user.POST("/followadd", controlers.FollowAdd)
			user.POST("/followremove", controlers.FollowRemove)
			user.GET("/checkvip", controlers.UserCheckVIP)

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

		article := v1.Group("/article")
		{
			article.GET("/list", controlers.Articles)
			article.Use(jwt.JWTAuth())
			article.POST("/add", controlers.ArticleAdd)
			article.GET("/my", controlers.MyArticles)
			article.POST("/remove", controlers.ArticleRemove)

		}

	}
}
