// Copyright 2025 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package amber

// Tariff information about how your tariff affects an interval.
type Tariff struct {
	Period Period `json:"period,omitempty"`       // The Time of Use period that is currently active. [offPeak, shoulder, solarSponge, peak]
	Season Season `json:"season,omitempty"`       // The Time of Use season that is currently active. [default, summer, autumn, winter, spring, nonSummer, holiday, weekend, weekendHoliday, weekday]
	Block  int    `json:"block,omitempty"`        // The block that is currently active.
	Demand bool   `json:"demandWindow,omitempty"` // Is this interval currently in the demand window?
}

// Period is the Time of Use period that is currently active.
type Period string

const (
	PeriodOffPeak     Period = "offPeak"
	PeriodShoulder    Period = "shoulder"
	PeriodSolarSponge Period = "solarSponge"
	PeriodPeak        Period = "peak"
)

// Season is tThe Time of Use season that is currently active.
type Season string

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
