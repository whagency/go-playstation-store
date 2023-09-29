package response

type ErrorsList []struct {
	Message    string `json:"message"`
	Type       string `json:"type"`
	StatusCode int    `json:"statusCode"`
	ApiName    string `json:"apiName"`
}

type Response struct {
	StatusCode int
	Message    string      `json:"message"`
	Errors     ErrorsList  `json:"errors"`
	Data       interface{} `json:"data"`
}
