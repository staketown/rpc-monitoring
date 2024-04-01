package main

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var (
	ListenAddress string

	LogLevel string

	ConstLabels map[string]string
)

var log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()

var rootCmd = &cobra.Command{
	Use:  "rpc-exporter",
	Long: "Scrape the data about the rpc.",
	Run:  Execute,
}

func Execute(_ *cobra.Command, _ []string) {
	logLevel, err := zerolog.ParseLevel(LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not parse log level")
	}

	zerolog.SetGlobalLevel(logLevel)

	log.Info().
		Str("--listen-address", ListenAddress).
		Str("--log-level", LogLevel).
		Msg("Started with following parameters")

	http.HandleFunc("/metrics/general", func(w http.ResponseWriter, r *http.Request) {
		GeneralHandler(w, r)
	})

	log.Info().Str("address", ListenAddress).Msg("Listening")
	err = http.ListenAndServe(ListenAddress, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("Could not start application")
	}
}

func main() {
	rootCmd.PersistentFlags().StringVar(&ListenAddress, "listen-address", ":9300", "The address this exporter would listen on")
	rootCmd.PersistentFlags().StringVar(&LogLevel, "log-level", "info", "Logging level")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("Could not start application")
	}
}
