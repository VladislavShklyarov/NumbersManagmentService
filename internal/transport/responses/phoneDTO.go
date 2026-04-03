package responses

type PhoneDTO struct {
	Number   string `json:"number"`
	Country  string `json:"country"`
	Region   string `json:"region"`
	Provider string `json:"provider"`
	Source   string `json:"source"`
}
