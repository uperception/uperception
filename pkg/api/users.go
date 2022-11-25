package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	user, err := a.App.GetUserInfo(c.GetString("token"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
