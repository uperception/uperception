package models

type LighthouseResult struct {
	Project Project
	Url     string
}

func NewLighthouseResult(project Project, url string) *LighthouseResult {
	return &LighthouseResult{
		Project: project,
		Url:     url,
	}
}
