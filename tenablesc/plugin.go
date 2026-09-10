package tenablesc

import (
	"fmt"
	"net/url"
)

const pluginEndpoint = "/plugin"

// Plugin represents the response structure for https://docs.tenable.com/tenablesc/api/Plugin.htm
type Plugin struct {
	BaseInfo
	Family              Family         `json:"family,omitempty"`
	Synopsis            string         `json:"synopsis,omitempty"`
	Solution            string         `json:"solution,omitempty"`
	SeeAlso             string         `json:"seeAlso,omitempty"`
	CPE                 string         `json:"cpe,omitempty"`
	RiskFactor          string         `json:"riskFactor,omitempty"`
	CVSSVector          string         `json:"cvssVector,omitempty"`
	BaseScore           string         `json:"baseScore,omitempty"`
	TemporalScore       string         `json:"temporalScore,omitempty"`
	CVSSV3BaseScore     string         `json:"cvssV3BaseScore,omitempty"`
	CVSSV3TemporalScore string         `json:"cvssV3TemporalScore,omitempty"`
	CVSSV3Vector        string         `json:"cvssV3Vector,omitempty"`
	VPRScore            string         `json:"vprScore,omitempty"`
	VPRContext          string         `json:"vprContext,omitempty"`
	EPSSScore           string         `json:"epssScore,omitempty"`
	StigSeverity        string         `json:"stigSeverity,omitempty"`
	ExploitAvailable    string         `json:"exploitAvailable,omitempty"`
	ExploitEase         string         `json:"exploitEase,omitempty"`
	ExploitFrameworks   string         `json:"exploitFrameworks,omitempty"`
	XRefs               string         `json:"xrefs,omitempty"`
	PluginPubDate       ProbablyString `json:"pluginPubDate,omitempty"`
	PluginModDate       ProbablyString `json:"pluginModDate,omitempty"`
	PatchPubDate        ProbablyString `json:"patchPubDate,omitempty"`
	PatchModDate        ProbablyString `json:"patchModDate,omitempty"`
	VulnPubDate         ProbablyString `json:"vulnPubDate,omitempty"`
	ModifiedTime        ProbablyString `json:"modifiedTime,omitempty"`
	CheckType           string         `json:"checkType,omitempty"`
}

type Family struct {
	BaseInfo
	Type string `json:"type,omitempty"`
}

func (c *Client) GetPlugin(id string) (*Plugin, error) {
	plugin := &Plugin{}

	if _, err := c.getResource(fmt.Sprintf("%s/%s", pluginEndpoint, id), plugin); err != nil {
		return nil, fmt.Errorf("failed to get plugin id %s: %w", id, err)
	}

	return plugin, nil
}

func (c *Client) GetPluginsByName(name string) ([]*Plugin, error) {
	var plugins []*Plugin

	query := url.Values{}
	query.Add("sortDirection", "ASC")
	query.Add("sortField", "name")
	query.Add("filterField", "name")
	query.Add("op", "eq")
	query.Add("value", name)

	if _, err := c.getResource(fmt.Sprintf("%s?%s", pluginEndpoint, query.Encode()), &plugins); err != nil {
		return nil, fmt.Errorf("failed to find plugin with name %s: %w", name, err)
	}

	return plugins, nil
}
