package types

type Response struct {
	Status      int    `json:"status"`
	StatusText  string `json:"statusText"`
	HTTPVersion string `json:"httpVersion"`
	Headers     []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	Cookies []Cookie `json:"cookies"`
	Content struct {
		Size     int    `json:"size"`
		MimeType string `json:"mimeType"`
		Text     string `json:"text"`
	} `json:"content"`
	RedirectURL  string `json:"redirectURL"`
	HeadersSize  int    `json:"headersSize"`
	BodySize     int    `json:"bodySize"`
	TransferSize int    `json:"_transferSize"`
	Error        string `json:"_error"`
}
