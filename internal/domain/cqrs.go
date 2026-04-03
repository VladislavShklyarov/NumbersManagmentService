package domain

type ImportCommand struct {
	Numbers []string
	Source  string
}

type SearchQuery struct {
	Number   string
	Country  string
	Region   string
	Provider string
	Limit    int
	Offset   int
}
