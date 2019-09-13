package fio

import (
	"github.com/trustwallet/blockatlas"
	"github.com/trustwallet/blockatlas/coin"
)

type Platform struct {
	client Client
}

func (p *Platform) Init() error {
	p.client = InitClient()
	return nil
}

func (p *Platform) Coin() coin.Coin {
	return coin.Coins[coin.FIO]
}

func (p *Platform) GetTxsByAddress(address string) (page blockatlas.TxPage, err error) {
	return page, err
}
