package wallet

type Option func(o *Options)

type Options struct {
	Mnemonic string
	Index    uint32
}

type Wallet struct {
	Options Options
}

func AddressIndex(i uint32) Option {
	return func(o *Options) {
		o.Index = i
	}
}

func Mnemonic(mnemonic string) Option {
	return func(o *Options) {
		o.Mnemonic = mnemonic
	}
}

func NewWallet(opts ...Option) *Wallet {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return &Wallet{
		Options: opt,
	}
}
