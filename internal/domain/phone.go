package domain

import "time"

type ImportResult struct {
	Accepted   int
	Duplicates int
	Errors     int
}

type SearchResult struct {
	Total   int
	Numbers []PhoneNumber
}

type PhoneNumber struct {
	ID        int64
	Number    string
	Country   string
	Region    string
	Provider  string
	Source    string
	CreatedAt time.Time
}
