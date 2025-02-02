package v2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

func (e *Endpoint) GetSubmissionEvents(submissionCode string) ([]Event, error) {
	payOutURL := fmt.Sprintf("%s/%s/%s/events", e.URLAPI, apiSubmissions, submissionCode)
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

	var events []Event

	if err := json.Unmarshal(b, &events); err != nil {
		e.Logger.Error(errors.Wrap(err, "could not decode events"))
	}

	return events, nil
}

type Event struct {
	Type struct {
		ID    int    `json:"id,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"type,omitempty"`
	From struct {
		Status struct {
			ID    int    `json:"id,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"status,omitempty"`
		CloseReason            interface{} `json:"closeReason,omitempty"`
		DuplicateSubmissionURL string      `json:"duplicateSubmissionUrl,omitempty"`
		UserID                 string      `json:"userId,omitempty"`
		UserName               string      `json:"userName,omitempty"`
		AvatarURL              string      `json:"avatarUrl,omitempty"`
		Email                  string      `json:"email,omitempty"`
		Role                   string      `json:"role,omitempty"`
		DuplicateInfo          struct {
			ParentSubmissionCode interface{} `json:"parentSubmissionCode,omitempty"`
			ChildSubmissionCodes interface{} `json:"childSubmissionCodes,omitempty"`
		} `json:"duplicateInfo,omitempty"`
	} `json:"from,omitempty"`
	To struct {
		Status struct {
			ID    int    `json:"id,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"status,omitempty"`
		CloseReason            interface{} `json:"closeReason,omitempty"`
		DuplicateSubmissionURL string      `json:"duplicateSubmissionUrl,omitempty"`
		UserID                 string      `json:"userId,omitempty"`
		UserName               string      `json:"userName,omitempty"`
		AvatarURL              string      `json:"avatarUrl,omitempty"`
		Email                  string      `json:"email,omitempty"`
		Role                   string      `json:"role,omitempty"`
		DuplicateInfo          struct {
			ParentSubmissionCode interface{} `json:"parentSubmissionCode,omitempty"`
			ChildSubmissionCodes interface{} `json:"childSubmissionCodes,omitempty"`
		} `json:"duplicateInfo,omitempty"`
	} `json:"to,omitempty"`
	Message                string       `json:"message,omitempty"`
	Attachments            []Attachment `json:"attachments,omitempty"`
	DestroyMessageMetadata interface{}  `json:"destroyMessageMetadata,omitempty"`
	LastEditedAt           int64        `json:"lastEditedAt,omitempty"`
	DeletedAt              int64        `json:"deletedAt,omitempty"`
	CreatedAt              int64        `json:"createdAt"`
	Visibility             interface{}  `json:"visibility"`
	User                   User         `json:"user,omitempty"`
	Amount                 struct {
		Value    float64 `json:"value,omitempty"`
		Currency string  `json:"currency,omitempty"`
	} `json:"amount,omitempty"`
	PayoutType struct {
		ID    int    `json:"id,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"payoutType"`
	PayoutID string `json:"payoutId,omitempty"`
}
