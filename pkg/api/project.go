package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leometzger/mmonitoring/pkg/models"
)

// GET /projects
func (a *Api) QueryProjects(c *gin.Context) {
	projects, _ := a.App.QueryProjects()

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

// GET /projects/:id
func (a *Api) GetProject(c *gin.Context) {
	project, err := a.App.FindProject(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

// POST /projects
func (a *Api) CreateProject(c *gin.Context) {
	var input models.CreateProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := a.App.CreateProject(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

// PUT /projects/:id
func (a *Api) UpdateProject(c *gin.Context) {
	var input models.UpdateProjectInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project, err := a.App.UpdateProject(c.Param("id"), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

// DELETE /projects/:id
func (a *Api) DeleteProject(c *gin.Context) {
	err := a.App.DeleteProject(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted!"})
}
