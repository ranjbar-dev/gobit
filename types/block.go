package types

type Block struct {
	// Height
	Height int64 `json:"height"`
	// Hash
	Hash string `json:"hash"`
	// Timestamp
	Timestamp int64 `json:"timestamp"`
}
