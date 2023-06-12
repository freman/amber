// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

import (
	"encoding/json"
	"time"
)

type Duration time.Duration

const (
	IntervalDuration5  Duration = Duration(time.Minute * 5)
	IntervalDuration15 Duration = Duration(time.Minute * 15)
	IntervalDuration31 Duration = Duration(time.Minute * 30)
)

func (id *Duration) UnmarshalJSON(in []byte) error {
	var i int64
	if err := json.Unmarshal(in, &i); err != nil {
		return err
	}

	d := Duration(time.Duration(i) * time.Minute)

	*id = d

	return nil
}

func (id *Duration) MarshalJSON() ([]byte, error) {
	i := time.Duration(*id) / time.Minute

	return json.Marshal(i)
}
