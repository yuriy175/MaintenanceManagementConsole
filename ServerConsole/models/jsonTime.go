package models

import (
	"encoding/json"
	"strings"
	"time"
)

// First create a type alias
type JsonTime time.Time

// Implement Marshaler and Unmarshaler interface
func (j *JsonTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JsonTime(t)
	return nil
}

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j JsonTime) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
