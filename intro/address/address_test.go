package address

import "testing"

func TestAddressType(test *testing.T) {
	addressToTest := "Avenida Paulista"
	expectedAddressType := "Avenida"
	receivedAddressType := AddressType(addressToTest)
	if receivedAddressType != expectedAddressType {
		test.Error("The received type has an error")
	}
}
