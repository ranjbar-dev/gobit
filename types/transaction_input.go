package types

type TransactionInput struct {
	// ID
	ID string `json:"id"`
	// Transaction ID
	TransactionID string `json:"transaction_id"`
	// Hash
	Hash string `json:"hash"`
	// Index
	Index int64 `json:"index"`
	// Address
	Address string `json:"address"`
	// Value
	Value int64 `json:"value"`
}
