package blockstream

import "github.com/go-resty/resty/v2"

type BlockStream struct {
	c    *resty.Client
	host string
}

func NewBlockStream() *BlockStream {

	return &BlockStream{
		c:    resty.New(),
		host: "https://blockstream.info/api",
	}
}
