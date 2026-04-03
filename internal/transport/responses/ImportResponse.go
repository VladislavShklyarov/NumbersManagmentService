package responses

type ImportResponse struct {
	Accepted   int `json:"accepted"`
	Duplicates int `json:"duplicates"`
	Errors     int `json:"errors"`
}
