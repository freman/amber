// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

// Range is used when prices are particularly volatile, the API may return a range of prices that are possible.
type Range struct {
	Min float64 `json:"min"` // Estimated minimum price (c/kWh)
	Max float64 `json:"max"` // Estimated maximum price (c/kWh)
}
