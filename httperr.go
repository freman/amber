// Copyright 2025 Shannon Wynter
//
// This software may be modified and distributed under the terms
// of the MIT license.  See the LICENSE file for details.

package amber

import "fmt"

type HTTPError struct {
	StatusCode int
	Status     string
	Body       []byte
	Err        error
}

func (err HTTPError) Error() string {
	msg := fmt.Sprintf("unexpected http status %s (%d) was encountered", err.Status, err.StatusCode)

	if l := len(err.Body); l > 0 {
		msg += fmt.Sprintf(", http body contained %d bytes", l)
	}

	if err.Err != nil {
		msg += fmt.Sprintf(", additionally %v was encountered reading body", err.Err)
	}

	return msg
}
