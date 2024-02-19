# currency-test

## Usage (example)

`main.go`
```go
// Инициализируем по курсу из файла
currency.LoadRate("exchange.rate.txt")

// Создаем 2 кошелька
m1 := currency.Money{1, currency.USD}
m2 := m1.As(currency.RUB)

// Складываем деньги из 2 кошельков представляя в виде рублей и в виде долларов по курсу из r2
m3 := m1.Add(m2).As(currency.RUB)
m4 := m1.Add(m2).As(currency.USD)

// Проверим также вычитание
m5 := m4.Sub(m2).As(currency.RUB)

// Выводим получившийся результат в консоль
fmt.Println(m3)
fmt.Println(m5)
```

## Run
`go run main.go`
