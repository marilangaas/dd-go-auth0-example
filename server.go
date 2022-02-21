package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"go-auth0-datadog-tracing/config"
	"go-auth0-datadog-tracing/services"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"net/http"
	"time"
)

func main() {
	tracer.Start(tracer.WithService("example"))
	defer tracer.Stop()

	cfg, err := config.NewConfig("config")
	if err != nil {
		log.Fatal().Err(err).Msg("unable to init config")
	}

	ams, _ := services.New(&services.Config{
		Auth0AudienceURL:  cfg.Auth0ManagementServiceSettings.Auth0Domain,
		Auth0ClientID:     cfg.Auth0ManagementServiceSettings.Auth0ClientID,
		Auth0ClientSecret: cfg.Auth0ManagementServiceSettings.Auth0ClientSecret,
		HTTPClient:        createClient(),
	})

	email, err := ams.GetUserEmail("auth0|123")
	if err != nil {
		fmt.Println(fmt.Errorf("%w",err))
	}

	fmt.Println(email)
}

func createClient() *http.Client {

	c := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   4 * time.Second,
	}
	client := httptrace.WrapClient(c,
		httptrace.RTWithServiceName(config.DataDogServiceName),
		httptrace.RTWithResourceNamer(func(r *http.Request) string {
			return r.URL.String()
		}),
	)

	return client
}
