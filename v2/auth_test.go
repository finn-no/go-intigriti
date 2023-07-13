package v2

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
	// intigriti "github.com/hazcod/go-intigriti/pkg/api"
	"github.com/hazcod/go-intigriti/pkg/config"
)

func TestLogin(t *testing.T) {

	clientID := "yy"
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
	// fmt.Println(inti.GetSubmissions())
	// cc, _ := inti.GetSubmission("SCHIBSTED-WQGZO0J8")
	p, err := inti.GetProgram("f65f73f6-e60f-405e-a7fd-4674ecb618bc")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("CC:", cc)
	fmt.Println("PG:", p)
	assert.Equal(t, err, nil)
}
