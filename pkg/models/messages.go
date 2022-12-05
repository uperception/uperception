package models

type RunLighthouseProjectMessage struct {
	ProjectID uint
}

type RunLighthouseEndpointMessage struct {
	ProjectID  uint
	EndpointID uint
}
