package main

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
	"tiktok/middlewave"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.GET("/feed/", controller.Feed)

	authed := apiRouter.Use(middlewave.Auth)
	authed.GET("/user/", controller.UserInfo)
	authed.POST("/publish/action/", controller.Publish)
	authed.GET("/publish/list/", controller.PublishList)

	authed.POST("/favorite/action/", controller.FavoriteAction)
	authed.GET("/favorite/list/", controller.FavoriteList)
	authed.POST("/comment/action/", controller.CommentAction)
	authed.GET("/comment/list/", controller.CommentList)

	authed.POST("/relation/action/", controller.RelationAction)
	authed.GET("/relation/follow/list/", controller.FollowList)
	authed.GET("/relation/follower/list/", controller.FollowerList)
	authed.GET("/relation/friend/list/", controller.FriendList)
	authed.GET("/message/chat/", controller.MessageChat)
	authed.POST("/message/action/", controller.MessageAction)
}
