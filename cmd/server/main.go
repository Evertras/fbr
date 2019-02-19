package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/Evertras/fbr/lib/server"
	metrics "github.com/armon/go-metrics"
	"github.com/armon/go-metrics/prometheus"
	"github.com/spf13/pflag"
)

type metricsConfig struct {
	statsdAddr           string
	prometheusListenAddr string
}

type config struct {
	metrics metricsConfig
	server  server.Config
}

func main() {
	// TODO: handle cancel via interrupts instead of just using Background which does nothing.
	// This should include draining to wait for graceful shutdown.
	ctx := context.Background()

	cfg := getConfig()

	initMetrics(cfg.metrics)

	s := server.New(ctx, cfg.server)

	log.Println("Starting")

	log.Fatal(s.Listen("0.0.0.0:8000"))
}

func getConfig() config {
	//tickRate := pflag.IntP("tick-rate", "t", 5, "How many ticks per second to update clients.")
	devMode := pflag.BoolP("dev", "d", false, "If set, serve files from disk rather than in-memory so changes can be served without a restart.")
	statsd := pflag.StringP("statsd", "", "", "The address to send statsd metrics to.  Example: localhost:8125")
	prometheus := pflag.StringP("prometheus", "", "", "The address to listen on for Prometheus scrapes.  Example: :9090")

	pflag.Parse()

	cfg := server.Config{
		ReadStaticFilesPerRequest: *devMode,
	}

	return config{
		metrics: metricsConfig{
			statsdAddr:           *statsd,
			prometheusListenAddr: *prometheus,
		},
		server: cfg,
	}
}

func initMetrics(cfg metricsConfig) {
	var err error

	var sinks metrics.FanoutSink = make([]metrics.MetricSink, 0)

	// Always include an in-memory sink that we can poke with SIGUSR1 for metric dumps to stdout
	inmem := metrics.NewInmemSink(10*time.Second, time.Minute)
	metrics.DefaultInmemSignal(inmem)

	sinks = append(sinks, inmem)

	if cfg.statsdAddr != "" {
		statsd, err := metrics.NewStatsdSink(cfg.statsdAddr)
		if err != nil {
			log.Fatalf("Unable to create statsd sink: %v", err)
		}

		sinks = append(sinks, statsd)
	}

	if cfg.prometheusListenAddr != "" {
		promSink, err := prometheus.NewPrometheusSink()

		if err != nil {
			log.Fatalf("Unable to create Prometheus sink: %v", err)
		}

		sinks = append(sinks, promSink)

		go func() {
			handler := promhttp.Handler()
			promServer := http.Server{
				Addr:    cfg.prometheusListenAddr,
				Handler: handler,
			}

			promServer.ListenAndServe()
		}()
	}

	_, err = metrics.NewGlobal(
		&metrics.Config{
			ServiceName:        "gopong",
			EnableServiceLabel: true,

			HostName:            "tmp",
			EnableHostname:      false,
			EnableHostnameLabel: false,

			EnableTypePrefix: false,

			TimerGranularity:     time.Millisecond,
			ProfileInterval:      time.Second,
			EnableRuntimeMetrics: true,

			FilterDefault: true,
		},
		sinks,
	)

	if err != nil {
		log.Fatalf("Unable to create global metrics: %v", err)
	}
}
