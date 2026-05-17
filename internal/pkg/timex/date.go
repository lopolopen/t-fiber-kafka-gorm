package timex

import (
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) String() string {
	return d.Time.Format(time.DateOnly)
}

// MarshalJSON: convert Date to JSON string in "yyyy-MM-dd" format
func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		// encode zero value as empty string
		return []byte(`""`), nil
	}
	s := fmt.Sprintf("\"%s\"", d)
	return []byte(s), nil
}

// UnmarshalJSON: parse JSON string in "yyyy-MM-dd" format into Date
func (d *Date) UnmarshalJSON(b []byte) error {
	// remove surrounding quotes
	s := string(b)
	if len(s) >= 2 && s[0] == '"' && s[len(s)-1] == '"' {
		s = s[1 : len(s)-1]
	}
	if s == "" {
		*d = Date{time.Time{}}
		return nil
	}
	t, err := time.ParseInLocation(time.DateOnly, s, time.Local)
	if err != nil {
		return err
	}
	d.Time = t
	return nil
}

func NowDate() Date {
	return Date{
		Time: time.Now().Truncate(time.Hour),
	}
}

func TimeToDate(t time.Time) Date {
	return Date{
		Time: t.Truncate(time.Hour),
	}
}
