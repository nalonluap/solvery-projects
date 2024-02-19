package currency

import (
	"bufio"
	"fmt"
	"github.com/shopspring/decimal"
	"os"
	"strconv"
	"strings"
)

const (
	USD = "USD"
	RUB = "RUB"
)

type (
	Currency string
	Quantity int
)

type Money struct {
	Quantity Quantity
	Currency Currency
}

type RateCQType map[Currency]Quantity

var (
	defaultBaseCurrency = USD
	defaultRateCQ       = RateCQType{
		USD: 1,
		RUB: 90,
	}
	baseCurrency = defaultBaseCurrency
	RateCQ       = defaultRateCQ
)

func LoadRateFromMap(rate RateCQType) {
	RateCQ = rate
}

func LoadRateFromFile(filePath string) error {
	var lines []string

	file, err := os.Open(filePath)
	if err != nil {
		lines, err = nil, fmt.Errorf("Ошибка при открытии файла: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err != nil {
		return fmt.Errorf("Ошибка чтения из файла: %w", err)
	}

	for _, line := range lines {
		var rateValue int
		rateStr := strings.Split(line, " ")
		rateValue, err = strconv.Atoi(rateStr[1])
		if err != nil {
			return fmt.Errorf("Введен некорректный формат числа: %w", err)
		}

		RateCQ[Currency(rateStr[0])] = Quantity(rateValue)
	}

	return nil
}

func (m1 Money) Add(m2 Money) Money {
	return Money{m1.Quantity + m2.As(m1.Currency).Quantity, m1.Currency}
}

func (m1 Money) Sub(m2 Money) Money {
	return Money{m1.Quantity - m2.As(m1.Currency).Quantity, m1.Currency}
}

func (m1 Money) As(c Currency) Money {
	var result decimal.Decimal
	if m1.Currency == c {
		result = decimal.NewFromInt(int64(m1.Quantity))
	} else {
		if m1.Currency == Currency(baseCurrency) {
			result = decimal.NewFromInt(int64(m1.Quantity)).Mul(decimal.NewFromInt(int64(RateCQ[c])))
		} else {
			result = decimal.NewFromInt(int64(m1.Quantity)).Div(decimal.NewFromInt(int64(RateCQ[m1.Currency])))
		}
	}

	m := Money{Quantity(result.IntPart()), c}
	return m
}

func (m1 Money) String() string {
	return fmt.Sprint(m1.Quantity, " ", m1.Currency)
}
