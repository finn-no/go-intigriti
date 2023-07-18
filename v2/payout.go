package v2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (e *Endpoint) GetSubmissionPayouts(submissionCode string) ([]Payout, error) {
	payOutURL := fmt.Sprintf("%s/%s/%s/payouts", e.URLAPI, apiSubmissions, submissionCode)
	req, err := http.NewRequest(http.MethodGet, payOutURL, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create get programs")
	}

	resp, err := e.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "could not get programs")
	}

	if resp.StatusCode > 399 {
		return nil, errors.Errorf("returned status %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read response")
	}

	var payouts []Payout

	if err := json.Unmarshal(b, &payouts); err != nil {
		e.Logger.Error(errors.Wrap(err, "could not decode payouts"))
	}

	return payouts, nil
}

type Payout struct {
	ID          string `json:"id"`
	Originators struct {
		ProgramID       string      `json:"programId"`
		PentestCode     interface{} `json:"pentestCode"`
		SubmissionCode  string      `json:"submissionCode"`
		RewardRequestID interface{} `json:"rewardRequestId"`
	} `json:"originators"`
	Amount struct {
		Value    float64 `json:"value"`
		Currency string  `json:"currency"`
	} `json:"amount"`
	Type struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"type"`
	Researcher struct {
		Ranking struct {
			Rank       int `json:"rank"`
			Reputation int `json:"reputation"`
			Streak     struct {
				ID    int    `json:"id"`
				Value string `json:"value"`
			} `json:"streak"`
		} `json:"ranking"`
		IdentityChecked bool        `json:"identityChecked"`
		UserID          string      `json:"userId"`
		UserName        string      `json:"userName"`
		AvatarURL       interface{} `json:"avatarUrl"`
		Role            string      `json:"role"`
	} `json:"researcher"`
	Status struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"status"`
	CreatedAt     int `json:"createdAt"`
	PaidAt        int `json:"paidAt"`
	LastUpdatedAt int `json:"lastUpdatedAt"`
}
