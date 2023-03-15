package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type submissionsResponse []struct {
	Code              string `json:"code"`
	InternalReference struct {
		Reference string `json:"reference"`
		URL       string `json:"url"`
	} `json:"internalReference"`
	Title         string `json:"title"`
	OriginalTitle string `json:"originalTitle"`
	Program       struct {
		Name                 string `json:"name"`
		Handle               string `json:"handle"`
		LogoURL              string `json:"logoUrl"`
		ConfidentialityLevel struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"confidentialityLevel"`
		Status struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"status"`
		StatusTrigger struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"statusTrigger"`
	} `json:"program"`
	Type struct {
		Name     string `json:"name"`
		Category string `json:"category"`
		Cwe      string `json:"cwe"`
	} `json:"type"`
	Severity struct {
		ID     int    `json:"id"`
		Vector string `json:"vector"`
		Value  string `json:"value"`
	} `json:"severity"`
	Domain struct {
		Value      string `json:"value"`
		Motivation string `json:"motivation"`
		Type       struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"type"`
		BountyTable struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"bountyTable"`
	} `json:"domain"`
	EndpointVulnerableComponent string `json:"endpointVulnerableComponent"`
	State                       struct {
		Status struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"status"`
		CloseReason struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"closeReason"`
	} `json:"state"`
	TotalPayout      float64 `json:"totalPayout"`
	CreatedAt        int64   `json:"createdAt"`
	LastUpdatedAt    int64   `json:"lastUpdatedAt"`
	ClosedAt         int64   `json:"closedAt"`
	ValidatedAt      int64   `json:"validatedAt"`
	AcceptedAt       int64   `json:"acceptedAt"`
	ArchivedAt       int64   `json:"archivedAt"`
	AwaitingFeedback bool    `json:"awaitingFeedback"`
	Assignee         struct {
		UserName  string `json:"userName"`
		AvatarURL string `json:"avatarUrl"`
		Email     string `json:"email"`
		Role      string `json:"role"`
	} `json:"assignee"`
	Researcher struct {
		UserName  string `json:"userName"`
		AvatarURL string `json:"avatarUrl"`
		Ranking   struct {
			Rank       int `json:"rank"`
			Reputation int `json:"reputation"`
			Streak     struct {
				ID    int    `json:"id"`
				Value string `json:"value"`
			} `json:"streak"`
		} `json:"ranking"`
		IdentityChecked bool `json:"identityChecked"`
	} `json:"researcher"`
	LastUpdater struct {
		UserName  string `json:"userName"`
		AvatarURL string `json:"avatarUrl"`
		Email     string `json:"email"`
		Role      string `json:"role"`
	} `json:"lastUpdater"`
	Links struct {
		Details string `json:"details"`
	} `json:"links"`
	WebLinks struct {
		Details string `json:"details"`
	} `json:"webLinks"`
	SubmissionDetailURL string `json:"submissionDetailUrl"`
}

type Researcher struct {
	Username  string
	AvatarURL string
}

type Program struct {
	Handle string
	Name   string
}

type Submission struct {
	Program    Program
	Researcher Researcher

	DateLastUpdated time.Time
	DateCreated     time.Time
	DateClosed      time.Time

	CWE      string
	Type     string
	Category string

	InternalReference string

	ID       string
	URL      string
	Title    string
	Severity string
	Endpoint string
	State    string
	Payout   float64

	CloseReason string
}

/*
	INTIGRITI Report structure
*/
type SubmissionDetails struct {
	Submission        Submission
	Code              string `json:"code"`
	URL               string
	InternalReference struct {
		Reference string `json:"reference"`
		URL       string `json:"url"`
	} `json:"internalReference"`
	Title         string `json:"title"`
	OriginalTitle string `json:"originalTitle"`
	Type          struct {
		Name     string `json:"name"`
		Category string `json:"category"`
		Cwe      string `json:"cwe"`
	} `json:"type"`
	Questions []struct {
		Question string `json:"question"`
		Type     struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"type"`
		Answer string `json:"answer"`
	} `json:"questions"`
	Program struct {
		ID                   string `json:"id"`
		Name                 string `json:"name"`
		Handle               string `json:"handle"`
		LogoURL              string `json:"logoUrl"`
		ConfidentialityLevel struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"confidentialityLevel"`
		Status struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"status"`
		StatusTrigger struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"statusTrigger"`
	} `json:"program"`
	Domain struct {
		Value      string `json:"value"`
		Motivation string `json:"motivation"`
		Type       struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"type"`
		BountyTable struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"bountyTable"`
		Description string `json:"description"`
	} `json:"domain"`
	EndpointVulnerableComponent string `json:"endpointVulnerableComponent"`
	PocDescription              string `json:"pocDescription"`
	Impact                      string `json:"impact"`
	PersonalData                bool   `json:"personalData"`
	RecommendedSolution         string `json:"recommendedSolution"`
	State                       struct {
		Status struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"status"`
		CloseReason struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"closeReason"`
		DuplicateSubmissionURL string `json:"duplicateSubmissionUrl"`
	} `json:"state"`
	DuplicatedSubmissionUrls []string `json:"duplicatedSubmissionUrls"`
	Severity                 struct {
		ID     int    `json:"id"`
		Vector string `json:"vector"`
		Value  string `json:"value"`
	} `json:"severity"`
	AwaitingFeedback bool `json:"awaitingFeedback"`
	DaysOpen         int  `json:"daysOpen"`
	Payouts          []struct {
		Amount    interface{} `json:"amount"`
		CreatedAt interface{} `json:"createdAt"`
		Type      struct {
			ID    interface{} `json:"id"`
			Value string      `json:"value"`
		} `json:"type"`
	} `json:"payouts"`
	TotalPayout       interface{} `json:"totalPayout"`
	TotalBountyPayout interface{} `json:"totalBountyPayout"`
	TotalBonusPayout  interface{} `json:"totalBonusPayout"`
	CreatedAt         int         `json:"createdAt"`
	LastUpdatedAt     int         `json:"lastUpdatedAt"`
	ValidatedAt       int         `json:"validatedAt"`
	AcceptedAt        int         `json:"acceptedAt"`
	ClosedAt          int         `json:"closedAt"`
	ArchivedAt        int         `json:"archivedAt"`
	Destroyed         struct {
		DestroyedBy struct {
			UserID    string `json:"userId"`
			UserName  string `json:"userName"`
			AvatarURL string `json:"avatarUrl"`
			Email     string `json:"email"`
			Role      string `json:"role"`
		} `json:"destroyedBy"`
		DestroyedAt string `json:"destroyedAt"`
	} `json:"destroyed"`
	Assignee struct {
		UserID    string `json:"userId"`
		UserName  string `json:"userName"`
		AvatarURL string `json:"avatarUrl"`
		Email     string `json:"email"`
		Role      string `json:"role"`
	} `json:"assignee"`
	Tags           []string     `json:"tags"`
	Attachments    []Attachment `json:"attachments"`
	AttachmentUrls []string     `json:"attachmentUrls"`
	GroupLegacy    struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"groupLegacy"`
	Researcher struct {
		UserID    string `json:"userId"`
		UserName  string `json:"userName"`
		AvatarURL string `json:"avatarUrl"`
		Email     string `json:"email"`
		Role      string `json:"role"`
		Ranking   struct {
			Rank       interface{} `json:"rank"`
			Reputation interface{} `json:"reputation"`
			Streak     struct {
				ID    interface{} `json:"id"`
				Value string      `json:"value"`
			} `json:"streak"`
		} `json:"ranking"`
		IdentityChecked bool `json:"identityChecked"`
	} `json:"researcher"`
	LastUpdater struct {
		UserID    string `json:"userId"`
		UserName  string `json:"userName"`
		AvatarURL string `json:"avatarUrl"`
		Email     string `json:"email"`
		Role      string `json:"role"`
	} `json:"lastUpdater"`
	Files  []string `json:"files"`
	Events []struct {
		Type struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"type"`
		CreatedAt  int `json:"createdAt"`
		Visibility struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"visibility"`
		User struct {
			UserID    string      `json:"userId"`
			UserName  string      `json:"userName"`
			AvatarURL string      `json:"avatarUrl"`
			Email     interface{} `json:"email"`
			Role      string      `json:"role"`
		} `json:"user"`
		From struct {
			Status struct {
				ID    int    `json:"id"`
				Value string `json:"value"`
			} `json:"status"`
			CloseReason            interface{} `json:"closeReason"`
			DuplicateSubmissionURL string      `json:"duplicateSubmissionUrl"`
			UserID                 string      `json:"userId"`
			UserName               string      `json:"userName"`
			AvatarURL              string      `json:"avatarUrl"`
			Email                  string      `json:"email"`
			Role                   string      `json:"role"`
		} `json:"from,omitempty"`
		To struct {
			Status struct {
				ID    int    `json:"id"`
				Value string `json:"value"`
			} `json:"status"`
			CloseReason            interface{} `json:"closeReason"`
			DuplicateSubmissionURL string      `json:"duplicateSubmissionUrl"`
			UserID                 string      `json:"userId"`
			UserName               string      `json:"userName"`
			AvatarURL              string      `json:"avatarUrl"`
			Email                  string      `json:"email"`
			Role                   string      `json:"role"`
		} `json:"to,omitempty"`
		Amount                 float64      `json:"amount"`
		Message                interface{}  `json:"message"`
		Attachments            []Attachment `json:"attachments"`
		AttachmentUrls         []string     `json:"attachmentUrls"`
		DestroyMessageMetadata interface{}  `json:"destroyMessageMetadata"`
	} `json:"events"`
	WebLinks struct {
		Details string `json:"details"`
	} `json:"webLinks"`
	IntegrationReferences []struct {
		Type struct {
			ID    int    `json:"id"`
			Value string `json:"value"`
		} `json:"type"`
		Reference string `json:"reference"`
		URL       string `json:"url"`
	} `json:"integrationReferences"`
}

/*
	Attachment structure used in Submission
*/
type Attachment struct {
	ID             string `json:"id"`
	LinkID         string `json:"linkId"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	IsSafe         bool   `json:"isSafe"`
	Size           int    `json:"size"`
	CreatedAt      int    `json:"createdAt"`
	Code           int    `json:"code"`
	AttachmentType int    `json:"attachmentType"`
}

func (s *Submission) IsReady() bool {
	switch strings.ToLower(s.State) {
	case "triage":
		return false
	case "closed":
		return false
	case "accepted":
		return false
	case "archived":
		return false
	default:
		return true
	}
}

func (s *Submission) IsClosed() bool {
	return s.CloseReason != ""
}

func getBearerTokenHeader(authToken string) string {
	return "Bearer " + authToken
}

func (e *Endpoint) GetSubmissions() ([]Submission, error) {
	var findings []Submission

	if err := authenticate(e); err != nil {
		return findings, errors.Wrap(err, "could not authenticate to intigriti API")
	}

	// req, err := http.NewRequest(http.MethodGet, e.URLApiSubmissions, nil) //ORIGINAL
	req, err := http.NewRequest(http.MethodGet, apiEndpointV1, nil)
	if err != nil {
		return findings, errors.Wrap(err, "could not create http request to intigriti")
	}

	req.Header.Set("Content-Type", mimeFormUrlEncoded)
	req.Header.Set("X-Client", clientTag)
	req.Header.Set("Authorization", getBearerTokenHeader(e.authToken))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return findings, errors.Wrap(err, "fetching to intigriti failed")
	}

	defer func() { _ = resp.Body.Close() }()

	// the token was invalidated before it expired
	if resp.StatusCode == http.StatusUnauthorized {
		// fetch a new token
		if err := authenticate(e); err != nil {
			return nil, errors.Wrap(err, "could not reauthenticate to intigriti API")
		}

		e.Logger.Debug("reauthenticated because of an invalidated token")
	}

	if resp.StatusCode > 399 {
		return findings, errors.Errorf("fetch from intigriti returned status code: %d", resp.StatusCode)
	}

	var fetchResp submissionsResponse
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return findings, errors.Wrap(err, "could not read response")
	}

	if err := json.Unmarshal(respBytes, &fetchResp); err != nil {
		return findings, errors.Wrap(err, "could not decode intigriti response")
	}

	for _, entry := range fetchResp {
		findings = append(findings, Submission{
			Program: Program{
				Handle: entry.Program.Handle,
				Name:   entry.Program.Name,
			},

			Researcher: Researcher{
				Username:  entry.Researcher.UserName,
				AvatarURL: entry.Researcher.AvatarURL,
			},

			State:    entry.State.Status.Value,
			Type:     entry.Type.Name,
			CWE:      entry.Type.Cwe,
			Category: entry.Type.Category,
			ID:       entry.Code,

			URL:      entry.WebLinks.Details,
			Title:    entry.Title,
			Severity: entry.Severity.Value,

			DateCreated:     time.Unix(entry.CreatedAt, 0),
			DateClosed:      time.Unix(entry.ClosedAt, 0),
			DateLastUpdated: time.Unix(entry.LastUpdatedAt, 0),

			Endpoint: entry.EndpointVulnerableComponent,

			InternalReference: entry.InternalReference.Reference,
			CloseReason:       entry.State.CloseReason.Value,
		})
	}

	return findings, nil
}

/*
	GetSubmission gets report/submission by code/id
*/
func (e *Endpoint) GetSubmission(code string) (*SubmissionDetails, error) {
	var submi SubmissionDetails
	var respBytes []byte
	var err error
	var req *http.Request
	url := apiEndpointV1 + "/SCHIBSTED-LKJ2GIWL"

	if err := authenticate(e); err != nil {
		return nil, errors.Wrap(err, "could not authenticate to intigriti API")
	}

	req, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.Wrap(err, "could not create http request to intigriti")
	}
	req.Header.Set("Content-Type", mimeFormUrlEncoded)
	req.Header.Set("X-Client", clientTag)
	req.Header.Set("Authorization", getBearerTokenHeader(e.authToken))

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "fetching to intigriti failed")
	}

	defer func() { _ = resp.Body.Close() }()

	// the token was invalidated before it expired
	if resp.StatusCode == http.StatusUnauthorized {
		// fetch a new token
		if err := authenticate(e); err != nil {
			return nil, errors.Wrap(err, "could not reauthenticate to intigriti API")
		}
		e.Logger.Debug("reauthenticated because of an invalidated token")
	}

	if resp.StatusCode > 399 {
		return nil, errors.Errorf("fetch from intigriti returned status code: %d", resp.StatusCode)
	}

	respBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "could not read response")
	}
	err = json.Unmarshal(respBytes, &submi)
	if err != nil && submi.Code == "" {
		return nil, err
	}
	return &submi, nil
}
