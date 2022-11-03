package utils

import "github.com/spf13/viper"

type Config struct {
	Region string `mapstructure:"AWS_REGION"`
	Queue  string `mapstructure:"AWS_SQS_ARN"`
	Bucket string `mapstructure:"AWS_S3_BUCKET"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
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
