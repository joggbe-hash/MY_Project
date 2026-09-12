package handlers

import (
	"net/http"

	"me-too-backend/database"
	"me-too-backend/models"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname"`
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "請填寫完整帳號與密碼"})
		return
	}

	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "帳號或密碼錯誤"})
		return
	}

	// Verify password (allows 123456 for default demo or stored password)
	if user.Password != "" && user.Password != req.Password && req.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "帳號或密碼錯誤"})
		return
	}

	// Simple demo token
	token := "token_" + user.Username

	c.JSON(http.StatusOK, gin.H{
		"message": "登入成功",
		"token":   token,
		"user": gin.H{
			"id":          user.ID,
			"username":    user.Username,
			"nickname":    user.Nickname,
			"avatar_url":  user.AvatarURL,
			"bio":         user.Bio,
			"weekly_info": user.WeeklyInfo,
		},
	})
}

func Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "請填寫完整註冊資訊"})
		return
	}

	var count int64
	database.DB.Model(&models.User{}).Where("username = ?", req.Username).Count(&count)
	if count > 0 {
		c.JSON(http.StatusConflict, gin.H{"error": "此帳號已被使用，請更換其他帳號"})
		return
	}

	nickname := req.Nickname
	if nickname == "" {
		nickname = req.Username
	}

	newUser := models.User{
		Username:   req.Username,
		Password:   req.Password,
		Nickname:   nickname,
		AvatarURL:  "",
		Bio:        "把靈感變成行動，把行動留下成果",
		WeeklyInfo: "新加入的創作者",
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "註冊失敗，請稍後再試"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "註冊成功，請登入",
		"user": gin.H{
			"id":       newUser.ID,
			"username": newUser.Username,
			"nickname": newUser.Nickname,
		},
	})
}

func GetMe(c *gin.Context) {
	var user models.User
	if err := database.DB.First(&user, getCurrentUserID()).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "找不到使用者"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": gin.H{
			"id":          user.ID,
			"username":    user.Username,
			"nickname":    user.Nickname,
			"avatar_url":  user.AvatarURL,
			"bio":         user.Bio,
			"weekly_info": user.WeeklyInfo,
		},
	})
}
