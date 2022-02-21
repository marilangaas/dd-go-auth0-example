package config

import (
	"github.com/spf13/viper"
)

var (
	DataDogServiceName = "idp-user-service-api"
)

type Config struct {
	Auth0ManagementServiceSettings *Auth0ManagementServiceConfig
}

func NewConfig(cfgPath string) (*Config, error) {
	if err := initEnv(cfgPath); err != nil {
		return nil, err
	}
	c := &Config{
		Auth0ManagementServiceSettings:       NewAuth0ManagementServiceSettings(),
	}
	return c, nil
}

func initEnv(cfgPath string) error {
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cfgPath)
	return viper.ReadInConfig()
}

type Auth0ManagementServiceConfig struct {
	Auth0Domain   string
	Auth0ClientID string
	Auth0ClientSecret string
}

func NewAuth0ManagementServiceSettings() *Auth0ManagementServiceConfig {
	return &Auth0ManagementServiceConfig{
		Auth0ClientID:     viper.GetString("auth0_client_id"),
		Auth0ClientSecret: viper.GetString("auth0_client_secret"),
		Auth0Domain:       viper.GetString("auth0_domain"),
	}
}

