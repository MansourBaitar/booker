package validate

import (
	"time"
)

const (
	ISOLayout = "2006-01-02"
)

func IsDate(input string) (time.Time, error) {
	return time.Parse(ISOLayout, input)
}
