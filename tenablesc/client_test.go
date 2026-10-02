package tenablesc

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
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
	} `json:"nested" tenable:"recurse"`
}

func TestGetFields(t *testing.T) {
	fields := []string{"bare", "first", "otherBare"}

	assert.Equal(t, fields, getFieldsForStruct(testGetFieldsStruct{}))
	assert.Equal(t, fields, getFieldsForStruct([]testGetFieldsStruct{}))

}

func TestRedirectPolicy(t *testing.T) {
	t.Run("disabled by default", func(t *testing.T) {
		redirected := false
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/start" {
				http.Redirect(w, r, "/finish", http.StatusFound)
				return
			}

			redirected = true
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		client := NewClient(server.URL)

		_, err := client.RestyClient().R().Get("/start")
		assert.ErrorContains(t, err, "auto redirect is disabled")
		assert.False(t, redirected)
	})

	t.Run("custom policy opts in", func(t *testing.T) {
		redirected := false
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == "/start" {
				http.Redirect(w, r, "/finish", http.StatusFound)
				return
			}

			redirected = true
			w.WriteHeader(http.StatusOK)
		}))
		defer server.Close()

		client := NewClient(server.URL)
		client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(10))

		_, err := client.RestyClient().R().Get("/start")
		assert.NoError(t, err)
		assert.True(t, redirected)
	})
}
