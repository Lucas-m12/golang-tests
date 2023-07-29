package address

import "strings"

func AddressType(address string) string {
	availableTypes := []string{
		"Rua", "Avenida", "Estrada", "Rodovia",
	}
	firstWord := strings.Split(address, " ")[0]
	isValidAddressType := false
	for _, addressType := range availableTypes {
		if strings.ToLower(firstWord) == strings.ToLower(addressType) {
			isValidAddressType = true
		}
	}
	if isValidAddressType == true {
		return firstWord
	} else {
		return "Ivalid address type"
	}
}
