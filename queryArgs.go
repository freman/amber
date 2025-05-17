// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package amber

import (
	"net/url"
	"strconv"
	"time"
)

type QueryArgument func(v url.Values)

func WithNext(intervals uint) QueryArgument {
	return func(v url.Values) {
		v.Set("next", strconv.Itoa(int(intervals)))
	}
}

func WithPrevious(intervals uint) QueryArgument {
	return func(v url.Values) {
		v.Set("previous", strconv.Itoa(int(intervals)))
	}
}

func WithResolution(resolution uint) QueryArgument {
	return func(v url.Values) {
		v.Set("resolution", strconv.Itoa(int(resolution)))
	}
}

func WithStartDate(date time.Time) QueryArgument {
	return func(v url.Values) {
		v.Set("startDate", date.Format(time.DateOnly))
	}
}

func WithEndDate(date time.Time) QueryArgument {
	return func(v url.Values) {
		v.Set("endDate", date.Format(time.DateOnly))
	}
}
