// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

// Tariff information about how your tariff affects an interval.
type Tariff struct {
	Period Period `json:"period,omitempty"`       // The Time of Use period that is currently active. [offPeak, shoulder, solarSponge, peak]
	Season Season `json:"season,omitempty"`       // The Time of Use season that is currently active. [default, summer, autumn, winter, spring, nonSummer, holiday, weekend, weekendHoliday, weekday]
	Block  int    `json:"block,omitempty"`        // The block that is currently active.
	Demand bool   `json:"demandWindow,omitempty"` // Is this interval currently in the demand window?
}
