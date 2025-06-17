package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "github.com/example/postapi/models"
    "golang.org/x/crypto/bcrypt"
    "github.com/example/postapi/utils"
)

func Register(c *gin.Context, db *gorm.DB) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)
    user := models.User{Username: input.Username, Password: string(hashedPassword)}
    db.Create(&user)

    c.JSON(http.StatusCreated, gin.H{"message": "registration successful"})
}

func Login(c *gin.Context, db *gorm.DB) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    db.Where("username = ?", input.Username).First(&user)
    if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
        return
    }
    token, _ := utils.GenerateToken(user.ID)
    c.JSON(http.StatusOK, gin.H{"token": token})
}
