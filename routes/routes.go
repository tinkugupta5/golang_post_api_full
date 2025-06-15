package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/example/postapi/controllers"
    "github.com/example/postapi/middleware"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    r.POST("/register", func(c *gin.Context) {
        controllers.Register(c, db)
    })
    r.POST("/login", func(c *gin.Context) {
        controllers.Login(c, db)
    })

    auth := r.Group("/")
    auth.Use(middleware.AuthMiddleware())

    auth.POST("/posts", func(c *gin.Context) {
        controllers.CreatePost(c, db, c.GetUint("userId"))
    })
    auth.PUT("/posts/:id", func(c *gin.Context) {
        controllers.UpdatePost(c, db, c.GetUint("userId"))
    })
    auth.DELETE("/posts/:id", func(c *gin.Context) {
        controllers.DeletePost(c, db, c.GetUint("userId"))
    })
    auth.POST("/posts/:postId/comments", func(c *gin.Context) {
        controllers.CreateComment(c, db, c.GetUint("userId"))
    })

    r.GET("/posts", func(c *gin.Context) {
        controllers.GetPosts(c, db)
    })
    r.GET("/posts/:id", func(c *gin.Context) {
        controllers.GetPost(c, db)
    })
}
