package currency

type CurrencyOptions struct {
	ContractAddress string
	Decimal         int
}

type Currency struct {
	Symbol  string
	Options CurrencyOptions
}

func NewCurrency(symbol string, options CurrencyOptions) *Currency {
	return &Currency{}
}
