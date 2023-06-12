// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

type SpikeStatus string

const (
	SpikeStatusNone      SpikeStatus = "none"
	SpikeStatusPotential SpikeStatus = "potential"
	SpikeStatusSpike     SpikeStatus = "spike"
)
