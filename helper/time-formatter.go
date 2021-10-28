package helper

import (
	"fmt"
	"time"
)

func FormatTime(date string) string {
	layout := "2006-01-02"
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Time parsing error ", err)
	}
	//Format ISO 8601 (RFC 3339)
	//Like this: 2017-01-28T01:22:14.398+00:00
	newTime := t.Format(time.RFC3339)

	return newTime
}
