package v2

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	programURI = "company/v2/programs"
)

func (e *Endpoint) GetPrograms() ([]Program, error) {
	programsUrl := fmt.Sprintf("%s/%s", e.URLAPI, programURI)
	req, err := http.NewRequest(http.MethodGet, programsUrl, nil)
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

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read response")
	}

	var programs []Program

	if err := json.Unmarshal(b, &programs); err != nil {
		e.Logger.Error(errors.Wrap(err, "could not decode programs"))
	}

	return programs, nil
}

func (e *Endpoint) GetProgram(id string) (Program, error) {
	programUrl := fmt.Sprintf("%s/%s/%s", e.URLAPI, programURI, id)
	req, err := http.NewRequest(http.MethodGet, programUrl, nil)
	fmt.Println("Final url:", programUrl)
	if err != nil {
		return Program{}, errors.Wrap(err, "could not create get programs")
	}

	resp, err := e.Client.Do(req)
	if err != nil {
		return Program{}, errors.Wrap(err, "could not get programs")
	}

	if resp.StatusCode > 399 {
		return Program{}, errors.Errorf("returned status %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Program{}, errors.Wrap(err, "could not read response")
	}

	var program Program

	if err := json.Unmarshal(b, &program); err != nil {
		e.Logger.Error(errors.Wrap(err, "could not decode program"))
	}

	return program, nil
}

type ProgramOld struct {
	ID            string `json:"id"`
	Handle        string `json:"handle"`
	CompanyID     string `json:"companyId"`
	CompanyHandle string `json:"companyHandle"`
	LogoURL       string `json:"logoUrl"`
	Name          string `json:"name"`
	Status        struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"status"`
	ConfidentialityLevel struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"confidentialityLevel"`
	WebLinks struct {
		Details string `json:"details"`
	} `json:"webLinks"`
	Type struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"type"`
}
type Program struct {
	ProgramID     string `json:"programId"`
	Handle        string `json:"handle"`
	Name          string `json:"name"`
	CompanyID     string `json:"companyId"`
	CompanyHandle string `json:"companyHandle"`
	Description   string `json:"description"`
	Options       []struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"options"`
	State struct {
		Status struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"status"`
		LastStatusTriggerID struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"lastStatusTriggerId"`
	} `json:"state"`
	ConfidentialityLevel struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"confidentialityLevel"`
	MaxConfidentialityLevel struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"maxConfidentialityLevel"`
	LogoURL string `json:"logoUrl"`
	Domain  []struct {
		ID   string `json:"id"`
		Type struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"type"`
		Endpoint string `json:"endpoint"`
		Tier     struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"tier"`
		Description string `json:"description"`
	} `json:"domain"`
	InScope             interface{}   `json:"inScope"`
	OutOfScope          interface{}   `json:"outOfScope"`
	Faq                 interface{}   `json:"faq"`
	SeverityAssessment  interface{}   `json:"severityAssessment"`
	RulesOfEngagement   interface{}   `json:"rulesOfEngagement"`
	SubmissionQuestions []interface{} `json:"submissionQuestions"`
	Bounties            []interface{} `json:"bounties"`
	AttachmentCount     int           `json:"attachmentCount"`
	BugBounty           struct {
		AutoSuspendThreshold struct {
			Value    float64 `json:"value"`
			Currency string  `json:"currency"`
		} `json:"autoSuspendThreshold"`
		LeaderboardVisibility struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"leaderboardVisibility"`
	} `json:"bugBounty"`
	Type struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
	} `json:"type"`
	ProgramBudget                interface{} `json:"programBudget"`
	TacRequired                  bool        `json:"tacRequired"`
	TacEmpty                     bool        `json:"tacEmpty"`
	IdentityCheckedRequired      bool        `json:"identityCheckedRequired"`
	SkipTriage                   bool        `json:"skipTriage"`
	AwardRep                     bool        `json:"awardRep"`
	AllowResearcherCollaboration bool        `json:"allowResearcherCollaboration"`
	DefaultAssignee              struct {
		UserID    string `json:"userId"`
		UserName  string `json:"userName"`
		AvatarURL string `json:"avatarUrl"`
		Role      string `json:"role"`
		Email     string `json:"email"`
	} `json:"defaultAssignee"`
	CreatedAt     int `json:"createdAt"`
	LastUpdatedAt int `json:"lastUpdatedAt"`
	WebLinks      struct {
		Details     string `json:"details"`
		Submissions string `json:"submissions"`
	} `json:"webLinks"`
}
