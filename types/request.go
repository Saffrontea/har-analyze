package types

type Request struct {
	Method      string `json:"method"`
	URL         string `json:"url"`
	HTTPVersion string `json:"httpVersion"`
	Headers     []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	QueryString []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"queryString"`
	Cookies     []Cookie `json:"cookies"`
	HeadersSize int      `json:"headersSize"`
	BodySize    int      `json:"bodySize"`
}
