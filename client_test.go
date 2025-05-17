// Copyright 2025 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package amber_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/freman/amber"
)

func TestClientGetPrices(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "test_data/sites_n_prices.json")
		},
	))
	defer s.Close()

	c := amber.New("testing", amber.WithHTTPClient(s.Client()), amber.WithBaseURL(s.URL))

	r, err := c.GetPrices("test")
	assert.NoError(t, err)

	assert.Len(t, r, 2) // No controlled load in test data
	assert.IsType(t, amber.ActualInterval{}, r[amber.ChannelTypeFeedIn][0])
	assert.IsType(t, amber.ActualInterval{}, r[amber.ChannelTypeFeedIn][204])
	assert.IsType(t, amber.CurrentInterval{}, r[amber.ChannelTypeFeedIn][205])
	assert.IsType(t, amber.ForecastInterval{}, r[amber.ChannelTypeFeedIn][206])

	assert.IsType(t, amber.ActualInterval{}, r[amber.ChannelTypeGeneral][0])
	assert.IsType(t, amber.ActualInterval{}, r[amber.ChannelTypeGeneral][204])
	assert.IsType(t, amber.CurrentInterval{}, r[amber.ChannelTypeGeneral][205])
	assert.IsType(t, amber.ForecastInterval{}, r[amber.ChannelTypeGeneral][206])
}

func TestClientGetCurrentPrices(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "test_data/sites_n_prices_current.json")
		},
	))
	defer s.Close()

	c := amber.New("testing", amber.WithHTTPClient(s.Client()), amber.WithBaseURL(s.URL))

	r, err := c.GetCurrentPrices("test")
	assert.NoError(t, err)

	assert.Len(t, r, 2) // No controlled load in test data
	assert.IsType(t, amber.CurrentInterval{}, r[amber.ChannelTypeFeedIn][0])
	assert.IsType(t, amber.ForecastInterval{}, r[amber.ChannelTypeFeedIn][1])

	assert.IsType(t, amber.CurrentInterval{}, r[amber.ChannelTypeGeneral][0])
	assert.IsType(t, amber.ForecastInterval{}, r[amber.ChannelTypeGeneral][1])
}

func TestClientGetUsage(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "test_data/sites_n_usage.json")
		},
	))
	defer s.Close()

	c := amber.New("testing", amber.WithHTTPClient(s.Client()), amber.WithBaseURL(s.URL))

	r, err := c.GetUsage("test")
	assert.NoError(t, err)

	assert.Len(t, r, 2) // No controlled load in test data
	assert.IsType(t, amber.Usage{}, r[amber.ChannelTypeFeedIn][0])
	assert.IsType(t, amber.Usage{}, r[amber.ChannelTypeGeneral][0])
}

func TestClientSites(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "test_data/sites.json")
		},
	))
	defer s.Close()

	c := amber.New("testing", amber.WithHTTPClient(s.Client()), amber.WithBaseURL(s.URL))

	r, err := c.GetSites()
	assert.NoError(t, err)
	assert.NotEmpty(t, r)
	assert.NotEmpty(t, r[0].ID)
}

func TestClientRenewables(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, []string{"10"}, r.URL.Query()["next"])
			assert.Equal(t, []string{"10"}, r.URL.Query()["previous"])
			http.ServeFile(w, r, "test_data/renewables.json")
		},
	))
	defer s.Close()

	c := amber.New("testing", amber.WithHTTPClient(s.Client()), amber.WithBaseURL(s.URL))

	r, err := c.GetCurrentRenewables("qld", amber.WithPrevious(10), amber.WithNext(10))
	assert.NoError(t, err)

	assert.Len(t, r, 21) // No controlled load in test data
}
