//go:build !darwin

package timestamp

func v2Now() int64 {
	return v1Now()
}
