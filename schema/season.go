// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

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
