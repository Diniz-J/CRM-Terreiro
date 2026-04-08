package validator

import (
	"regexp"
	"strings"
)

var cpfRegex = regexp.MustCompile(`^\d{11}$`)

func CPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	if !cpfRegex.MatchString(cpf) {
		return false
	}

	// rejeita sequências como 00000000000, 11111111111, etc.
	allSame := true
	for i := 1; i < 11; i++ {
		if cpf[i] != cpf[0] {
			allSame = false
			break
		}
	}
	if allSame {
		return false
	}

	// valida 1º dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		sum += int(cpf[i]-'0') * (10 - i)
	}
	resto := sum % 11
	primeiro := 0
	if resto >= 2 {
		primeiro = 11 - resto
	}
	if int(cpf[9]-'0') != primeiro {
		return false
	}

	// valida 2º dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		sum += int(cpf[i]-'0') * (11 - i)
	}
	resto = sum % 11
	segundo := 0
	if resto >= 2 {
		segundo = 11 - resto
	}
	return int(cpf[10]-'0') == segundo
}
