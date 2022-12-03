package app

import (
	"context"
	"io"

	"github.com/leometzger/mmonitoring/pkg/models"
	"github.com/leometzger/mmonitoring/pkg/storage"
)

func (a App) GetAvatarURL(user *models.User) (*storage.SignedUrl, error) {
	url, err := a.storage.GetAvatarUrl(user.Avatar)

	return url, err
}

func (a App) SaveAvatar(accessToken string, avatarFile io.Reader, ext string) error {
	ctx := context.Background()
	user, err := a.keycloakClient.GetUserInfo(ctx, accessToken, a.config.KeycloakRealm)
	if err != nil {
		return err
	}

	path, err := a.storage.AddAvatar(*user.Sub, avatarFile, ext)
	if err != nil {
		return err
	}

	dbUser, err := a.userStore.FindByKeycloakId(*user.Sub)
	if err != nil {
		if err != models.ErrNotFound {
			return err
		}
		dbUser = &models.User{
			KeycloakID: *user.Sub,
			Avatar:     path,
		}
	}

	dbUser.Avatar = path
	err = a.userStore.Update(dbUser)
	if err != nil {
		return err
	}

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
	keycloakUser, err := a.keycloakClient.GetUserInfo(ctx, accessToken, a.config.KeycloakRealm)
	if err != nil {
		return nil, err
	}
	user, err := a.userStore.FindByKeycloakId(*keycloakUser.Sub)

	if err != nil && err != models.ErrNotFound {
		return nil, err
	}

	avatarUrl, _ := a.GetAvatarURL(user)

	userInfo := &models.UserInfo{
		Avatar:   avatarUrl.Url,
		Email:    *keycloakUser.Email,
		Name:     *keycloakUser.Name,
		LastName: *keycloakUser.FamilyName,
	}

	return userInfo, err
}
