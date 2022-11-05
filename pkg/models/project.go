package models

type Project struct {
	Domain string
	Urls   []string
}

func NewProject(domain string, urls []string) *Project {
	return &Project{
		Urls:   urls,
		Domain: domain,
	}
}

func (p Project) HasAuthentication() bool {
	return false
}
