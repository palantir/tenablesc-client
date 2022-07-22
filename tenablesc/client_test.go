package tenablesc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// tests for the utility getFieldsForStruct logic in the client.

type firstFieldStruct struct {
	First string `json:"first,omitempty"`
}

type testGetFieldsStruct struct {
	firstFieldStruct
	Bare   string `json:"bare,omitempty"`
	Nested struct {
		OtherBare string `json:"otherBare,omitempty"`
	} `json:"nested,omitempty" tenable:"recurse"`
}

func TestGetFields(t *testing.T) {
	fields := []string{"bare", "first", "otherBare"}

	assert.Equal(t, fields, getFieldsForStruct(testGetFieldsStruct{}))
	assert.Equal(t, fields, getFieldsForStruct([]testGetFieldsStruct{}))

}
