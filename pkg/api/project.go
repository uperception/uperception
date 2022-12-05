package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leometzger/mmonitoring/pkg/models"
)

// GET /projects
func (a *Api) QueryProjects(c *gin.Context) {
	projects, err := a.App.QueryProjects()
	if err != nil {
		c.JSON(getStatus(err), gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": projects})
}

// GET /projects/:id
func (a *Api) GetProject(c *gin.Context) {
	project, err := a.App.FindProject(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found!"})
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
		c.JSON(getStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": project})
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
		c.JSON(getStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

// DELETE /projects/:id
func (a *Api) DeleteProject(c *gin.Context) {
	err := a.App.DeleteProject(c.Param("id"))
	if err != nil {
		c.JSON(getStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// POST /projects/:id/lighthouse/endpoints
func (a *Api) CreateLighthouseEndpoint(c *gin.Context) {
	var input models.LighthouseEndpointInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	endpoint, err := a.App.CreateLighthouseEndpoint(c.Param("id"), input)
	if err != nil {
		c.JSON(getStatus(err), gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, endpoint)
}

// POST /projects/:id/lighthouse/endpoints/batch
func (a *Api) CreateLighthouseEndpointsBatch(c *gin.Context) {
}

// PUT /projects/:id/lighthouse/endpoints/:endpointID
func (a *Api) UpdateLighthouseEndpoint(c *gin.Context) {
	var input models.LighthouseEndpointInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	endpoint, err := a.App.UpdateLighthouseEndpoint(c.Param("id"), input)
	if err != nil {
		c.JSON(getStatus(err), gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, endpoint)
}

// DELETE /projects/:id/lighthouse/endpoints/:endpointID
func (a *Api) DeleteLighthouseEndpoint(c *gin.Context) {
	err := a.App.DeleteLighthouseEndpoint(c.Param("endpointID"))
	if err != nil {
		c.JSON(getStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

// GET /projects/:id/lighthouse/endpoints
func (a *Api) ListLighthouseEndpoints(c *gin.Context) {
	endpoints, err := a.App.ListLighthouseEndpoints(c.Param("id"))
	if err != nil {
		c.JSON(getStatus(err), gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": endpoints})
}
