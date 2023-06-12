// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

type Quality string

const (
	QualityEstimate Quality = "estimated"
	QualityBillable Quality = "billable"
)
