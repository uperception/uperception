package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/leometzger/mmonitoring/pkg/api/middleware"
	"github.com/leometzger/mmonitoring/pkg/app"
	"github.com/leometzger/mmonitoring/pkg/config"
)

type Api struct {
	App    *app.App
	Router *gin.Engine
}

func NewApi(config *config.Config) *Api {
	router := gin.Default()
	a := Api{
		App:    app.NewApp(config),
		Router: router,
	}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD"},
		AllowHeaders:     []string{"Origin", "content-type", "accept", "authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		// Projects
		v1.GET("/projects", middleware.IsAuthorized("test"), a.QueryProjects)
		v1.GET("/projects/:id", a.GetProject)
		v1.POST("/projects", a.CreateProject)
		v1.PUT("/projects/:id", a.UpdateProject)
		v1.DELETE("/projects/:id", a.DeleteProject)

		// Lighthouse
		v1.PUT("/projects/:id/lighthouse", a.UpdateLighthouseConfig)

		v1.GET("/projects/:id/lighthouse/endpoints", a.ListLighthouseEndpoints)
		v1.POST("/projects/:id/lighthouse/endpoints", a.CreateLighthouseEndpoint)
		v1.POST("/projects/:id/lighthouse/endpoints/batch", a.CreateLighthouseEndpointsBatch)
		v1.PUT("/projects/:id/lighthouse/endpoints/:endpointID", a.UpdateLighthouseEndpoint)
		v1.DELETE("/projects/:id/lighthouse/endpoints/:endpointID", a.DeleteLighthouseEndpoint)

		// Sessions
		// v1.GET("/projects/:id/sessions", a.QuerySessions)
		// v1.GET("/projects/:id/sessions/:id", a.QuerySessions)

		// User
		v1.GET("/profile", middleware.IsAuthorized(""), a.GetProfile)
		v1.PUT("/profile", middleware.IsAuthorized(""), a.UpdateProfile)
		v1.POST("/profile/avatar", middleware.IsAuthorized(""), a.AddAvatar)
		// v1.PUT("/profile", a.UpdateProfile)

		// Organizations
		v1.GET("/organizations", a.QueryOrganizations)
		v1.GET("/organizations/:id", a.FindOrganization)
		v1.POST("/organizations", a.CreateOrganization)
		v1.PUT("/organizations/:id", a.UpdateOrganization)
		// v1.GET("/organizations/:id/projects", a.QueryOrganizationProjects)
		// v1.GET("/organizations/:id/users", a.QueryOrganizationUsers)
		// v1.POST("/organizations/:id/users", a.AddUserToOrganization)
		// v1.DELETE("/organizations/:id/users/:userId", a.DeleteUserFromOrganization)
	}

	return &a
}

func (a *Api) Run() {
	a.Router.Run(":8080")
}

func VerifyToken(c *gin.Context) {
	SecretKey := "secret"

	reqToken := c.GetHeader("Authorization")
	key, er := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))

	if er != nil {
		fmt.Println(er)
		c.Abort()
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Writer.Write([]byte("Unauthorized"))
		return
	}

	token, err := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		fmt.Println(err)
		c.Abort()
		c.Writer.WriteHeader(http.StatusUnauthorized)
		c.Writer.Write([]byte("Unauthorized"))
		return
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("token is valid")
	}
}
