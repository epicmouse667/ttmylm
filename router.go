package main

import (
	"dou_sheng/controller"
	"dou_sheng/util"
	"github.com/gin-gonic/gin"
	"sync"
)

func initRouter(r *gin.Engine) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		r.Static("/static", "./public")

		apiRouter := r.Group("/douyin")

		// basic apis
		apiRouter.GET("/feed/", controller.Feed)
		//apiRouter.GET("/user/", controller.UserInfo)
		//apiRouter.POST("/user/register/", controller.Register)
		//apiRouter.POST("/user/login/", controller.Login)
		//apiRouter.POST("/publish/action/", controller.Publish)
		apiRouter.GET("/publish/list/", controller.PublishList)
		//
		//// extra apis - I
		apiRouter.POST("/favorite/action/", controller.FavoriteAction)
		apiRouter.GET("/favorite/list/", controller.FavoriteList)
		//apiRouter.POST("/comment/action/", controller.CommentAction)
		//apiRouter.GET("/comment/list/", controller.CommentList)
		//
		//// extra apis - II
		//apiRouter.POST("/relation/action/", controller.RelationAction)
		//apiRouter.GET("/relation/follow/list/", controller.FollowList)
		//apiRouter.GET("/relation/follower/list/", controller.FollowerList)
	}()
	go func() {
		defer wg.Done()
		util.Init()
	}()
	wg.Wait()
	// public directory is used to serve static resources

}
