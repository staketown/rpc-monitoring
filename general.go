package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"net/http"
	"strconv"
	"time"
)

var ()

func GeneralHandler(w http.ResponseWriter, r *http.Request) {
	requestStart := time.Now()

	rpcStatusGauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:        "rpc_status",
			Help:        "Current status of rpc getting from synced_up: true/false",
			ConstLabels: ConstLabels,
		},
		[]string{
			"chain_id",
			"rpc",
			"error",
		},
	)

	rpcSyncLatestBlockGauge := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:        "rpc_sync_blocks",
			Help:        "Current latest block on rpc",
			ConstLabels: ConstLabels,
		},
		[]string{
			"chain_id",
			"rpc",
			"error",
		},
	)

	sublogger := log.With().
		Str("request-id", uuid.New().String()).
		Logger()

	rpc := r.URL.Query().Get("rpc")

	registry := prometheus.NewRegistry()
	registry.MustRegister(rpcStatusGauge)
	registry.MustRegister(rpcSyncLatestBlockGauge)

	// doing this not in goroutine as we'll need params from oracle params response for calculation
	sublogger.Debug().Msg("Started querying rpc data")
	queryStart := time.Now()

	data, err := getRpcData(rpc)

	sublogger.Debug().
		Float64("request-time", time.Since(queryStart).Seconds()).
		Msg("Finished querying rpc data")

	if err != nil {
		rpcStatusGauge.With(prometheus.Labels{
			"chain_id": "",
			"rpc":      rpc,
			"error":    err.Error(),
		}).Set(0)

		rpcSyncLatestBlockGauge.With(prometheus.Labels{
			"chain_id": "",
			"rpc":      rpc,
			"error":    err.Error(),
		}).Set(0)
	} else {
		// caught up by default
		var catchUp float64 = 1
		if !data.Result.SyncInfo.CatchingUp {
			catchUp = 0
		}

		height, _ := strconv.ParseFloat(data.Result.SyncInfo.LatestBlockHeight, 64)

		rpcStatusGauge.With(prometheus.Labels{
			"chain_id": data.Result.NodeInfo.Network,
			"rpc":      rpc,
			"error":    "",
		}).Set(catchUp)

		rpcSyncLatestBlockGauge.With(prometheus.Labels{
			"chain_id": data.Result.NodeInfo.Network,
			"rpc":      rpc,
			"error":    "",
		}).Set(height)
	}

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
	sublogger.Info().
		Str("method", "GET").
		Str("endpoint", "/metrics/general?rpc="+rpc).
		Float64("request-time", time.Since(requestStart).Seconds()).
		Msg("Request processed")
}

func getRpcData(rpc string) (*RpcData, error) {
	response, err := http.Get(rpc + "/status")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// close response body
	defer response.Body.Close()

	// read response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response.StatusCode >= 300 || response.StatusCode < 200 {
		fmt.Println(string(body))
		return nil, errors.New(string(body))
	}

	var rpcData *RpcData
	err = json.Unmarshal(body, &rpcData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return rpcData, nil
}
