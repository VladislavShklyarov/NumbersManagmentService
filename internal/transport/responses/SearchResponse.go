package responses

type SearchResponse struct {
	Total   int        `json:"total"`
	Results []PhoneDTO `json:"results"`
}
