package response

type ErrorsList []struct {
	Message         string          `json:"message"`
	Type            string          `json:"type"`
	StatusCode      int             `json:"statusCode"`
	ApiName         string          `json:"apiName"`
	ErrorDetail     ErrorDetail     `json:"error"`
	ErrorExtensions ErrorExtensions `json:"extensions"`
}

type ErrorDetail struct {
	ReferenceID string `json:"referenceId"`
	Message     string `json:"message"`
	Reason      string `json:"reason"`
	Source      string `json:"source"`
}

type ErrorExtensions struct {
	ErrorDetail ErrorDetail `json:"error"`
}

type Response struct {
	StatusCode int
	Message    string      `json:"message"`
	Errors     ErrorsList  `json:"errors"`
	Data       interface{} `json:"data"`
}
