package validator

import (
	"regexp"
)

var phoneRegex = regexp.MustCompile(`^(?:\+55\s?)?(?:\(?\d{2}\)?[\s-]?)(?:9\d{4}|[2-8]\d{3})[\s-]?\d{4}$`)

func Phone(phone string) bool {
	return phoneRegex.MatchString(phone)
}
