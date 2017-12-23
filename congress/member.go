package congress

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Member holds the data of a member of Congress that is sent
// no matter which method was used to request their data.
type Member struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name"`
	Suffix     string `json:"suffix,omitempty"`
	Twitter    string `json:"twitter_account,omitempty"`
	Facebook   string `json:"facebook_account,omitempty"`
	Youtube    string `json:"youtube_account,omitempty"`

	// URL links to the endpoint for information about only this member.
	URL string `json:"api_uri"`
	TrackingIDs
}

// TrackingIDs collect the ID numbers of a politician across various
// websites tracking information about their actions. It is included in
// most (but not all) results.
type TrackingIDs struct {
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

	// Website links to the candidate's official website.
	Website string `json:"url,omitempty"`
	RSS     string `json:"rss_url,omitempty"`
}

// MemberSummary holds the data of a member of Congress when sent
// as part of a collection of multiple members.
type MemberSummary struct {
	Member
	ID          string `json:"id"`
	Birth       string `json:"date_of_birth"`
	State       string `json:"state,omitempty"`
	Title       string `json:"title"`
	ShortTitle  string `json:"short_title"`
	InOffice    bool   `json:"in_office"`
	SenateClass string `json:"senate_class,omitempty"` // TODO: Is this a different field for House members?
	StateRank   string `json:"state_rank,omitempty"`
	Party       string `json:"party"`
	Leadership  string `json:"leadership_role,omitempty"`
	Office      string `json:"office,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Fax         string `json:"fax,omitempty"`

	TrackingIDs

	// Contact links to a candidate's online contact form.
	Contact string `json:"contact_form,omitempty"`

	TotalVotes     int     `json:"total_votes,omitempty"`
	MissedVotes    int     `json:"missed_votes,omitempty"`
	PresentVotes   int     `json:"present_votes,omitempty"`
	NextElection   string  `json:"next_election,omitempty"`
	Seniority      string  `json:"seniority,omitempty"`
	MissedVotesPct float32 `json:"missed_votes_pct,omitempty"`
	VotesWithParty float32 `json:"votes_with_party_pct,omitempty"`

	// LIS identifies the politician's ID in the congressional
	// Legislative Information System.
	LIS string `json:"lis_id,omitempty"`

	FEC string `json:"fec_candidate_id,omitempty"`

	// DWNominate stores a politician's ideological score based on the
	// DW-NOMINATE (Dynamic Weighted NOMINAl Three-step Estimation) estimation.
	DWNominate float32 `json:"dw_nominate"`

	// *tk no idea what these are
	IdealPoint string `json:"ideal_point"`
	OCD        string `json:"ocd_id,omitempty"`
}

// MemberDetails holds the data of a member of Congress when it is requested
// specifically about that member.
type MemberDetails struct {
	Member
	ID             string       `json:"member_id"` // NOTE: JSON key is different from MemberSummary
	Birth          string       `json:"date_of_birth"`
	Gender         string       `json:"gender"`
	Party          string       `json:"current_party"` // NOTE: JSON key is different from MemberSummary
	State          string       `json:"state,omitempty"`
	InOffice       bool         `json:"in_office"`
	TimesTopics    string       `json:"times_topics_url,omitempty"`
	TimesTag       string       `json:"times_tag,omitempty"`
	MostRecentVote string       `json:"most_recent_vote,omitempty"`
	Roles          []MemberRole `json:"roles,omitempty"`
	TrackingIDs
}

// MemberSearch holds the data of a member of Congress when it is requested
// in a search specific to a state.
type MemberSearch struct {
	Member
	ID           string `json:"id"`
	Name         string `json:"name"`
	Title        string `json:"role"` // NOTE: JSON key is different than for MemberSummary and MemberRole
	Gender       string `json:"gender"`
	Party        string `json:"party"`
	TimesTopics  string `json:"times_topics_url,omitempty"`
	Seniority    string `json:"seniority,omitempty"`
	NextElection string `json:"next_election,omitempty"`
	URL          string `json:"api_uri"`
}

// MemberInTransition is the format of data sent from ProPublica
// about a single member when searching for members who are either
// new or leaving office. It doesn't really look like the others.
type MemberInTransition struct {
	ID         string `json:"id"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name,omitempty"`
	LastName   string `json:"last_name"`
	Suffix     string `json:"suffix,omitempty"`
	Party      string `json:"party"`
	Chamber    string `json:"chamber"`
	State      string `json:"state"`
	District   string `json:"district,omitempty"`
	StartDate  string `json:"start_date"`
	URL        string `json:"api_uri"`

	// Fields only included for departing members:
	EndDate string `json:"end_date"`
	Status  string `json:"status"`
	Note    string `json:"note"`
}

