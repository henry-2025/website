package internal

import (
	"strings"
	"time"
)

type Project struct {
	Date        Date   `json:"date"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Path        string
	Source      []byte
}

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	value := strings.Trim(string(b), `"`) //get rid of "
	if value == "" || value == "null" {
		return nil
	}

	t, err := time.Parse("2006-01-02", value) //parse time
	if err != nil {
		return err
	}
	*d = Date(t) //set result using the pointer
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format("2006-01-02") + `"`), nil
}

func (d Date) String() string {
	return time.Time(d).Format("2006-01-02")
}
