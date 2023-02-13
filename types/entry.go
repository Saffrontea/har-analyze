package types

import "time"

type Entry struct {
	Priority     string `json:"_priority"`
	ResourceType string `json:"_resourceType"`
	Cache        struct {
	} `json:"cache"`
	Connection      string    `json:"connection"`
	Pageref         string    `json:"pageref"`
	Request         Request   `json:"request"`
	Response        Response  `json:"response"`
	ServerIPAddress string    `json:"serverIPAddress"`
	StartedDateTime time.Time `json:"startedDateTime"`
	Time            float64   `json:"time"`
	Timings         struct {
		Blocked         float64 `json:"blocked"`
		DNS             int     `json:"dns"`
		Ssl             int     `json:"ssl"`
		Connect         int     `json:"connect"`
		Send            float64 `json:"send"`
		Wait            float64 `json:"wait"`
		Receive         float64 `json:"receive"`
		BlockedQueueing float64 `json:"_blocked_queueing"`
	} `json:"timings"`
	Initiator struct {
		Type       string `json:"type"`
		URL        string `json:"url,omitempty"`
		LineNumber int    `json:"lineNumber,omitempty"`
		Stack      struct {
			CallFrames []struct {
				FunctionName string `json:"functionName"`
				ScriptID     string `json:"scriptId"`
				URL          string `json:"url"`
				LineNumber   int    `json:"lineNumber"`
				ColumnNumber int    `json:"columnNumber"`
			} `json:"callFrames,omitempty"`
			Parent Parent `json:"parent,omitempty"`
		} `json:"stack,omitempty"`
	} `json:"_initiator,omitempty"`
}
