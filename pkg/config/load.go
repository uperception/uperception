package config

import "github.com/spf13/viper"

type Config struct {
	Region         string `mapstructure:"AWS_REGION"`
	Queue          string `mapstructure:"AWS_SQS_ARN"`
	QueueUrl       string `mapstructure:"AWS_SQS_URL"`
	Bucket         string `mapstructure:"AWS_S3_BUCKET"`
	DBDatabase     string `mapstructure:"DB_DATABASE"`
	DBUser         string `mapstructure:"DB_USER"`
	DBPassword     string `mapstructure:"DB_PASSWORD"`
	DBHost         string `mapstructure:"DB_HOST"`
	DBPort         string `mapstructure:"DB_PORT"`
	KeycloakRealm  string `mapstructure:"KEYCLOAK_REALM"`
	KeycloakUrl    string `mapstructure:"KEYCLOAK_URL"`
	KeycloakSecret string `mapstructure:"KEYCLOAK_SECRET"`
	KeycloakClient string `mapstructure:"KEYCLOAK_CLIENT"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
