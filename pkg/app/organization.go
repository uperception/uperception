package app

import "github.com/leometzger/mmonitoring/pkg/models"

func (a App) CreateOrganization(org models.CreateOrganizationInput) (*models.Organization, error) {
	organization := models.Organization{
		Name:        org.Name,
		Description: org.Description,
	}

	err := a.store.OrganizationStore().Save(&organization)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}

func (a App) UpdateOrganization(id string, org models.UpdateOrganizationInput) (*models.Organization, error) {
	organization, err := a.store.OrganizationStore().FindById(id)
	if err != nil {
		return nil, err
	}
	organization.Description = org.Description
	organization.Logo = org.Logo

	err = a.store.OrganizationStore().Save(organization)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (a App) QueryOrganizations() []*models.Organization {
	organizations, _ := a.store.OrganizationStore().List()

	return organizations
}

func (a App) FindOrganization(id string) (*models.Organization, error) {
	return a.store.OrganizationStore().FindById(id)
}

func (a App) AddProjectToOrg(projectId string, orgId string) {

}
