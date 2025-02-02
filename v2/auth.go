package v2

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/hazcod/go-intigriti/pkg/config"

	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

const (
	// timeout of every http request
	httpTimeoutSec = 5
	// the length of our Oauth2 state parameter
	stateLengthLetters = 10
	// timeout of the local callback listener
	callbackTimeoutSec = 120

	// local callback url listener
	localCallbackPort = 1337
	localCallbackHost = "localhost"
	localCallbackURI  = "/"

	// default production API endpoints
	defaultApiTokenURL = "https://login.intigriti.com/connect/token"
	defaultApiAuthzURL = "https://login.intigriti.com/connect/authorize"
	defaultApiEndpoint = "https://api.intigriti.com/external"
)

var (
	// used to override the API endpoints at runtime for testing
	tokenURL = os.Getenv("INTI_TOKEN_URL")
	authzURL = os.Getenv("INTI_AUTH_URL")
	apiURL   = os.Getenv("INTI_API_URL")
)

// used if we do local testing to non-production endpoints
func init() {
	if tokenURL == "" {
		tokenURL = defaultApiTokenURL
	}

	if authzURL == "" {
		authzURL = defaultApiAuthzURL
	}

	if apiURL == "" {
		apiURL = defaultApiEndpoint
	}
}

// retrieve the oauth2 configuration to use
func (e *Endpoint) GetOauth2Config(apiScopes []string) oauth2.Config {
	e.Logger.WithField("api_url", apiURL).Debug("set api url")

	oauthConfig := oauth2.Config{
		ClientID:     e.clientID,
		ClientSecret: e.clientSecret,
		Endpoint: oauth2.Endpoint{
			TokenURL: tokenURL,
			AuthURL:  authzURL,
		},
		RedirectURL: fmt.Sprintf("http://%s:%d%s", localCallbackHost, localCallbackPort, localCallbackURI),
		Scopes:      apiScopes,
	}
	e.URLAPI = apiURL
	e.Logger.Tracef("%+v", oauthConfig)

	return oauthConfig
}

// fetch the latest (valid) oauth2 access and refresh token
func (e *Endpoint) GetToken() (*oauth2.Token, error) {
	// don't do anything when the token is ok
	if e.OauthToken != nil && e.OauthToken.Valid() {
		return e.OauthToken, nil
	}

	// get out oauth2 config to use
	conf := e.GetOauth2Config(e.ApiScopes)

	// get valid refresh and access tokens
	tokenSrc := conf.TokenSource(context.Background(), e.OauthToken)
	token, err := tokenSrc.Token()
	if err != nil {
		return nil, errors.Wrap(err, "could not retrieve refresh token")
	}

	return token, nil
}

// return the http client which automatically injects the right authentication credentials
func (e *Endpoint) GetClient(tc *config.CachedToken, auth *config.InteractiveAuthenticator) (*http.Client, error) {
	ctx := context.Background()

	conf := e.GetOauth2Config(e.ApiScopes)

	httpClient := &http.Client{Timeout: httpTimeoutSec * time.Second}

	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)

	e.OauthToken = &oauth2.Token{}

	if tc == nil {
		tc = &config.CachedToken{}
	}

	// if our configuration contains a cached token, re-use it
	if tc.RefreshToken != "" {
		e.Logger.Debug("trying to use cached token")
		e.OauthToken.AccessToken = tc.AccessToken
		e.OauthToken.RefreshToken = tc.RefreshToken
		e.OauthToken.Expiry = tc.ExpiryDate
		e.OauthToken.TokenType = tc.Type
	}

	// if the current token is invalid, fetch a new one
	if !e.OauthToken.Valid() {
		e.Logger.Debug("authenticating for new token")

		authzCode, err := e.authenticate(ctx, &conf, auth)
		if err != nil {
			return nil, errors.Wrap(err, "failed to authenticate")
		}

		e.Logger.WithField("code", authzCode).Debug("exchanging code")

		e.OauthToken, err = conf.Exchange(ctx, authzCode)
		if err != nil {
			return nil, errors.Wrap(err, "could not exchange code")
		}
	}

	// ensure our http client uses our oauth2 credentials
	authHttpClient := conf.Client(ctx, e.OauthToken)

	// inject a logging middleware into our http client
	authHttpClient.Transport = TaggedRoundTripper{Proxied: authHttpClient.Transport, Logger: e.Logger}
	e.Logger.Debug("successfully created client")

	return authHttpClient, nil
}

// authenticate versus the Intigriti API, this requires user interaction
func (e *Endpoint) authenticate(ctx context.Context, oauth2Config *oauth2.Config, auth *config.InteractiveAuthenticator) (string, error) {
	state := randomString(stateLengthLetters)

	resultChan := make(chan callbackResult, 1)

	ctx, cancel := context.WithTimeout(ctx, time.Second*callbackTimeoutSec)

	go e.listenForCallback(localCallbackURI, localCallbackHost, localCallbackPort, state, resultChan)
	defer func() { go func() { cancel(); resultChan <- callbackResult{} }() }()

	// Redirect user to consent page to ask for permission 	for the scopes specified above.
	url := oauth2Config.AuthCodeURL(state, oauth2.AccessTypeOffline)
	e.Logger.Warnf("Please authenticate: %s", url)

	if auth != nil {
		e.Logger.Info("opening system browser to authenticate")

		authenticator := *auth
		if err := authenticator.OpenURL(url); err != nil {
			e.Logger.WithField("url", url).WithError(err).Warnf("could not open browser")
		}
	}

	e.Logger.Debug("waiting for callback click")

	var chanResult callbackResult
	select {
	case <-ctx.Done():
		resultChan <- callbackResult{Error: errors.New("timeout")}
		break
	case chanResult = <-resultChan:
		break
	}

	e.Logger.WithField("result", chanResult).Debug("received callback result")

	if chanResult.Error != nil {
		return "", chanResult.Error
	}

	if chanResult.Code == "" {
		return "", errors.New("got empty code")
	}

	e.Logger.WithField("code", chanResult.Code).Debug("successfully retrieved new code")
	return chanResult.Code, nil
}
