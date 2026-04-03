package business

import "strings"

var ruDEF = map[string]struct {
	Region   string
	Provider string
}{
	"901": {"Москва", "Beeline"},
	"902": {"Москва", "Tele2"},
	"903": {"Москва", "Beeline"},
	"904": {"Северо-Запад", "Tele2"},
	"905": {"Москва", "Beeline"},
	"906": {"Москва", "Beeline"},
	"910": {"Москва", "МТС"},
	"911": {"Северо-Запад", "МТС"},
	"912": {"Урал", "МТС"},
	"913": {"Сибирь", "МТС"},
	"914": {"Дальний Восток", "МТС"},
	"915": {"Москва", "МТС"},
	"916": {"Москва", "МТС"},
	"917": {"Москва", "МТС"},
	"918": {"Юг", "МТС"},
	"919": {"Москва", "МТС"},
}

func DetectRUProvider(phone string) (string, string) {
	phone = strings.TrimPrefix(phone, "+")
	if len(phone) < 11 || phone[0] != '7' {
		return "UNKNOWN", "UNKNOWN"
	}

	if len(phone) < 4 {
		return "UNKNOWN", "UNKNOWN"
	}

	def := phone[1:4]

	if v, ok := ruDEF[def]; ok {
		return v.Region, v.Provider
	}

	return "UNKNOWN", "UNKNOWN"
}
