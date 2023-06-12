// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

import "time"

// Renewable data.
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
