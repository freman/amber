// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

type Period string

const (
	PeriodOffPeak     Period = "offPeak"
	PeriodShoulder    Period = "shoulder"
	PeriodSolarSponge Period = "solarSponge"
	PeriodPeak        Period = "peak"
)
