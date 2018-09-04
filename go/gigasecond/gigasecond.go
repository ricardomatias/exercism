package gigasecond

import (
	"time"
)

// A gigasecond is 10**9 (1,000,000,000) seconds.
const gigasecond time.Duration = 1000000000

func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * gigasecond)
}
