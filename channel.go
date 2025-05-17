// Copyright 2025 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package amber

// Channel Describes a power meter channel.
//
// The General channel provides continuous power - it's the channel all of your appliances and lights are attached to.
//
// Controlled loads are only on for a limited time during the day (usually when the load on the network is
// low, or generation is high) - you may have your hot water system attached to this channel.
//
// The feed in channel sends power back to the grid - you will have these types of channels
// if you have solar or batteries.
type Channel struct {
	Identifier string      `json:"identifier"` // Identifier of the channel.
	Type       ChannelType `json:"type"`       // Channel type. [general, controlledLoad, feedIn]
	Tariff     string      `json:"tariff"`     // The tariff code of the channel.
}

type ChannelType string

const (
	ChannelTypeGeneral        ChannelType = "general"
	ChannelTypeControlledLoad ChannelType = "controlledLoad"
	ChannelTypeFeedIn         ChannelType = "feedIn"
)
