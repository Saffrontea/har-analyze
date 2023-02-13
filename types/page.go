package types

import "time"

type Page struct {
	StartedDateTime time.Time `json:"startedDateTime"`
	ID              string    `json:"id"`
	Title           string    `json:"title"`
	PageTimings     struct {
		OnContentLoad float64 `json:"onContentLoad"`
		OnLoad        float64 `json:"onLoad"`
	} `json:"pageTimings"`
}
