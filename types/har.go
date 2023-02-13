package types

import (
	"time"
)

type Har struct {
	Log struct {
		Version string `json:"version"`
		Creator struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"creator"`
		Pages []struct {
			StartedDateTime time.Time `json:"startedDateTime"`
			ID              string    `json:"id"`
			Title           string    `json:"title"`
			PageTimings     struct {
				OnContentLoad float64 `json:"onContentLoad"`
				OnLoad        float64 `json:"onLoad"`
			} `json:"pageTimings"`
		} `json:"pages"`
		Entries []Entry `json:"entries"`
	} `json:"log"`
}
