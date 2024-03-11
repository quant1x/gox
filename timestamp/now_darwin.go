package timestamp

import (
	_ "unsafe" // for go:linkname
)

//go:linkname walltime runtime.walltime
func walltime() (int64, int32)

func v2Now() int64 {
	sec, nsec := walltime()
	sec += int64(offsetInSecondsEastOfUTC)
	milli := sec*MillisecondsPerSecond + int64(nsec)/1e6%MillisecondsPerSecond
	return milli
}
