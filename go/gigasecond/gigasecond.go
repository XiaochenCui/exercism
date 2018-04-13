package gigasecond

// import path for the time package from the standard library
import "time"
import "math"

// AddGigasecond should have a comment documenting it.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * time.Duration(int(math.Pow10(9))))
}
