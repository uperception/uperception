package models

type LighthouseConfig struct {
	ID          uint `gorm:"primary_key"`
	Enabled     bool
	Periodicity uint8
	Endpoints   []LighthouseEndpoint
}

type UpdateLighthouseConfigInput struct {
	Enabled     bool
	Periodicity uint8
	Endpoints   []LighthouseEndpoint
}

type LighthouseEndpoint struct {
	ID     uint `gorm:"primary_key"`
	Url    string
	Header string
}
