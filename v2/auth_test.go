package v2

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	// intigriti "github.com/hazcod/go-intigriti/pkg/api"
	"github.com/hazcod/go-intigriti/pkg/config"
)

func TestLogin(t *testing.T) {

	clientID := "xx"
	clientSecret := "xx"
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
	fmt.Println(inti.GetSubmissions())
	cc, _ := inti.GetSubmission("SCHIBSTED-WQGZO0J8")
	fmt.Println("CC:", cc)
	assert.Equal(t, err, nil)
}
