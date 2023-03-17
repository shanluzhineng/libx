package timeonly

import (
	"errors"
	"fmt"
	"time"
)

var (
	Error_TimeOnlyParse = errors.New("TimeOnlyParseError: should be a string formatted as '15:04:05'")
)

// Represents a time of day, as would be read from a clock, within the range 00:00:00 to 23:59:59
type TimeOnly struct {
	_time time.Time
}

// create a instance
func NewTimeOnly(hour int, minute int, second int) (TimeOnly, error) {
	s := fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
	return Parse(s)
}

func NewTimeOnlyFromTime(time time.Time) TimeOnly {
	return TimeOnly{
		_time: time,
	}
}

// Now returns the current time.
func NewTimeOnlyFromNow() TimeOnly {
	return TimeOnly{
		_time: time.Now(),
	}
}

func Parse(s string) (TimeOnly, error) {
	time, err := time.Parse(time.TimeOnly, s)
	if err != nil {
		return TimeOnly{}, Error_TimeOnlyParse
	}
	return TimeOnly{
		_time: time,
	}, nil
}

// MarshalJSON implements the json.Marshaler interface.
func (t TimeOnly) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t._time.Format(time.TimeOnly) + `"`), nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *TimeOnly) UnmarshalJSON(b []byte) error {
	s := string(b)
	// len(`"23:59:00"`) == 7
	if len(s) != len(time.TimeOnly) {
		return Error_TimeOnlyParse
	}
	ret, err := time.Parse(time.TimeOnly, s[1:len(s)-2])
	if err != nil {
		return err
	}
	t._time = ret
	return nil
}

func (t TimeOnly) ToString() string {
	return t._time.Format(time.TimeOnly)
}

func (t TimeOnly) ToTime(year int, month int, day int) time.Time {
	s := fmt.Sprintf("%04d-%02d-%02d", year, month, day)
	s = s + " " + t.ToString()
	time, _ := time.Parse(s, time.DateTime)
	return time
}

// After reports whether the TimeOnly instant t is after u.
func (t TimeOnly) After(u TimeOnly) bool {
	return t._time.After(u._time)
}

// Before reports whether the TimeOnly instant t is before u.
func (t TimeOnly) Before(u TimeOnly) bool {
	return t._time.Before(u._time)
}

// Compare compares the TimeOnly instant t with u. If t is before u, it returns -1;
// if t is after u, it returns +1; if they're the same, it returns 0.
func (t TimeOnly) Compare(u TimeOnly) int {
	return t._time.Compare(u._time)
}

// Equal reports whether t and u represent the same TimeOnly instant.
func (t TimeOnly) Equal(u TimeOnly) bool {
	return t._time.Equal(u._time)
}

// Clock returns the hour, minute, and second within the day specified by t.
func (t TimeOnly) Clock() (hour, min, sec int) {
	return t._time.Clock()
}

// Hour returns the hour within the day specified by t, in the range [0, 23].
func (t TimeOnly) Hour() int {
	return t._time.Hour()
}

// Minute returns the minute offset within the hour specified by t, in the range [0, 59].
func (t TimeOnly) Minute() int {
	return t._time.Minute()
}

// Second returns the second offset within the minute specified by t, in the range [0, 59].
func (t TimeOnly) Second() int {
	return t._time.Second()
}

// Nanosecond returns the nanosecond offset within the second specified by t,
// in the range [0, 999999999].
func (t TimeOnly) Nanosecond() int {
	return t._time.Nanosecond()
}

// Sub returns the duration t-u. If the result exceeds the maximum (or minimum)
// value that can be stored in a Duration, the maximum (or minimum) duration
// will be returned.
func (t TimeOnly) Sub(u TimeOnly) time.Duration {
	return t._time.Sub(u._time)
}
