package tests

import (
	"currency/pkg/currency"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMoney_Add(t *testing.T) {
	// Arrange
	money1 := currency.Money{Quantity: 100, Currency: currency.USD}
	money2 := currency.Money{Quantity: 200, Currency: currency.USD}

	// Act
	result := money1.Add(money2)

	// Assert
	assert.Equal(t, currency.Quantity(300), result.Quantity, "they should be equal")
}

func TestMoney_Sub(t *testing.T) {
	// Arrange
	money1 := currency.Money{Quantity: 200, Currency: currency.USD}
	money2 := currency.Money{Quantity: 100, Currency: currency.USD}

	// Act
	result := money1.Sub(money2)

	// Assert
	assert.Equal(t, currency.Quantity(100), result.Quantity, "they should be equal")
}

func TestMoney_As(t *testing.T) {
	// Arrange
	money := currency.Money{Quantity: 1, Currency: currency.USD}
	currency.LoadRateFromMap(currency.RateCQType{
		currency.USD: 1,
		currency.RUB: 90,
	})

	// Act
	result := money.As(currency.RUB)

	// Assert
	assert.Equal(t, currency.Quantity(90), result.Quantity, "conversion should match the rate")
}

func TestMoney_String(t *testing.T) {
	// Arrange
	money := currency.Money{Quantity: 100, Currency: currency.USD}

	// Act
	result := money.String()

	// Assert
	assert.Equal(t, "100 USD", result, "string representation should match")
}

func TestLoadRateFromFile(t *testing.T) {
	// Arrange
	filePath := "./exchange.rate.txt"
	content := []byte("USD 1\nRUB 75")
	err := os.WriteFile(filePath, content, 0644)
	require.NoError(t, err)
	defer os.Remove(filePath)

	// Act
	err = currency.LoadRateFromFile(filePath)

	// Assert
	require.NoError(t, err)
	assert.Equal(t, currency.Quantity(1), currency.RateCQ[currency.USD], "USD rate should be 1")
	assert.Equal(t, currency.Quantity(75), currency.RateCQ[currency.RUB], "RUB rate should be 75")
}

func TestLoadRateFromFile_Error(t *testing.T) {
	// Arrange
	filePath := "./non_existent_file.txt"

	// Act
	err := currency.LoadRateFromFile(filePath)

	// Assert
	assert.Error(t, err, "should return an error for non-existent file")
}
