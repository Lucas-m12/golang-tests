package address

import "testing"

type testCase struct {
	insertedAddress string
	expectedReturn  string
}

func TestAddressType(test *testing.T) {
	testCases := []testCase{
		{
			insertedAddress: "Avenida Paulista",
			expectedReturn:  "Avenida",
		},
		{
			insertedAddress: "Rua Paulista",
			expectedReturn:  "Rua",
		},
		{
			insertedAddress: "Rodovia Paulista",
			expectedReturn:  "Rodovia",
		},
		{
			insertedAddress: "Estrada Paulista",
			expectedReturn:  "Estrada",
		},
		{
			insertedAddress: "Praça das Rosas",
			expectedReturn:  "Invalid address type",
		},
	}
	for _, item := range testCases {
		receivedAddressType := AddressType(item.insertedAddress)
		if receivedAddressType != item.expectedReturn {
			test.Error("The received type has an error")
		}
	}
}
