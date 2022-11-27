package app

import (
	"context"
	"io"

	"github.com/leometzger/mmonitoring/pkg/models"
)

func (a App) GetAvatarURL(accessToken string) (string, error) {
	ctx := context.Background()
	user, err := a.keycloakClient.GetUserInfo(ctx, accessToken, a.config.KeycloakRealm)
	if err != nil {
		return "", err
	}
	url, err := a.storage.GetAvatarUrl(*user.Sub)

	return url, err
}

func (a App) SaveAvatar(accessToken string, avatarFile io.Reader) error {
	ctx := context.Background()
	user, err := a.keycloakClient.GetUserInfo(ctx, accessToken, a.config.KeycloakRealm)
	if err != nil {
		return err
	}

	err = a.storage.AddAvatar(*user.Sub, avatarFile)
	return err
}

func (a App) UpdateProfile(accessToken string, input models.UpdateProfileInput) (*models.UserInfo, error) {
	ctx := context.Background()
	a.refreshKeycloakToken()

	keycloakUserInfo, err := a.keycloakClient.GetUserInfo(ctx, accessToken, a.config.KeycloakRealm)
	if err != nil {
		return nil, err
	}

	user, err := a.keycloakClient.GetUserByID(ctx, a.keycloakAdminToken.AccessToken, a.config.KeycloakRealm, *keycloakUserInfo.Sub)
	if err == nil {
		user.FirstName = &input.Name
		user.LastName = &input.LastName
	} else {
		return nil, err
	}

	err = a.keycloakClient.UpdateUser(ctx, a.keycloakAdminToken.AccessToken, a.config.KeycloakRealm, *user)
	if err != nil {
		return nil, err
	}

	userInfo := &models.UserInfo{
		Email:    *user.Email,
		Name:     input.Name,
		LastName: input.LastName,
	}

	return userInfo, err
}

func (a App) GetUserInfo(accessToken string) (*models.UserInfo, error) {
	ctx := context.Background()
	user, err := a.keycloakClient.GetUserInfo(ctx, accessToken, a.config.KeycloakRealm)

	if err != nil {
		return nil, err
	}

	avatarUrl, _ := a.GetAvatarURL(accessToken)

	userInfo := &models.UserInfo{
		Avatar:   avatarUrl,
		Email:    *user.Email,
		Name:     *user.Name,
		LastName: *user.FamilyName,
	}

	return userInfo, err
}
