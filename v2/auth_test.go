package v2

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	// intigriti "github.com/hazcod/go-intigriti/pkg/api"
	"github.com/hazcod/go-intigriti/pkg/config"
)

func TestLogin(t *testing.T) {

	clientID := os.Getenv("INTIGRITI_CLIENT_ID_V2")
	clientSecret := os.Getenv("INTIGRITI_CLIENT_SECRET_V2")
	fmt.Println(clientID, clientSecret)
	t.Log(clientID, clientSecret)
	t.Log(clientSecret, clientID)
	inti, err := New(config.Config{
		Credentials: struct {
			ClientID     string
			ClientSecret string
		}{
			ClientID:     clientID,
			ClientSecret: clientSecret,
		},
	})
	if err != nil {
		t.Log(err)
	}
	fmt.Println("Is auth:", inti.IsAuthenticated())
	assert.Equal(t, err, nil)
}
