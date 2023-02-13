package types

type Parent struct {
	Description string `json:"description"`
	CallFrames  []struct {
		FunctionName string `json:"functionName"`
		ScriptID     string `json:"scriptId"`
		URL          string `json:"url"`
		LineNumber   int    `json:"lineNumber"`
		ColumnNumber int    `json:"columnNumber"`
	} `json:"callFrames"`
	Parent *Parent `json:"parent,omitempty"`
}
