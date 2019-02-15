package utils

import (
	"fmt"
	"strconv"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(*t).Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	*t = Timestamp(time.Unix(int64(ts), 0))

	return nil
}

func (t *Timestamp) IsZero() bool {
	return time.Time(*t).IsZero()
}

func (t *Timestamp) Time() time.Time {
	return time.Time(*t)
}
