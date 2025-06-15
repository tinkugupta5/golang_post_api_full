package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "github.com/example/postapi/models"
)

func CreatePost(c *gin.Context, db *gorm.DB, userId uint) {
    var post models.Post
    if err := c.ShouldBindJSON(&post); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    post.UserID = userId
    db.Create(&post)
    c.JSON(http.StatusCreated, post)
}

func GetPosts(c *gin.Context, db *gorm.DB) {
    var posts []models.Post
    db.Preload("Comments").Find(&posts)
    c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context, db *gorm.DB) {
    var post models.Post
    if err := db.Preload("Comments").First(&post, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "post not found"})
        return
    }
    c.JSON(http.StatusOK, post)
}

func UpdatePost(c *gin.Context, db *gorm.DB, userId uint) {
    var post models.Post
    if err := db.First(&post, c.Param("id")).Error; err != nil || post.UserID != userId {
        c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
        return
    }

    var input models.Post
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Model(&post).Updates(input)
    c.JSON(http.StatusOK, post)
}

func DeletePost(c *gin.Context, db *gorm.DB, userId uint) {
    var post models.Post
    if err := db.First(&post, c.Param("id")).Error; err != nil || post.UserID != userId {
        c.JSON(http.StatusForbidden, gin.H{"error": "not authorized"})
        return
    }

    db.Delete(&post)
    db.Where("post_id = ?", post.ID).Delete(&models.Comment{})
    c.JSON(http.StatusOK, gin.H{"message": "post deleted"})
}
