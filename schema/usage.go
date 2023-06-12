// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

type Usage struct {
	Interval
	ChannelIdentifier string  `json:"channelIdentifier"` // Meter channel identifier.
	KWh               float64 `json:"kwh"`               // Number of kWh you consumed or generated. Generated numbers will be negative.
	Quality           Quality `json:"quality"`           // If the metering company has had trouble contacting your meter, they may make an estimate of your usage for that period. Billable data is data that will appear on your bill.
	Cost              float64 `json:"cost"`              // The total cost of your consumption or generation for this period - includes GST
}
