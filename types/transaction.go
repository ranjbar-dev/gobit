package types

type Transaction struct {
	// ID
	ID string `json:"id"`
	// Block Height
	BlockHeight int64 `json:"block_height"`
}
