package types

type Actor struct {
	ChainID string `json:"chain"` // this is empty unless it comes from a different chain
	App     string `json:"app"`   // the app that the actor belongs to
	Address string `json:"addr"`  // arbitrary app-specific unique id
}

type Coin struct {
	Denom  string `json:"denom"`
	Amount int64  `json:"amount"`
}

type Coins []Coin

type Tx struct {
	TxInner "json:\"unwrap\""
}

type TxInner interface {
	Wrap() Tx

	// ValidateBasic should be a stateless check and just verify that the
	// tx is properly formated (required strings not blank, signatures exist, etc.)
	// this can also be run on the client-side for better debugging before posting a tx
	ValidateBasic() error
}

type RequestSign struct {
	Name     string `json:"name,omitempty" validate:"required,min=3,printascii"`
	Password string `json:"password,omitempty" validate:"required,min=10"`

	Tx map[string]interface{} `json:"tx" validate:"required"`
}

type SendInput struct {
	Fees     *Coin  `json:"fees"`
	Multi    bool   `json:"multi,omitempty"`
	Sequence uint32 `json:"sequence"`

	To     *Actor `json:"to"`
	From   *Actor `json:"from"`
	Amount Coins  `json:"amount"`
}