// MemberRole stores information about a candidate's positions in a single
// chamber of a single Congress.
type MemberRole struct {
	Congress   string `json:"congress"`
	Chamber    string `json:"chamber"`
	Title      string `json:"title"`
	ShortTitle string `json:"short_title"`
	State      string `json:"state"`
	Party      string `json:"party"`
	Leadership string `json:"leadership_role,omitempty"`
	FEC        string `json:"fed_candidate_id,omitempty"`
	Seniority  string `json:"seniority"`
	District   string `json:"district,omitempty"`
	AtLarge    bool   `json:"at_large"`
	OCD        string `json:"ocd_id,omitempty"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
	Office     string `json:"office"`
	Phone      string `json:"phone"`
	Fax        string `json:"fax"`
	Contact    string `json:"contact_form"`

	// Sponsored is the number of bills sponsored in a congress.
	Sponsored   int `json:"bills_sponsored"`
	Cosponsored int `json:"bills_cosponsored"`

	MissedVotesPct float32           `json:"missed_votes_pct"`
	VotesWithParty float32           `json:"votes_with_party_pct"`
	Committees     []MemberCommittee `json:"committees"`
}

// MemberCommittee is information about a politician's role on a committee
// in a single congress.
type MemberCommittee struct {
	Name string `json:"name"`

	// Code is the abbreviation of the committee.
	Code string `json:"code"`

	// URL is the link to more information about the committee within the API.
	URL string `json:"api_url"`

	// Side indicates whether the politician was in the majority or minority
	// on the committee in a single congress.
	Side string `json:"side"`

	Title     string `json:"member"`
	PartyRank int    `json:"rank_in_party"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
}

// MemberSubcommittee is information about a politician's role on a subcommittee
// in a single congress.
type MemberSubcommittee struct {
	MemberCommittee
	// Parent indicates the ID of the committee under which this
	// subcommittee operates.
	Parent string `json:"parent_committee_id"`
}

// getMembersResponse is the format of the response received from the
// Congress API "get members" endpoint.
type getMembersResponse struct {
	Status    string              `json:"status"`
	Copyright string              `json:"copyright"`
	Results   []getMembersResults `json:"results"`
}

type getMembersResults struct {
	Congress    string          `json:"congress"`
	Chamber     string          `json:"chamber"`
	ResultCount int             `json:"num_results"`
	Offset      int             `json:"offset"`
	Members     []MemberSummary `json:"members"`
}

// GetMembers fetches a list of members of a defined chamber of Congress ("house" or "senate")
// in a particular congress (i.e. 115).
func (c *Client) GetMembers(congress int, chamber string) (members []MemberSummary, err error) {
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

// getMemberResponse is the format of the response received from the
// Congress API "get member" endpoint.
type getMemberResponse struct {
	Status    string          `json:"status"`
	Copyright string          `json:"copyright"`
	Results   []MemberDetails `json:"results"`
}

// GetMember fetches detailed information about a single politician spanning
// their congressional career
func (c *Client) GetMember(id string) (member MemberDetails, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/members/%s.json", c.Endpoint, id), nil)
	if err != nil {
		return
	}
	req.Header.Add("X-API-Key", c.Key)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	var unmarshaled getMemberResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &unmarshaled)
	if err != nil {
		return
	}
	if len(unmarshaled.Results) > 0 {
		member = unmarshaled.Results[0]
	}
	return
}

