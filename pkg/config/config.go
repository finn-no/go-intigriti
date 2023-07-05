package config

import (
	"time"

	"github.com/sirupsen/logrus"
)

type CachedToken struct {
	RefreshToken string
	AccessToken  string
	ExpiryDate   time.Time
	Type         string
}

type InteractiveAuthenticator interface {
	OpenURL(url string) error
}

type Config struct {
	// required authentication credentials
	Credentials struct {
		ClientID     string
		ClientSecret string
	}

	// optionally open a browser to complete authentication if user interaction is required
	OpenBrowser   bool
	Authenticator InteractiveAuthenticator

	// optional token cache if caching previous credentials
	TokenCache *CachedToken

	// optional logger instance
	Logger *logrus.Logger
}
