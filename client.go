// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package amber

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/freman/amber/schema"
)

type Client struct {
	token string

	httpClient *http.Client
}

var ErrInvalidState = errors.New("invalid state provided, valid states are nsw, sa, qld, or vic")

func New(token string, opts ...Option) *Client {
	c := Client{
		token: token,
	}

	for _, opt := range opts {
		opt(&c)
	}

	if c.httpClient == nil {
		c.httpClient = &http.Client{
			Timeout: time.Minute,
		}
	}

	return &c
}

// GetCurrentRenewables returns the current percentage of renewables in the grid
//
// State you would like the renewables for. Valid states: nsw, sa, qld, vic.
//
// Query Arquments:
//
//	WithNext - Return the next number of forecast intervals
//	WithPrevious - Return the previous number of actual intervals.
//	WithResolution - Specify the required interval duration resolution. Valid options: 30. Default: 30
func (c *Client) GetCurrentRenewables(state string, args ...QueryArgument) ([]schema.Renewable, error) {
	if !(state == "nsw" || state == "qld" || state == "vic" || state == "sa") {
		return nil, ErrInvalidState
	}

	uri := fmt.Sprintf("https://api.amber.com.au/v1/state/%s/renewables/current", state)

	query := url.Values{
		"resolution": []string{"30"},
	}

	for _, arg := range args {
		arg(query)
	}

	uri += "?" + query.Encode()

	var res []schema.Renewable
	if err := c.get(uri, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetSites returns all sites linked to your account.
func (c *Client) GetSites() ([]schema.Site, error) {
	var res []schema.Site
	if err := c.get("https://api.amber.com.au/v1/sites", &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetPrices returns all the prices between the start and end dates
//
// SiteID is the ID of the site you are fetching prices for.
//
// Query Arquments:
//
//	WithStartDate - Return all prices for each interval on and after this day. Defaults to today.
//	WithEndDate - Return all prices for each interval on and before this day. Defaults to today.
//	WithResolution - Specify the required interval duration resolution. Valid options: 5, 30. Default: 30
func (c *Client) GetPrices(siteID string, args ...QueryArgument) ([]schema.Interval, error) {
	uri := fmt.Sprintf("https://api.amber.com.au/v1/sites/%s/prices", siteID)

	now := time.Now().Format(time.DateOnly)

	query := url.Values{
		"startDate":  []string{now},
		"endDate":    []string{now},
		"resolution": []string{"30"},
	}

	for _, arg := range args {
		arg(query)
	}

	uri += "?" + query.Encode()

	var res []schema.Interval
	if err := c.get(uri, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetCurrentPrices returns the current price
//
// SiteID is the ID of the site you are fetching prices for.
//
// Query Arquments:
//
//	WithNext - Return the next number of forecast intervals
//	WithPrevious - Return the previous number of actual intervals.
//	WithResolution - Specify the required interval duration resolution. Valid options: 30. Default: 30
func (c *Client) GetCurrentPrices(siteID string, args ...QueryArgument) ([]schema.Interval, error) {
	uri := fmt.Sprintf("https://api.amber.com.au/v1/sites/%s/prices/current", siteID)

	query := url.Values{
		"resolution": []string{"30"},
	}

	for _, arg := range args {
		arg(query)
	}

	uri += "?" + query.Encode()

	var res []schema.Interval
	if err := c.get(uri, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetUsage returns all usage data between the start and end dates. To avoid request timing out,
// please keep date ranges to a month or less.
//
// SiteID is the ID  of the site you are fetching usage for.
//
// Query Arguments:
//
//	WithStartDate - Return all usage for each interval on and after this day.
//	WithEndDate - Return all usage for each interval on and before this day.
//	WithResolution - Specify the required interval duration resolution. Valid options: 30. Default: 30
func (c *Client) GetUsage(siteID string, args ...QueryArgument) ([]schema.Usage, error) {
	uri := fmt.Sprintf("https://api.amber.com.au/v1/sites/%s/usage", siteID)

	query := url.Values{
		"resolution": []string{"30"},
	}

	for _, arg := range args {
		arg(query)
	}

	uri += "?" + query.Encode()

	var res []schema.Usage
	if err := c.get(uri, &res); err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) get(url string, out interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.token)
	req.Header.Set("User-Agent", "github.com_freman_amber/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)

		return &HTTPError{
			resp.StatusCode,
			resp.Status,
			body,
			err,
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return fmt.Errorf("failed to decode JSON body: %w", err)
	}

	return nil
}
