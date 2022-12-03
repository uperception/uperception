package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leometzger/mmonitoring/pkg/models"
)

// GET /profile
func (a *Api) GetProfile(c *gin.Context) {
	user, err := a.App.GetUserInfo(c.GetString("token"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PUT /profile
func (a *Api) UpdateProfile(c *gin.Context) {
	var input models.UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.App.UpdateProfile(c.GetString("token"), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /profile/avatar
func (a *Api) AddAvatar(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	avatarFile, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	ext := file.Filename[len(file.Filename)-4:]

	err = a.App.SaveAvatar(c.GetString("token"), avatarFile, ext)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	c.Status(http.StatusCreated)
}
