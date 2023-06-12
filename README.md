# Unofficial GO api for [Amber](https://app.amber.com.au/developers/documentation/)

Early days, api is subject to change.

## Variables

```golang
var ErrInvalidState = errors.New("invalid state provided, valid states are nsw, sa, qld, or vic")
```

## Types

### type [Client](/client.go#L15)

`type Client struct { ... }`

#### func [New](/client.go#L23)

`func New(token string, opts ...Option) *Client`

#### func (*Client) [GetCurrentPrices](/client.go#L128)

`func (c *Client) GetCurrentPrices(siteID string, args ...QueryArgument) ([]schema.Interval, error)`

GetCurrentPrices returns the current price

SiteID is the ID of the site you are fetching prices for.

Query Arquments:

```go
WithNext - Return the next number of forecast intervals
WithPrevious - Return the previous number of actual intervals.
WithResolution - Specify the required interval duration resolution. Valid options: 30. Default: 30
```

#### func (*Client) [GetCurrentRenewables](/client.go#L50)

`func (c *Client) GetCurrentRenewables(state string, args ...QueryArgument) ([]schema.Renewable, error)`

GetCurrentRenewables returns the current percentage of renewables in the grid

State you would like the renewables for. Valid states: nsw, sa, qld, vic.

Query Arquments:

```go
WithNext - Return the next number of forecast intervals
WithPrevious - Return the previous number of actual intervals.
WithResolution - Specify the required interval duration resolution. Valid options: 30. Default: 30
```

#### func (*Client) [GetPrices](/client.go#L94)

`func (c *Client) GetPrices(siteID string, args ...QueryArgument) ([]schema.Interval, error)`

GetPrices returns all the prices between the start and end dates

SiteID is the ID of the site you are fetching prices for.

Query Arquments:

```go
WithStartDate - Return all prices for each interval on and after this day. Defaults to today.
WithEndDate - Return all prices for each interval on and before this day. Defaults to today.
WithResolution - Specify the required interval duration resolution. Valid options: 5, 30. Default: 30
```

#### func (*Client) [GetSites](/client.go#L76)

`func (c *Client) GetSites() ([]schema.Site, error)`

GetSites returns all sites linked to your account.

#### func (*Client) [GetUsage](/client.go#L159)

`func (c *Client) GetUsage(siteID string, args ...QueryArgument) ([]schema.Usage, error)`

GetUsage returns all usage data between the start and end dates. To avoid request timing out,
please keep date ranges to a month or less.

SiteID is the ID  of the site you are fetching usage for.

Query Arguments:

```go
WithStartDate - Return all usage for each interval on and after this day.
WithEndDate - Return all usage for each interval on and before this day.
WithResolution - Specify the required interval duration resolution. Valid options: 30. Default: 30
```

### type [HTTPError](/httperr.go#L5)

`type HTTPError struct { ... }`

#### func (HTTPError) [Error](/httperr.go#L12)

`func (err HTTPError) Error() string`

### type [Option](/options.go#L5)

`type Option func(c *Client)`

#### func [WithHTTPClient](/options.go#L7)

`func WithHTTPClient(client *http.Client) Option`

### type [QueryArgument](/queryArgs.go#L9)

`type QueryArgument func(v url.Values)`

#### func [WithEndDate](/queryArgs.go#L35)

`func WithEndDate(date time.Time) QueryArgument`

#### func [WithNext](/queryArgs.go#L11)

`func WithNext(intervals uint) QueryArgument`

#### func [WithPrevious](/queryArgs.go#L17)

`func WithPrevious(intervals uint) QueryArgument`

#### func [WithResolution](/queryArgs.go#L23)

`func WithResolution(resolution uint) QueryArgument`

#### func [WithStartDate](/queryArgs.go#L29)

`func WithStartDate(date time.Time) QueryArgument`

## Sub Packages

* [schema](./schema)

