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
	Value      string `json:"value,omitempty"`
}

type QueryResponse struct {
	Manageable []*Query `json:"manageable" tenable:"recurse"`
	Usable     []*Query `json:"usable" tenable:"recurse"`
}

func (o QueryResponse) allQueriesToExternal() []*Query {
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

func (c *Client) GetQuery(id string) (*Query, error) {
	query := &Query{}

	if _, err := c.getResource(fmt.Sprintf("%s/%s", queryEndpoint, id), query); err != nil {
		return nil, fmt.Errorf("failed to get query id %s: %w", id, err)
	}

	return query, nil
}

func (c *Client) GetAllQueries() ([]*Query, error) {
	var queries QueryResponse
	_, err := c.getResource(queryEndpoint, &queries)
	if err == nil {
		return queries.allQueriesToExternal(), nil
	}

	return queries.allQueriesToExternal(), nil

}
