package app

import (
	"context"

	"github.com/leometzger/mmonitoring/pkg/models"
)

func (a App) GetUserInfo(accessToken string) (*models.UserInfo, error) {
	ctx := context.Background()
	user, err := a.keycloakClient.GetUserInfo(ctx, accessToken, a.config.KeycloakRealm)

	if err != nil {
		return nil, err
	}

	userInfo := &models.UserInfo{
		Email:    *user.Email,
		Name:     *user.Name,
		LastName: *user.FamilyName,
	}

	return userInfo, err
}
