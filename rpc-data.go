package main

import "time"

type RpcData struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  struct {
		NodeInfo struct {
			Network string `json:"network"`
		} `json:"node_info"`
		SyncInfo struct {
			LatestBlockHash   string    `json:"latest_block_hash"`
			LatestAppHash     string    `json:"latest_app_hash"`
			LatestBlockHeight string    `json:"latest_block_height"`
			LatestBlockTime   time.Time `json:"latest_block_time"`
			CatchingUp        bool      `json:"catching_up"`
		} `json:"sync_info"`
	} `json:"result"`
}
