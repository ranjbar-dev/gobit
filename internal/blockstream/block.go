package blockstream

import (
	"errors"
	"strconv"
)

func (b *BlockStream) GetCurrentBlockHeight() (int64, error) {

	resp, err := b.c.R().Get(b.host + "/blocks/tip/height")
	if err != nil {

		return 0, err
	}
	if resp.StatusCode() != 200 {

		return 0, errors.New("[blockstream] GetCurrentBlockHeight: " + resp.String())
	}

	return strconv.ParseInt(resp.String(), 10, 64)
}

func (b *BlockStream) GetBlockHashByHeight(height int64) (string, error) {

	resp, err := b.c.R().Get(b.host + "/block-height/" + strconv.FormatInt(height, 10))
	if err != nil {

		return "", err
	}
	if resp.StatusCode() != 200 {

		return "", errors.New("[blockstream] GetBlockHashByHeight: " + resp.String())
	}

	return resp.String(), nil
}
