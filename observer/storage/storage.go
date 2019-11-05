package storage

import (
	"github.com/trustwallet/blockatlas/pkg/storage/redis"
)

type Storage struct {
	redis.Redis
	blockHeights BlockMap
}

func New() *Storage {
	s := new(Storage)
	s.blockHeights.heights = make(map[uint]int64)
	return s
}

type Tracker interface {
	GetBlockNumber(coin uint) (int64, error)
	SetBlockNumber(coin uint, num int64) error
}

type Addresses interface {
	Lookup(coin uint, addresses []string) ([]Subscription, error)
	AddSubscriptions(subscriptions []Subscription)
	DeleteSubscriptions(subscriptions []Subscription)
	GetXpubFromAddress(coin uint, address string) (xpub string, addresses []string, err error)
	GetXpub(coin uint, xpub string) ([]string, error)
}
