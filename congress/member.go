package congress

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Member holds the data of a member of Congress
type Member struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	ShortTitle   string `json:"short_title"`
	FirstName    string `json:"first_name"`
	MiddleName   string `json:"middle_name,omitempty"`
	LastName     string `json:"last_name"`
	Suffix       string `json:"suffix,omitempty"`
	Birth        string `json:"date_of_birth"`
	Party        string `json:"party"`
	State        string `json:"state,omitempty"`
	SenateClass  string `json:"senate_class,omitempty"`
	StateRank    string `json:"state_rank,omitempty"`
	Leadership   string `json:"leadership_role,omitempty"`
	InOffice     bool   `json:"in_office"`
	Seniority    string `json:"seniority,omitempty"`
	NextElection string `json:"next_election,omitempty"`
	TotalVotes   int    `json:"total_votes,omitempty"`
	MissedVotes  int    `json:"missed_votes,omitempty"`
	PresentVotes int    `json:"present_votes,omitempty"`

	// *tk what's this?
	OCD string `json:"ocd_id,omitempty"`

	Office string `json:"office,omitempty"`
	Phone  string `json:"phone,omitempty"`
	Fax    string `json:"fax,omitempty"`

	// LIS identifies the politician's ID in the congressional
	// Legislative Information System.
	LIS string `json:"lis_id,omitempty"`

	MissedVotesPct float32 `json:"missed_votes_pct,omitempty"`
	VotesWithParty float32 `json:"votes_with_party_pct,omitempty"`

	Twitter   string `json:"twitter_account,omitempty"`
	Facebook  string `json:"facebook_account,omitempty"`
	Youtube   string `json:"youtube_account,omitempty"`
	Govtrack  string `json:"govtrack_id,omitempty"`
	CSPAN     string `json:"cspan_id,omitempty"`
	Votesmart string `json:"votesmart_id,omitempty"`

	// ICPSR is "Inter-university Consortium for Political and Social Research"; this is the
	// politician's ID in their databases.
	ICPSR string `json:"icpsr_id,omitempty"`

	// CRP is the politician's ID from theCenter for Responsive Politics,
	// which runs OpenSecrets.org.
	CRP string `json:"crp_id,omitempty"`

	Google string `json:"google_entity_id,omitempty"`
	FEC    string `json:"fec_candidate_id,omitempty"`

	// DWNominate stores a politician's ideological score based on the
	// DW-NOMINATE (Dynamic Weighted NOMINAl Three-step Estimation) estimation.
	DWNominate float32 `json:"dw_nominate"`

	// *tk no idea what this is
	IdealPoint string `json:"ideal_point"`

	// URL links to the candidate's official website.
	URL string `json:"url,omitempty"`
	RSS string `json:"rss_url,omitempty"`

	// Contact links to a candidate's online contact form.
	Contact string `json:"contact_form,omitempty"`

	// API links to the endpoint for information about only this member.
	API string `json:"api_uri"`
}

// getMembersResponse is the format of the response received from the
// Congress API "get members" endpoint.
type getMembersResponse struct {
	Status    string              `json:"status"`
	Copyright string              `json:"copyright"`
	Results   []getMembersResults `json:"results"`
}

type getMembersResults struct {
	Congress    string   `json:"congress"`
	Chamber     string   `json:"chamber"`
	ResultCount int      `json:"num_results"`
	Offset      int      `json:"offset"`
	Members     []Member `json:"members"`
}

// GetMembers fetches a list of members of a defined chamber of Congress ("house" or "senate")
// in a particular session (i.e. 115).
func (c *Client) GetMembers(congress int, chamber string) (members []Member, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%d/%s/members.json", c.Endpoint, congress, chamber), nil)
	if err != nil {
		return
	}
	req.Header.Add("X-API-Key", c.Key)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	var unmarshaled getMembersResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &unmarshaled)
	if err != nil {
		return
	}
	if len(unmarshaled.Results) > 0 {
		members = unmarshaled.Results[0].Members
	}
	return
}
