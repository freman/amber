package amber_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/freman/amber"
	"github.com/freman/amber/schema"
)

func TestClientGetPrices(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "test_data/sites_n_prices.json")
		},
	))
	defer s.Close()

	c := amber.New("tesing", amber.WithHTTPClient(s.Client()), amber.WithBaseURL(s.URL))

	r, err := c.GetPrices("test")
	assert.NoError(t, err)

	assert.Len(t, r, 2) // No controlled load in test data
	assert.IsType(t, schema.ActualInterval{}, r[schema.ChannelTypeFeedIn][0])
	assert.IsType(t, schema.ActualInterval{}, r[schema.ChannelTypeFeedIn][204])
	assert.IsType(t, schema.CurrentInterval{}, r[schema.ChannelTypeFeedIn][205])
	assert.IsType(t, schema.ForecastInterval{}, r[schema.ChannelTypeFeedIn][206])

	assert.IsType(t, schema.ActualInterval{}, r[schema.ChannelTypeGeneral][0])
	assert.IsType(t, schema.ActualInterval{}, r[schema.ChannelTypeGeneral][204])
	assert.IsType(t, schema.CurrentInterval{}, r[schema.ChannelTypeGeneral][205])
	assert.IsType(t, schema.ForecastInterval{}, r[schema.ChannelTypeGeneral][206])

	spew.Dump(r[schema.ChannelTypeGeneral][206])
	spew.Dump(r[schema.ChannelTypeFeedIn][206])
}

func TestClientGetCurrentPrices(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "test_data/sites_n_prices_current.json")
		},
	))
	defer s.Close()

	c := amber.New("tesing", amber.WithHTTPClient(s.Client()), amber.WithBaseURL(s.URL))

	r, err := c.GetCurrentPrices("test")
	assert.NoError(t, err)

	assert.Len(t, r, 2) // No controlled load in test data
	assert.IsType(t, schema.CurrentInterval{}, r[schema.ChannelTypeFeedIn][0])
	assert.IsType(t, schema.ForecastInterval{}, r[schema.ChannelTypeFeedIn][1])

	assert.IsType(t, schema.CurrentInterval{}, r[schema.ChannelTypeGeneral][0])
	assert.IsType(t, schema.ForecastInterval{}, r[schema.ChannelTypeGeneral][1])

}
