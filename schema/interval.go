// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

import (
	"time"
)

// Interval is one time interval.
type Interval struct {
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
	Tarrif      Tariff      `json:"tariffInformation"` // Information about how your tariff is affecting this interval
	SpikeStatus SpikeStatus `json:"spikeStatus"`       // Indicates whether this interval will potentially spike, or is currently in a spike state. [none, potential, spike]
	Descriptor  Descriptor  `json:"descriptor"`        // Describes the current price. Gives you an indication of how cheap the price is in relation to the average VMO and DMO. Note: Negative is no longer used. It has been replaced with extremelyLow. [ negative, extremelyLow, veryLow, low, neutral, high, spike ]
}
