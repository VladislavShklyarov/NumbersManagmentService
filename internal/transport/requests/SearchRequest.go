package requests

type SearchRequest struct {
	Number   string `query:"number"`
	Country  string `query:"country"`
	Region   string `query:"region"`
	Provider string `query:"provider"`
	Limit    int    `query:"limit"`
	Offset   int    `query:"offset"`
}
