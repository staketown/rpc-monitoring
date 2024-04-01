module main

go 1.21

toolchain go1.21.6

require (
	github.com/google/uuid v1.4.0
	github.com/prometheus/client_golang v1.18.0
	github.com/rs/zerolog v1.30.0
	github.com/spf13/cobra v1.8.0
)

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/matttproud/golang_protobuf_extensions/v2 v2.0.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.45.0 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.15.0 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace cosmossdk.io/api => cosmossdk.io/api v0.3.1

replace github.com/cosmos/iavl => github.com/cosmos/iavl v0.20.0

// pin version! 126854af5e6d has issues with the store so that queries fail
replace github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7

replace github.com/linxGnu/grocksdb => github.com/linxGnu/grocksdb v1.8.9
