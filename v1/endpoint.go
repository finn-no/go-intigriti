package v1

import (
	"time"

	"github.com/sirupsen/logrus"
)

const (
	apiSubmissions = "https://api.intigriti.com/external/submission"
	apiAuth        = "https://login.intigriti.com/connect/token"
	apiEndpointV1  = "https://api.intigriti.com/external/v1.2/submissions"
	clientTag      = "Go intigriti library"
)

type Endpoint struct {
	Logger            *logrus.Logger
	URLApiAuth        string
	URLApiSubmissions string

	clientToken  string
	clientSecret string
	clientTag    string

	authToken    string
	authTokenExp time.Time
}

func New(clientToken string, clientSecret string) Endpoint {
	e := Endpoint{
		clientToken:  clientToken,
		clientSecret: clientSecret,
		clientTag:    clientTag,
	}

	e.Logger = logrus.New()
	e.URLApiAuth = apiAuth
	e.URLApiSubmissions = apiSubmissions

	return e
}
