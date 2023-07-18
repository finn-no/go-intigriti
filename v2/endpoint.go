package v2

import (
	"net/http"
	"strings"

	"github.com/hazcod/go-intigriti/pkg/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

const (
	apiAllScopes = "offline_access company_external_api core_platform:read"
)

type Endpoint struct {
	Logger *logrus.Logger

	URLAPI string

	clientID     string
	clientSecret string
	clientTag    string

	Client     *http.Client
	OauthToken *oauth2.Token

	ApiScopes []string
}

// New creates an Intigriti endpoint object to use
// this is the main object to interact with the SDK
func New(cfg config.Config) (Endpoint, error) {
	e := Endpoint{
		clientID:     cfg.Credentials.ClientID,
		clientSecret: cfg.Credentials.ClientSecret,
		clientTag:    clientTag,
		ApiScopes:    cfg.APIScopes,
	}

	if len(e.ApiScopes) == 0 {
		e.ApiScopes = strings.Split(apiAllScopes, " ")
	}

	// initialize the logger to use
	if cfg.Logger == nil {
		e.Logger = logrus.New()
	} else {
		e.Logger = cfg.Logger
	}

	// prepare our oauth2-ed http client
	authenticator := &cfg.Authenticator
	if !cfg.OpenBrowser {
		authenticator = nil
	}

	httpClient, err := e.GetClient(cfg.TokenCache, authenticator)
	if err != nil {
		return e, errors.Wrap(err, "could not init client")
	}

	e.Client = httpClient

	// ensure our current token is fetched or renewed if expired
	if _, err = e.GetToken(); err != nil {
		return e, errors.Wrap(err, "could not prepare token")
	}

	return e, nil
}

// IsAuthenticated returns whether the current SDK instance has successfully authenticated
func (e *Endpoint) IsAuthenticated() bool {
	if e.OauthToken == nil {
		return false
	}

	return e.OauthToken.Valid()
}
