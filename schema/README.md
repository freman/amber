# Unofficial GO api for [Amber](https://app.amber.com.au/developers/documentation/)

## Types

### type [Channel](/channel.go#L17)

`type Channel struct { ... }`

Channel Describes a power meter channel.

The General channel provides continuous power - it's the channel all of your appliances and lights are attached to.

Controlled loads are only on for a limited time during the day (usually when the load on the network is
low, or generation is high) - you may have your hot water system attached to this channel.

The feed in channel sends power back to the grid - you will have these types of channels
if you have solar or batteries.

### type [ChannelType](/channelType.go#L8)

`type ChannelType string`

### type [Date](/date.go#L13)

`type Date struct { ... }`

#### func (*Date) [MarshalJSON](/date.go#L29)

`func (d *Date) MarshalJSON() ([]byte, error)`

#### func (*Date) [UnmarshalJSON](/date.go#L17)

`func (d *Date) UnmarshalJSON(in []byte) (err error)`

### type [Descriptor](/descriptor.go#L8)

`type Descriptor string`

### type [Duration](/duration.go#L13)

`type Duration time.Duration`

#### func (*Duration) [MarshalJSON](/duration.go#L34)

`func (id *Duration) MarshalJSON() ([]byte, error)`

#### func (*Duration) [UnmarshalJSON](/duration.go#L21)

`func (id *Duration) UnmarshalJSON(in []byte) error`

### type [Interval](/interval.go#L13)

`type Interval struct { ... }`

Interval is one time interval.

### type [Period](/period.go#L8)

`type Period string`

### type [Quality](/quality.go#L8)

`type Quality string`

### type [Range](/range.go#L9)

`type Range struct { ... }`

Range is used when prices are particularly volatile, the API may return a range of prices that are possible.

### type [Renewable](/renewable.go#L11)

`type Renewable struct { ... }`

Renewable data.

### type [RenewableDescriptor](/descriptor.go#L20)

`type RenewableDescriptor string`

### type [Season](/season.go#L8)

`type Season string`

### type [Site](/site.go#L8)

`type Site struct { ... }`

### type [SpikeStatus](/spikeStatus.go#L8)

`type SpikeStatus string`

### type [Tariff](/tariff.go#L9)

`type Tariff struct { ... }`

Tariff information about how your tariff affects an interval.

### type [Usage](/usage.go#L8)

`type Usage struct { ... }`

