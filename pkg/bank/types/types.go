package types

type Money int64

type Currency string

const (
        TJS Currency = "TJS"
        RUB Currency = "RUB"
        USD Currency = "USD"
)

type PAN string

type Card struct {
        ID int
        PAN PAN
        Balance Money
        MinBalance Money
        Currency Currency
        Color string
        Name string
        Active bool
}

// Category категории
type Category string

type Payment struct{
        ID int
        Amount Money
        Category Category
}

type PaymentSource struct{
    Type string // card
    Number string // number as '5058 xxxx xxxx 8888'
    Balance Money // balance in diram
}