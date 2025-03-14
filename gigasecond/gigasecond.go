// package gigasecond provides utility to add gigasecond to time.
package gigasecond

import (
	"time"
)

// AddGigasecond adds 1 gigasecond to the provided time.
func AddGigasecond(t time.Time) time.Time {
	future := t.Add(1_000_000_000 * time.Second)

	return future
}
