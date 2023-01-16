package controller

import (
	"github.com/gin-gonic/gin"
	"tiktok/middlewave"
)

func InitRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	apiRouter.POST("/user/register/", Register)
	apiRouter.POST("/user/login/", Login)
	apiRouter.GET("/feed/", Feed)

	authed := apiRouter.Use(middlewave.Auth)
	authed.GET("/user/", UserInfo)
	authed.POST("/publish/action/", Publish)
	authed.GET("/publish/list/", PublishList)

	authed.POST("/favorite/action/", FavoriteAction)
	authed.GET("/favorite/list/", FavoriteList)
	authed.POST("/comment/action/", CommentAction)
	authed.GET("/comment/list/", CommentList)

	authed.POST("/relation/action/", RelationAction)
	authed.GET("/relation/follow/list/", FollowList)
	authed.GET("/relation/follower/list/", FollowerList)
	authed.GET("/relation/friend/list/", FriendList)
	authed.GET("/message/chat/", MessageChat)
	authed.POST("/message/action/", MessageAction)
}
