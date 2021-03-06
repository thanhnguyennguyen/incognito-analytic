package models

import (
	"time"
)

type ShardBlock struct {
	ID            string `db:"id"`
	BlockHash     string `db:"block_hash"`
	BlockVersion  int    `db:"block_version"`
	BlockHeight   uint64 `db:"block_height"`
	BlockProducer string `db:"block_producer"`
	Epoch         uint64 `db:"epoch"`
	Round         int    `db:"round"`

	CreatedTime       time.Time `db:"created_time"`
	Data              string    `db:"data"`
	CountTx           int       `db:"count_tx"`
	ShardID           int       `db:"shard_id"`
	PreBlock          string    `db:"pre_block"`
	NextBlock         string    `db:"next_block"`
	ListHashTx        string    `db:"list_hash_tx"`
	BeaconBlockHeight uint64    `db:"beacon_block_height"`
}