// getMembersByStateResponse is the format of the response received from the
// Congress API "get current members by state/district" endpoint.
type getMembersByStateResponse struct {
	Status    string         `json:"status"`
	Copyright string         `json:"copyright"`
	Results   []MemberSearch `json:"results"`
}

// GetChamberMembersByState fetches basic information about the congressional delegation
// of a single chamber for a single state. ("state" param is case-insensitive two-character abbreviation.)
func (c *Client) GetChamberMembersByState(state, chamber string) (members []MemberSearch, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/members/%s/%s/current.json", c.Endpoint, chamber, state), nil)
	if err != nil {
		return
	}
	req.Header.Add("X-API-Key", c.Key)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	var unmarshaled getMembersByStateResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &unmarshaled)
	if err != nil {
		return
	}
	members = unmarshaled.Results

	return
}

// GetMembersByState fetches basic information about the entire congressional
// delegation for a single state
func (c *Client) GetMembersByState(state string) (members []MemberSearch, err error) {
	members, err = c.GetChamberMembersByState(state, "house")
	if err != nil {
		return
	}
	senate, err := c.GetChamberMembersByState(state, "senate")
	if err != nil {
		return
	}
	members = append(members, senate...)
	return
}

// GetChamberMembersByDistrict fetches basic information about the congressional delegation
// of a single chamber for a single district of a state.
func (c *Client) GetChamberMembersByDistrict(state string, district int, chamber string) (members []MemberSearch, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/members/%s/%s/%d/current.json", c.Endpoint, chamber, state, district), nil)
	if err != nil {
		return
	}
	req.Header.Add("X-API-Key", c.Key)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	var unmarshaled getMembersByStateResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &unmarshaled)
	if err != nil {
		return
	}
	members = unmarshaled.Results

	return
}

// getMembersInTransitionResponse is the format of the response received from the
// Congress API "get new members" and "get members leaving office" endpoints.
type getMembersInTransitionResponse struct {
	Status    string                          `json:"status"`
	Copyright string                          `json:"copyright"`
	Results   []getMembersInTransitionResults `json:"results"`
}

type getMembersInTransitionResults struct {
	Members []MemberInTransition `json:"members"`

	// NOTE: These two fields are excluded because there's no known pagination in
	// the API yet, and because they appear inconsistently in results: When requesting
	// new members, they're strings. When requesting departing members, they are ints.
	// ResultCount int `json:"num_results"`
	// Offset      int `json:"offset"`

	// Fields only included in response for departing members:
	Congress string `json:"congress"`
	Chamber  string `json:"chamber"`
}

// GetNewMembers fetches basic information about the first-time members
// of either chamber.
func (c *Client) GetNewMembers() (members []MemberInTransition, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/members/new.json", c.Endpoint), nil)
	if err != nil {
		return
	}
	req.Header.Add("X-API-Key", c.Key)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	var unmarshaled getMembersInTransitionResponse
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

// GetDepartingMembers fetches basic information about the outgoing members
// of both chambers for a particular Congress.
func (c *Client) GetDepartingMembers(congress int, chamber string) (members []MemberInTransition, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%d/%s/members/leaving.json", c.Endpoint, congress, chamber), nil)
	if err != nil {
		return
	}
	req.Header.Add("X-API-Key", c.Key)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	var unmarshaled getMembersInTransitionResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &unmarshaled)
	if err != nil {
		return
	}
	// Results are broken down by chamber
	if len(unmarshaled.Results) > 0 {
		members = unmarshaled.Results[0].Members
	}
	if len(unmarshaled.Results) > 1 {
		members = append(members, unmarshaled.Results[1].Members...)
	}
	return
}
