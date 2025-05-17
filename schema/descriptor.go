// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

type Descriptor string

const (
	DescriptorNegative     Descriptor = "negative"
	DescriptorExtremelyLow Descriptor = "extremelyLow"
	DescriptorVeryLow      Descriptor = "veryLow"
	DescriptorLow          Descriptor = "low"
	DescriptorNeutral      Descriptor = "neutral"
	DescriptorHigh         Descriptor = "high"
	DescriptorSpike        Descriptor = "spike"
)

type RenewableDescriptor string

const (
	RenewableDescriptorBest     RenewableDescriptor = "best"
	RenewableDescriptorGreat    RenewableDescriptor = "great"
	RenewableDescriptorOk       RenewableDescriptor = "ok"
	RenewableDescriptorNotGreat RenewableDescriptor = "notGreat"
	RenewableDescriptorWorst    RenewableDescriptor = "worst"
)
