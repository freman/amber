// Copyright 2023 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package schema

import (
	"encoding/json"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) UnmarshalJSON(in []byte) (err error) {
	var tmp string

	if err := json.Unmarshal(in, &tmp); err != nil {
		return err
	}

	d.Time, err = time.Parse(time.DateOnly, tmp)

	return err
}

func (d *Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(time.DateOnly))
}
