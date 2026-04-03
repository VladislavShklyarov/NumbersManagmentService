package business

import "strings"

var countryByCode = map[string]string{
	"7":  "RU",
	"1":  "US",
	"44": "UK",
	"49": "DE",
	"33": "FR",
	"86": "CN",
}

func DetectCountry(phone string) string {
	phone = strings.TrimPrefix(phone, "+")

	if len(phone) >= 2 {
		if c, ok := countryByCode[phone[:2]]; ok {
			return c
		}
	}

	if len(phone) >= 1 {
		if c, ok := countryByCode[string(phone[0])]; ok {
			return c
		}
	}

	return "UNKNOWN"
}
