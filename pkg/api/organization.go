package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leometzger/mmonitoring/pkg/models"
)

// GET /api/organizations
func (a *Api) QueryOrganizations(c *gin.Context) {
	organizations := a.App.QueryOrganizations()

	c.JSON(http.StatusOK, gin.H{"data": organizations})
}

// GET /api/organizations/:id
func (a *Api) FindOrganization(c *gin.Context) {
	organization, err := a.App.FindOrganization(c.Param("id"))
	if err != nil {
		c.JSON(getStatus(err), gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": organization})
}

// POST /api/organizations
func (a *Api) CreateOrganization(c *gin.Context) {
	var input models.CreateOrganizationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization, err := a.App.CreateOrganization(input)
	if err != nil {
		c.JSON(getStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": organization})
}

// PUT /api/organizations/:id
func (a *Api) UpdateOrganization(c *gin.Context) {
	var input models.UpdateOrganizationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization, err := a.App.UpdateOrganization(c.Param("id"), input)
	if err != nil {
		c.JSON(getStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": organization})
}
