// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

import (
	"encoding/json"
	"fmt"
	"time"
)

type Interval interface {
	GetType() string
	GetChannelType() ChannelType
}

// Interval is one time interval.
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

func (b BaseInterval) GetChannelType() ChannelType {
	return b.ChannelType
}

type ActualInterval struct {
	BaseInterval
}

func (b ActualInterval) GetType() string {
	return b.Type
}

type ForecastInterval struct {
	BaseInterval
	Range         Range         `json:"range,omitempty"`
	AdvancedPrice AdvancedPrice `json:"advancedPrice,omitempty"`
}

func (b ForecastInterval) GetType() string {
	return b.Type
}

type CurrentInterval struct {
	BaseInterval
	Range         Range         `json:"range,omitempty"`
	AdvancedPrice AdvancedPrice `json:"advancedPrice,omitempty"`
	Estimate      bool          `json:"estimate,omitempty"`
}

func (b CurrentInterval) GetType() string {
	return b.Type
}

type Usage struct {
	BaseInterval
	ChannelIdentifier string  `json:"channelIdentifier"` // Meter channel identifier.
	KWh               float64 `json:"kwh"`               // Number of kWh you consumed or generated. Generated numbers will be negative.
	Quality           Quality `json:"quality"`           // If the metering company has had trouble contacting your meter, they may make an estimate of your usage for that period. Billable data is data that will appear on your bill.
	Cost              float64 `json:"cost"`              // The total cost of your consumption or generation for this period - includes GST
}

type IntervalLength int

const (
	IntervalLength5Mins  IntervalLength = 5
	IntervalLength10Mins IntervalLength = 10
)

func (i IntervalLength) Valid() bool {
	return i == IntervalLength5Mins || i == IntervalLength10Mins
}

func (i IntervalLength) Duration() time.Duration {
	return time.Second * time.Duration(i)
}

// Advamcedprice is part of Ambers advanced forecast system that represents their confidence in the AEMO
// forecast. The range indicates where they think the price will land for a given interval.
type AdvancedPrice struct {
	Low       float64 `json:"low,omitempty"`       // The lower bound of prediction band. (c/kWh).
	Predicted float64 `json:"predicted,omitempty"` // The predicted price. Use this if you need a single number to forecast against. (c/kWh).
	High      float64 `json:"high,omitempty"`      // The upper bound of prediction band. Price includes network and market fees. (c/kWh).
}

type Duration time.Duration

const (
	IntervalDuration5  Duration = Duration(time.Minute * 5)
	IntervalDuration15 Duration = Duration(time.Minute * 15)
	IntervalDuration30 Duration = Duration(time.Minute * 30)
)

func (id *Duration) UnmarshalJSON(in []byte) error {
	var i int64
	if err := json.Unmarshal(in, &i); err != nil {
		return err
	}

	d := Duration(time.Duration(i) * time.Minute)

	*id = d

	return nil
}

func (id *Duration) MarshalJSON() ([]byte, error) {
	i := time.Duration(*id) / time.Minute

	return json.Marshal(i)
}

// IntervalSlice is a generic slice of Intervals
type IntervalSlice[T Interval] []T

// IntervalMap is a generic map of ChannelType to IntervalSlice
type IntervalMap[T Interval] map[ChannelType]IntervalSlice[T]

// Custom JSON unmarshaling for IntervalMap
func (im *IntervalMap[T]) UnmarshalJSON(data []byte) error {
	// Temporary slice to hold raw JSON objects
	var raw []json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	// Initialize the map
	*im = make(IntervalMap[T])

	for _, item := range raw {
		// Peek at the type and channelType fields
		var typePeek struct {
			Type        string      `json:"type"`
			ChannelType ChannelType `json:"channelType"`
		}
		if err := json.Unmarshal(item, &typePeek); err != nil {
			return err
		}

		if typePeek.ChannelType == "" {
			typePeek.ChannelType = "unknown"
		}

		var interval T
		switch typePeek.Type {
		case "ActualInterval":
			var ai ActualInterval
			if err := json.Unmarshal(item, &ai); err != nil {
				return err
			}
			interval, _ = any(ai).(T)
		case "ForecastInterval":
			var fi ForecastInterval
			if err := json.Unmarshal(item, &fi); err != nil {
				return err
			}
			interval, _ = any(fi).(T)
		case "CurrentInterval":
			var ci CurrentInterval
			if err := json.Unmarshal(item, &ci); err != nil {
				return err
			}
			interval, _ = any(ci).(T)
		default:
			return fmt.Errorf("unknown interval type: %s", typePeek.Type)
		}

		// Group by ChannelType
		(*im)[typePeek.ChannelType] = append((*im)[typePeek.ChannelType], interval)
	}

	return nil
}
