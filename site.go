// Copyright 2025 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package amber

type Site struct {
	ID             string         `json:"id"`                       // Unique Site Identifier.
	NMI            string         `json:"nmi"`                      // National Metering Identifier (NMI) for the site.
	Channels       []Channel      `json:"channels"`                 // List of channels that are readable from your meter.
	Network        string         `json:"network"`                  // The name of the site's network.
	Status         SiteStatus     `json:"status"`                   // Site status. [pending, active, closed]
	ActiveFrom     Date           `json:"activeFrom,omitempty"`     // Date the site became active [2006-01-02].
	ClosedOn       Date           `json:"closedOn,omitempty"`       // Date the site closed.
	IntervalLength IntervalLength `json:"intervalLength,omitempty"` // Length of interval that you will be billed on. [5, 30] minutes
}

type SiteStatus string

const (
	SiteStatusPending ChannelType = "pending"
	SiteStatusActive  ChannelType = "active"
	SiteStatusClosed  ChannelType = "cloased"
)
