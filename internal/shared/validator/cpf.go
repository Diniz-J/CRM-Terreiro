package validator

import (
	"regexp"
	"strings"
)

//^começo // $fim

// \d numero \w letra ou numero \s espaço . qq caracteres

// \D!numero \W!letra/numero \S!espaço

var cpfRegex = regexp.MustCompile(`^(\d{11})$`)

func CPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if len(cpf) != 11 {
		return false
	}

	allSame := true
	for i := 0; i < 11; i++ {
		if cpf[i] != cpf[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	return cpfRegex.MatchString(cpf)
}
