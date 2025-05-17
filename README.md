# Unofficial GO api for [Amber](https://app.amber.com.au/developers/documentation/)

Early days, api is subject to change.

# amber

```go
import "github.com/freman/amber"
```

## Index

- [Variables](<#variables>)
- [type ActualInterval](<#ActualInterval>)
  - [func \(b ActualInterval\) GetType\(\) string](<#ActualInterval.GetType>)
- [type AdvancedPrice](<#AdvancedPrice>)
- [type BaseInterval](<#BaseInterval>)
  - [func \(b BaseInterval\) GetChannelType\(\) ChannelType](<#BaseInterval.GetChannelType>)
- [type Channel](<#Channel>)
- [type ChannelType](<#ChannelType>)
- [type Client](<#Client>)
  - [func New\(token string, opts ...Option\) \*Client](<#New>)
  - [func \(c \*Client\) GetCurrentPrices\(siteID string, args ...QueryArgument\) \(IntervalMap\[Interval\], error\)](<#Client.GetCurrentPrices>)
  - [func \(c \*Client\) GetCurrentRenewables\(state string, args ...QueryArgument\) \(\[\]Renewable, error\)](<#Client.GetCurrentRenewables>)
  - [func \(c \*Client\) GetPrices\(siteID string, args ...QueryArgument\) \(IntervalMap\[Interval\], error\)](<#Client.GetPrices>)
  - [func \(c \*Client\) GetSites\(\) \(\[\]Site, error\)](<#Client.GetSites>)
  - [func \(c \*Client\) GetUsage\(siteID string, args ...QueryArgument\) \(map\[ChannelType\]\[\]Usage, error\)](<#Client.GetUsage>)
- [type CurrentInterval](<#CurrentInterval>)
  - [func \(b CurrentInterval\) GetType\(\) string](<#CurrentInterval.GetType>)
- [type Date](<#Date>)
  - [func \(d \*Date\) MarshalJSON\(\) \(\[\]byte, error\)](<#Date.MarshalJSON>)
  - [func \(d Date\) String\(\) string](<#Date.String>)
  - [func \(d \*Date\) UnmarshalJSON\(in \[\]byte\) \(err error\)](<#Date.UnmarshalJSON>)
- [type Descriptor](<#Descriptor>)
- [type Duration](<#Duration>)
  - [func \(id \*Duration\) MarshalJSON\(\) \(\[\]byte, error\)](<#Duration.MarshalJSON>)
  - [func \(id \*Duration\) UnmarshalJSON\(in \[\]byte\) error](<#Duration.UnmarshalJSON>)
- [type ForecastInterval](<#ForecastInterval>)
  - [func \(b ForecastInterval\) GetType\(\) string](<#ForecastInterval.GetType>)
- [type HTTPError](<#HTTPError>)
  - [func \(err HTTPError\) Error\(\) string](<#HTTPError.Error>)
- [type Interval](<#Interval>)
- [type IntervalLength](<#IntervalLength>)
  - [func \(i IntervalLength\) Duration\(\) time.Duration](<#IntervalLength.Duration>)
  - [func \(i IntervalLength\) Valid\(\) bool](<#IntervalLength.Valid>)
- [type IntervalMap](<#IntervalMap>)
  - [func \(im \*IntervalMap\[T\]\) UnmarshalJSON\(data \[\]byte\) error](<#IntervalMap[T].UnmarshalJSON>)
- [type IntervalSlice](<#IntervalSlice>)
- [type Option](<#Option>)
  - [func WithBaseURL\(uri string\) Option](<#WithBaseURL>)
  - [func WithHTTPClient\(client \*http.Client\) Option](<#WithHTTPClient>)
- [type Period](<#Period>)
- [type Quality](<#Quality>)
- [type QueryArgument](<#QueryArgument>)
  - [func WithEndDate\(date time.Time\) QueryArgument](<#WithEndDate>)
  - [func WithNext\(intervals uint\) QueryArgument](<#WithNext>)
  - [func WithPrevious\(intervals uint\) QueryArgument](<#WithPrevious>)
  - [func WithResolution\(resolution uint\) QueryArgument](<#WithResolution>)
  - [func WithStartDate\(date time.Time\) QueryArgument](<#WithStartDate>)
- [type Range](<#Range>)
- [type Renewable](<#Renewable>)
- [type RenewableDescriptor](<#RenewableDescriptor>)
- [type Season](<#Season>)
- [type Site](<#Site>)
- [type SiteStatus](<#SiteStatus>)
- [type SpikeStatus](<#SpikeStatus>)
- [type Tariff](<#Tariff>)
- [type Usage](<#Usage>)


## Variables

<a name="ErrInvalidState"></a>

```go
var ErrInvalidState = errors.New("invalid state provided, valid states are nsw, sa, qld, or vic")
```

<a name="ActualInterval"></a>
## type [ActualInterval](<https://github.com/freman/amber/blob/master/interval.go#L40-L42>)



```go
type ActualInterval struct {
    BaseInterval
}
```

<a name="ActualInterval.GetType"></a>
### func \(ActualInterval\) [GetType](<https://github.com/freman/amber/blob/master/interval.go#L44>)

```go
func (b ActualInterval) GetType() string
```



<a name="AdvancedPrice"></a>
## type [AdvancedPrice](<https://github.com/freman/amber/blob/master/interval.go#L94-L98>)

Advamcedprice is part of Ambers advanced forecast system that represents their confidence in the AEMO forecast. The range indicates where they think the price will land for a given interval.

```go
type AdvancedPrice struct {
    Low       float64 `json:"low,omitempty"`       // The lower bound of prediction band. (c/kWh).
    Predicted float64 `json:"predicted,omitempty"` // The predicted price. Use this if you need a single number to forecast against. (c/kWh).
    High      float64 `json:"high,omitempty"`      // The upper bound of prediction band. Price includes network and market fees. (c/kWh).
}
```

<a name="BaseInterval"></a>
## type [BaseInterval](<https://github.com/freman/amber/blob/master/interval.go#L20-L34>)

Interval is one time interval.

```go
type BaseInterval struct {
    Type        string      `json:"type"`
    Duration    Duration    `json:"duration"`          // Length of the interval in minutes. [5,15,30]
    SpotPerKwh  float64     `json:"spotPerKwh"`        // NEM spot price (c/kWh).
    PerKwh      float64     `json:"perKwh"`            // Number of cents you will pay per kilowatt-hour (c/kWh) - includes GST
    Date        Date        `json:"date"`              // Date the interval belongs to (in NEM time). This may be different to the date component of nemTime, as the last interval of the day ends at 12:00 the following day. [2006-01-02]
    NemTime     time.Time   `json:"nemTime"`           // The interval's NEM time. This represents the time at the end of the interval UTC+10. Formatted as a ISO 8601 time
    StartTime   time.Time   `json:"startTime"`         // Start time of the interval in UTC. Formatted as a ISO 8601 time
    EndTime     time.Time   `json:"endTime"`           // End time of the interval in UTC. Formatted as a ISO 8601 time
    Renewables  float64     `json:"renewables"`        // Percentage of renewables in the grid 0-100
    ChannelType ChannelType `json:"channelType"`       // Meter channel type [general, controlledLoad, feedIn]
    Tariff      Tariff      `json:"tariffInformation"` // Information about how your tariff is affecting this interval
    SpikeStatus SpikeStatus `json:"spikeStatus"`       // Indicates whether this interval will potentially spike, or is currently in a spike state. [none, potential, spike]
    Descriptor  Descriptor  `json:"descriptor"`        // Describes the current price. Gives you an indication of how cheap the price is in relation to the average VMO and DMO. Note: Negative is no longer used. It has been replaced with extremelyLow. [ negative, extremelyLow, veryLow, low, neutral, high, spike ]
}
```

<a name="BaseInterval.GetChannelType"></a>
### func \(BaseInterval\) [GetChannelType](<https://github.com/freman/amber/blob/master/interval.go#L36>)

```go
func (b BaseInterval) GetChannelType() ChannelType
```



<a name="Channel"></a>
## type [Channel](<https://github.com/freman/amber/blob/master/channel.go#L17-L21>)

Channel Describes a power meter channel.

The General channel provides continuous power \- it's the channel all of your appliances and lights are attached to.

Controlled loads are only on for a limited time during the day \(usually when the load on the network is low, or generation is high\) \- you may have your hot water system attached to this channel.

The feed in channel sends power back to the grid \- you will have these types of channels if you have solar or batteries.

```go
type Channel struct {
    Identifier string      `json:"identifier"` // Identifier of the channel.
    Type       ChannelType `json:"type"`       // Channel type. [general, controlledLoad, feedIn]
    Tariff     string      `json:"tariff"`     // The tariff code of the channel.
}
```

<a name="ChannelType"></a>
## type [ChannelType](<https://github.com/freman/amber/blob/master/channel.go#L23>)



```go
type ChannelType string
```

<a name="ChannelTypeGeneral"></a>

```go
const (
    ChannelTypeGeneral        ChannelType = "general"
    ChannelTypeControlledLoad ChannelType = "controlledLoad"
    ChannelTypeFeedIn         ChannelType = "feedIn"
)
```

<a name="SiteStatusPending"></a>

```go
const (
    SiteStatusPending ChannelType = "pending"
    SiteStatusActive  ChannelType = "active"
    SiteStatusClosed  ChannelType = "cloased"
)
```

<a name="Client"></a>
## type [Client](<https://github.com/freman/amber/blob/master/client.go#L18-L23>)



```go
type Client struct {
    // contains filtered or unexported fields
}
```

<a name="New"></a>
### func [New](<https://github.com/freman/amber/blob/master/client.go#L27>)

```go
func New(token string, opts ...Option) *Client
```



<a name="Client.GetCurrentPrices"></a>
### func \(\*Client\) [GetCurrentPrices](<https://github.com/freman/amber/blob/master/client.go#L132>)

```go
func (c *Client) GetCurrentPrices(siteID string, args ...QueryArgument) (IntervalMap[Interval], error)
```

GetCurrentPrices returns the current price

SiteID is the ID of the site you are fetching prices for.

Query Arquments:

```
WithNext - Return the next number of forecast intervals
WithPrevious - Return the previous number of actual intervals.
WithResolution - Specify the required interval duration resolution. Valid options: 30. Default: 30
```

<a name="Client.GetCurrentRenewables"></a>
### func \(\*Client\) [GetCurrentRenewables](<https://github.com/freman/amber/blob/master/client.go#L58>)

```go
func (c *Client) GetCurrentRenewables(state string, args ...QueryArgument) ([]Renewable, error)
```

GetCurrentRenewables returns the current percentage of renewables in the grid

State you would like the renewables for. Valid states: nsw, sa, qld, vic.

Query Arquments:

```
WithNext - Return the next number of forecast intervals
WithPrevious - Return the previous number of actual intervals.
WithResolution - Specify the required interval duration resolution. Valid options: 30. Default: 30
```

<a name="Client.GetPrices"></a>
### func \(\*Client\) [GetPrices](<https://github.com/freman/amber/blob/master/client.go#L102>)

```go
func (c *Client) GetPrices(siteID string, args ...QueryArgument) (IntervalMap[Interval], error)
```

GetPrices returns all the prices between the start and end dates

SiteID is the ID of the site you are fetching prices for.

Query Arquments:

```
WithStartDate - Return all prices for each interval on and after this day. Defaults to today.
WithEndDate - Return all prices for each interval on and before this day. Defaults to today.
WithResolution - Specify the required interval duration resolution. Valid options: 5, 30. Default: 30
```

<a name="Client.GetSites"></a>
### func \(\*Client\) [GetSites](<https://github.com/freman/amber/blob/master/client.go#L84>)

```go
func (c *Client) GetSites() ([]Site, error)
```

GetSites returns all sites linked to your account.

<a name="Client.GetUsage"></a>
### func \(\*Client\) [GetUsage](<https://github.com/freman/amber/blob/master/client.go#L161>)

```go
func (c *Client) GetUsage(siteID string, args ...QueryArgument) (map[ChannelType][]Usage, error)
```

GetUsage returns all usage data between the start and end dates. To avoid request timing out, please keep date ranges to a month or less.

SiteID is the ID of the site you are fetching usage for.

Query Arguments:

```
WithStartDate - Return all usage for each interval on and after this day.
WithEndDate - Return all usage for each interval on and before this day.
WithResolution - Deprecated, upstream will ignore it
```

<a name="CurrentInterval"></a>
## type [CurrentInterval](<https://github.com/freman/amber/blob/master/interval.go#L58-L63>)



```go
type CurrentInterval struct {
    BaseInterval
    Range         Range         `json:"range,omitempty"`
    AdvancedPrice AdvancedPrice `json:"advancedPrice,omitempty"`
    Estimate      bool          `json:"estimate,omitempty"`
}
```

<a name="CurrentInterval.GetType"></a>
### func \(CurrentInterval\) [GetType](<https://github.com/freman/amber/blob/master/interval.go#L65>)

```go
func (b CurrentInterval) GetType() string
```



<a name="Date"></a>
## type [Date](<https://github.com/freman/amber/blob/master/date.go#L13-L15>)



```go
type Date struct {
    time.Time
}
```

<a name="Date.MarshalJSON"></a>
### func \(\*Date\) [MarshalJSON](<https://github.com/freman/amber/blob/master/date.go#L29>)

```go
func (d *Date) MarshalJSON() ([]byte, error)
```



<a name="Date.String"></a>
### func \(Date\) [String](<https://github.com/freman/amber/blob/master/date.go#L33>)

```go
func (d Date) String() string
```



<a name="Date.UnmarshalJSON"></a>
### func \(\*Date\) [UnmarshalJSON](<https://github.com/freman/amber/blob/master/date.go#L17>)

```go
func (d *Date) UnmarshalJSON(in []byte) (err error)
```



<a name="Descriptor"></a>
## type [Descriptor](<https://github.com/freman/amber/blob/master/interval.go#L189>)



```go
type Descriptor string
```

<a name="DescriptorNegative"></a>

```go
const (
    DescriptorNegative     Descriptor = "negative"
    DescriptorExtremelyLow Descriptor = "extremelyLow"
    DescriptorVeryLow      Descriptor = "veryLow"
    DescriptorLow          Descriptor = "low"
    DescriptorNeutral      Descriptor = "neutral"
    DescriptorHigh         Descriptor = "high"
    DescriptorSpike        Descriptor = "spike"
)
```

<a name="Duration"></a>
## type [Duration](<https://github.com/freman/amber/blob/master/interval.go#L100>)



```go
type Duration time.Duration
```

<a name="IntervalDuration5"></a>

```go
const (
    IntervalDuration5  Duration = Duration(time.Minute * 5)
    IntervalDuration15 Duration = Duration(time.Minute * 15)
    IntervalDuration30 Duration = Duration(time.Minute * 30)
)
```

<a name="Duration.MarshalJSON"></a>
### func \(\*Duration\) [MarshalJSON](<https://github.com/freman/amber/blob/master/interval.go#L121>)

```go
func (id *Duration) MarshalJSON() ([]byte, error)
```



<a name="Duration.UnmarshalJSON"></a>
### func \(\*Duration\) [UnmarshalJSON](<https://github.com/freman/amber/blob/master/interval.go#L108>)

```go
func (id *Duration) UnmarshalJSON(in []byte) error
```



<a name="ForecastInterval"></a>
## type [ForecastInterval](<https://github.com/freman/amber/blob/master/interval.go#L48-L52>)



```go
type ForecastInterval struct {
    BaseInterval
    Range         Range         `json:"range,omitempty"`
    AdvancedPrice AdvancedPrice `json:"advancedPrice,omitempty"`
}
```

<a name="ForecastInterval.GetType"></a>
### func \(ForecastInterval\) [GetType](<https://github.com/freman/amber/blob/master/interval.go#L54>)

```go
func (b ForecastInterval) GetType() string
```



<a name="HTTPError"></a>
## type [HTTPError](<https://github.com/freman/amber/blob/master/httperr.go#L10-L15>)



```go
type HTTPError struct {
    StatusCode int
    Status     string
    Body       []byte
    Err        error
}
```

<a name="HTTPError.Error"></a>
### func \(HTTPError\) [Error](<https://github.com/freman/amber/blob/master/httperr.go#L17>)

```go
func (err HTTPError) Error() string
```



<a name="Interval"></a>
## type [Interval](<https://github.com/freman/amber/blob/master/interval.go#L14-L17>)



```go
type Interval interface {
    GetType() string
    GetChannelType() ChannelType
}
```

<a name="IntervalLength"></a>
## type [IntervalLength](<https://github.com/freman/amber/blob/master/interval.go#L77>)



```go
type IntervalLength int
```

<a name="IntervalLength5Mins"></a>

```go
const (
    IntervalLength5Mins  IntervalLength = 5
    IntervalLength10Mins IntervalLength = 10
)
```

<a name="IntervalLength.Duration"></a>
### func \(IntervalLength\) [Duration](<https://github.com/freman/amber/blob/master/interval.go#L88>)

```go
func (i IntervalLength) Duration() time.Duration
```



<a name="IntervalLength.Valid"></a>
### func \(IntervalLength\) [Valid](<https://github.com/freman/amber/blob/master/interval.go#L84>)

```go
func (i IntervalLength) Valid() bool
```



<a name="IntervalMap"></a>
## type [IntervalMap](<https://github.com/freman/amber/blob/master/interval.go#L131>)

IntervalMap is a generic map of ChannelType to IntervalSlice

```go
type IntervalMap[T Interval] map[ChannelType]IntervalSlice[T]
```

<a name="IntervalMap[T].UnmarshalJSON"></a>
### func \(\*IntervalMap\[T\]\) [UnmarshalJSON](<https://github.com/freman/amber/blob/master/interval.go#L134>)

```go
func (im *IntervalMap[T]) UnmarshalJSON(data []byte) error
```

Custom JSON unmarshaling for IntervalMap

<a name="IntervalSlice"></a>
## type [IntervalSlice](<https://github.com/freman/amber/blob/master/interval.go#L128>)

IntervalSlice is a generic slice of Intervals

```go
type IntervalSlice[T Interval] []T
```

<a name="Option"></a>
## type [Option](<https://github.com/freman/amber/blob/master/options.go#L10>)



```go
type Option func(c *Client)
```

<a name="WithBaseURL"></a>
### func [WithBaseURL](<https://github.com/freman/amber/blob/master/options.go#L18>)

```go
func WithBaseURL(uri string) Option
```



<a name="WithHTTPClient"></a>
### func [WithHTTPClient](<https://github.com/freman/amber/blob/master/options.go#L12>)

```go
func WithHTTPClient(client *http.Client) Option
```



<a name="Period"></a>
## type [Period](<https://github.com/freman/amber/blob/master/tariff.go#L17>)

Period is the Time of Use period that is currently active.

```go
type Period string
```

<a name="PeriodOffPeak"></a>

```go
const (
    PeriodOffPeak     Period = "offPeak"
    PeriodShoulder    Period = "shoulder"
    PeriodSolarSponge Period = "solarSponge"
    PeriodPeak        Period = "peak"
)
```

<a name="Quality"></a>
## type [Quality](<https://github.com/freman/amber/blob/master/interval.go#L201>)



```go
type Quality string
```

<a name="QualityEstimate"></a>

```go
const (
    QualityEstimate Quality = "estimated"
    QualityBillable Quality = "billable"
)
```

<a name="QueryArgument"></a>
## type [QueryArgument](<https://github.com/freman/amber/blob/master/queryArgs.go#L14>)



```go
type QueryArgument func(v url.Values)
```

<a name="WithEndDate"></a>
### func [WithEndDate](<https://github.com/freman/amber/blob/master/queryArgs.go#L40>)

```go
func WithEndDate(date time.Time) QueryArgument
```



<a name="WithNext"></a>
### func [WithNext](<https://github.com/freman/amber/blob/master/queryArgs.go#L16>)

```go
func WithNext(intervals uint) QueryArgument
```



<a name="WithPrevious"></a>
### func [WithPrevious](<https://github.com/freman/amber/blob/master/queryArgs.go#L22>)

```go
func WithPrevious(intervals uint) QueryArgument
```



<a name="WithResolution"></a>
### func [WithResolution](<https://github.com/freman/amber/blob/master/queryArgs.go#L28>)

```go
func WithResolution(resolution uint) QueryArgument
```



<a name="WithStartDate"></a>
### func [WithStartDate](<https://github.com/freman/amber/blob/master/queryArgs.go#L34>)

```go
func WithStartDate(date time.Time) QueryArgument
```



<a name="Range"></a>
## type [Range](<https://github.com/freman/amber/blob/master/interval.go#L209-L212>)

Range is used when prices are particularly volatile, the API may return a range of prices that are possible.

```go
type Range struct {
    Min float64 `json:"min"` // Estimated minimum price (c/kWh)
    Max float64 `json:"max"` // Estimated maximum price (c/kWh)
}
```

<a name="Renewable"></a>
## type [Renewable](<https://github.com/freman/amber/blob/master/renewable.go#L11-L20>)

Renewable data.

```go
type Renewable struct {
    Type       string              `json:"type"`
    Duration   Duration            `json:"duration"`   // Length of the interval in minutes.
    Date       Date                `json:"date"`       // Date the interval belongs to (in NEM time). This may be different to the date component of nemTime, as the last interval of the day ends at 12:00 the following day. [2006-01-02]
    NemTime    time.Time           `json:"nemTime"`    // The interval's NEM time. This represents the time at the end of the interval UTC+10. Formatted as a ISO 8601 time
    StartTime  time.Time           `json:"startTime"`  // Start time of the interval in UTC. Formatted as a ISO 8601 time
    EndTime    time.Time           `json:"endTime"`    // End time of the interval in UTC. Formatted as a ISO 8601 time
    Renewables float64             `json:"renewables"` // Percentage of renewables in the grid 0-100
    Descriptor RenewableDescriptor `json:"descriptor"` // Describes the state of renewables. Gives you an indication of how green power is right now
}
```

<a name="RenewableDescriptor"></a>
## type [RenewableDescriptor](<https://github.com/freman/amber/blob/master/renewable.go#L22>)



```go
type RenewableDescriptor string
```

<a name="RenewableDescriptorBest"></a>

```go
const (
    RenewableDescriptorBest     RenewableDescriptor = "best"
    RenewableDescriptorGreat    RenewableDescriptor = "great"
    RenewableDescriptorOk       RenewableDescriptor = "ok"
    RenewableDescriptorNotGreat RenewableDescriptor = "notGreat"
    RenewableDescriptorWorst    RenewableDescriptor = "worst"
)
```

<a name="Season"></a>
## type [Season](<https://github.com/freman/amber/blob/master/tariff.go#L27>)

Season is tThe Time of Use season that is currently active.

```go
type Season string
```

<a name="SeasonDefault"></a>

```go
const (
    SeasonDefault        Season = "default"
    SeasonSummer         Season = "summer"
    SeasonAutumn         Season = "autumn"
    SeasonWinter         Season = "winter"
    SeasonSpring         Season = "spring"
    SeasonNonSummer      Season = "nonSummer"
    SeasonHoliday        Season = "holiday"
    SeasonWeekend        Season = "weekend"
    SeasonWeekendHoliday Season = "weekendHoliday"
    SeasonWeekday        Season = "weekday"
)
```

<a name="Site"></a>
## type [Site](<https://github.com/freman/amber/blob/master/site.go#L8-L17>)



```go
type Site struct {
    ID             string         `json:"id"`                       // Unique Site Identifier.
    NMI            string         `json:"nmi"`                      // National Metering Identifier (NMI) for the site.
    Channels       []Channel      `json:"channels"`                 // List of channels that are readable from your meter.
    Network        string         `json:"network"`                  // The name of the site's network.
    Status         SiteStatus     `json:"status"`                   // Site status. [pending, active, closed]
    ActiveFrom     Date           `json:"activeFrom,omitempty"`     // Date the site became active [2006-01-02].
    ClosedOn       Date           `json:"closedOn,omitempty"`       // Date the site closed.
    IntervalLength IntervalLength `json:"intervalLength,omitempty"` // Length of interval that you will be billed on. [5, 30] minutes
}
```

<a name="SiteStatus"></a>
## type [SiteStatus](<https://github.com/freman/amber/blob/master/site.go#L19>)



```go
type SiteStatus string
```

<a name="SpikeStatus"></a>
## type [SpikeStatus](<https://github.com/freman/amber/blob/master/interval.go#L214>)



```go
type SpikeStatus string
```

<a name="SpikeStatusNone"></a>

```go
const (
    SpikeStatusNone      SpikeStatus = "none"
    SpikeStatusPotential SpikeStatus = "potential"
    SpikeStatusSpike     SpikeStatus = "spike"
)
```

<a name="Tariff"></a>
## type [Tariff](<https://github.com/freman/amber/blob/master/tariff.go#L9-L14>)

Tariff information about how your tariff affects an interval.

```go
type Tariff struct {
    Period Period `json:"period,omitempty"`       // The Time of Use period that is currently active. [offPeak, shoulder, solarSponge, peak]
    Season Season `json:"season,omitempty"`       // The Time of Use season that is currently active. [default, summer, autumn, winter, spring, nonSummer, holiday, weekend, weekendHoliday, weekday]
    Block  int    `json:"block,omitempty"`        // The block that is currently active.
    Demand bool   `json:"demandWindow,omitempty"` // Is this interval currently in the demand window?
}
```

<a name="Usage"></a>
## type [Usage](<https://github.com/freman/amber/blob/master/interval.go#L69-L75>)



```go
type Usage struct {
    BaseInterval
    ChannelIdentifier string  `json:"channelIdentifier"` // Meter channel identifier.
    KWh               float64 `json:"kwh"`               // Number of kWh you consumed or generated. Generated numbers will be negative.
    Quality           Quality `json:"quality"`           // If the metering company has had trouble contacting your meter, they may make an estimate of your usage for that period. Billable data is data that will appear on your bill.
    Cost              float64 `json:"cost"`              // The total cost of your consumption or generation for this period - includes GST
}
```