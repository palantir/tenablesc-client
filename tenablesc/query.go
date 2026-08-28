package tenablesc

import (
	"fmt"
)

const queryEndpoint = "/query"

// Query represents the request/response structure for https://docs.tenable.com/tenablesc/api/Query.htm
type Query struct {
	BaseInfo
	Tool         string              `json:"tool,omitempty"`
	Type         string              `json:"type,omitempty"`
	Tags         string              `json:"tags,omitempty"`
	Context      string              `json:"context,omitempty"`
	CreatedTime  UnixEpochStringTime `json:"createdTime,omitempty"`
	ModifiedTime UnixEpochStringTime `json:"modifiedTime,omitempty"`
	Status       string              `json:"status,omitempty"`
	Filters      []Filter            `json:"filters,omitempty"`
	CanManage    string              `json:"canManage,omitempty"`
	CanUse       string              `json:"canUse,omitempty"`
	Creator      UserInfo            `json:"creator,omitempty"`
	Owner        UserInfo            `json:"owner,omitempty"`
	OwnerGroup   BaseInfo            `json:"ownerGroup,omitempty"`
	TargetGroup  BaseInfo            `json:"targetGroup,omitempty"`
}

type Filter struct {
	FilterName string `json:"filterName,omitempty"`
	Operator   string `json:"operator,omitempty"`
	// "value" format depends on filter's "filterName" parameter
	Value interface{} `json:"value,omitempty"`
}

func (c *Client) GetQuery(id string) (*Query, error) {
	resp := &Query{}

	if _, err := c.getResource(fmt.Sprintf("%s/%s", queryEndpoint, id), resp); err != nil {
		return nil, fmt.Errorf("failed to get query id %s: %w", id, err)
	}

	return resp, nil
}

type queryResponse struct {
	Manageable []*Query `json:"manageable" tenable:"recurse"`
	Usable     []*Query `json:"usable" tenable:"recurse"`
}

func (o queryResponse) allQueriesToExternal() []*Query {
	var spOut []*Query

	spMap := make(map[ProbablyString]bool)

	for _, o := range o.Usable {
		spOut = append(spOut, o)
		spMap[o.ID] = true
	}

	for _, o := range o.Manageable {
		if _, exists := spMap[o.ID]; !exists {
			spOut = append(spOut, o)
			spMap[o.ID] = true
		}
	}

	return spOut
}

func (c *Client) GetAllQueries() ([]*Query, error) {
	var queries queryResponse

	_, err := c.getResource(queryEndpoint, &queries)
	if err == nil {
		return queries.allQueriesToExternal(), nil
	}

	var q []*Query
	if _, err := c.getResource(queryEndpoint, &q); err != nil {
		return nil, fmt.Errorf("failed to get queries: %w", err)
	}

	return q, nil
}
