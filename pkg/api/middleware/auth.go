package middleware

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	oidc "github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	ResourceAccess client `json:"resource_access,omitempty"`
	JTI            string `json:"jti,omitempty"`
}

type client struct {
	ServiceClient clientRoles `json:"mmonitoring,omitempty"`
}

type clientRoles struct {
	Roles []string `json:"roles,omitempty"`
}

var RealmConfigURL string = "http://localhost:3000/realms/UPerception"
var clientID string = "account"

func IsAuthorized(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		accessTokenHeader := c.Request.Header["Authorization"]

		if len(accessTokenHeader) != 1 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		rawAccessToken := accessTokenHeader[0]
		c.Set("token", rawAccessToken)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{
			Timeout:   time.Duration(6000) * time.Second,
			Transport: tr,
		}
		ctx := oidc.ClientContext(context.Background(), client)
		provider, err := oidc.NewProvider(ctx, RealmConfigURL)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		oidcConfig := &oidc.Config{
			ClientID: clientID,
		}
		verifier := provider.Verifier(oidcConfig)
		idToken, err := verifier.Verify(ctx, rawAccessToken)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if role == "" {
			c.Next()
			return
		}

		var IDTokenClaims Claims
		if err := idToken.Claims(&IDTokenClaims); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		user_access_roles := IDTokenClaims.ResourceAccess.ServiceClient.Roles
		for _, b := range user_access_roles {
			if b == role {
				c.Next()
				return
			}
		}

		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
