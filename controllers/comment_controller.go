package controllers

import (
	"net/http"
	"strconv"

	"github.com/example/postapi/models"
	"github.com/example/postapi/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateComment(c *gin.Context, db *gorm.DB, userId uint) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postIdStr := c.Param("postId")
	postIdUint64, err := strconv.ParseUint(postIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid postId"})
		return
	}
	comment.PostID = uint(postIdUint64)
	comment.PostID = uint(utils.ParseUint(c.Param("postId")))
	db.Create(&comment)
	c.JSON(http.StatusCreated, comment)
}
