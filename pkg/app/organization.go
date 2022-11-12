package app

import "github.com/leometzger/mmonitoring/pkg/models"

func (a App) CreateOrganization(org models.CreateOrganizationInput) (*models.Organization, error) {
	organization := models.Organization{
		Name:        org.Name,
		Description: org.Description,
	}

	err := a.organizationStore.Save(&organization)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}

func (a App) UpdateOrganization(id string, org models.UpdateOrganizationInput) (*models.Organization, error) {
	organization, err := a.organizationStore.FindById(id)
	if err != nil {
		return nil, err
	}
	organization.Description = org.Description
	organization.Logo = org.Logo

	err = a.organizationStore.Save(organization)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (a App) QueryOrganizations() []*models.Organization {
	organizations, _ := a.organizationStore.List()

	return organizations
}

func (a App) FindOrganization(id string) (*models.Organization, error) {
	return a.organizationStore.FindById(id)
}

func (a App) DeleteOrganization(id string) error {
	return a.organizationStore.Delete(id)
}

func (a App) AddProjectToOrg(projectId string, orgId string) {

}
