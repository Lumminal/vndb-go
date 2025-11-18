package wrapper

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	releaseDateFull    = "2006-01-02"
	releaseDateMonth   = "2006-01"
	releaseDateYear    = "2006"
	releaseDateTBA     = "TBA"
	releaseDateUnknown = "unknown"
	releaseDateToday   = "today"
)

type ReleaseDate struct {
	Release string
	Time    time.Time
}

func (rd *ReleaseDate) UnmarshalJSON(data []byte) error {
	var stu string // string to unmarshal
	if err := json.Unmarshal(data, &stu); err != nil {
		return err
	}

	rd.Release = stu

	if rd.Release == releaseDateTBA || rd.Release == releaseDateUnknown {
		rd.Time = time.Time{}
		return nil
	}

	if rd.Release == releaseDateToday {
		rd.Time = time.Now()
		return nil
	}

	// Parse all dates - Start
	if t, err := time.Parse(releaseDateFull, rd.Release); err == nil {
		rd.Time = t
		return nil
	}

	if t, err := time.Parse(releaseDateMonth, rd.Release); err == nil {
		rd.Time = t
		return nil
	}

	if t, err := time.Parse(releaseDateYear, rd.Release); err == nil {
		rd.Time = t
		return nil
	}
	// Parse all dates - End

	return fmt.Errorf("invalid release date format: %s", rd.Release)
}
