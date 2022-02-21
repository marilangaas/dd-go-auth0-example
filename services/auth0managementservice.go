package services

import (
	"github.com/auth0/go-auth0/management"
	"github.com/rs/zerolog/log"
	"net/http"
)

//Auth0ManagementService is a service to perform operations against the auth0 management API
type Auth0ManagementService struct {
	mgmtClient *management.Management
}

//Config wraps the config for the Auth0ManagementService
type Config struct {
	Auth0AudienceURL  string
	Auth0ClientID     string
	Auth0ClientSecret string
	HTTPClient        *http.Client
}

func New(cfg *Config) (*Auth0ManagementService, error) {
	if cfg.HTTPClient == nil {
		log.Fatal().Msg("failure initializing http client in auth0managementservice")
	}
	m, err := management.New(cfg.Auth0AudienceURL, management.WithClientCredentials(cfg.Auth0ClientID, cfg.Auth0ClientSecret), management.WithClient(cfg.HTTPClient))
	if err != nil {
		return nil, err
	}
	return &Auth0ManagementService{mgmtClient: m}, nil
}

func (ams *Auth0ManagementService) GetUserEmail(id string) (*management.User, error) {
	user, err := ams.mgmtClient.User.Read(id, management.IncludeFields("email"))
	if err != nil {
		return nil, err
	}
	return user, nil
}
