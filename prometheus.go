package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var prometheusPrefix = "pi_temp_"

var (
	uptimeSeconds = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: prometheusPrefix + "uptime_seconds",
			Help: "Uptime of the server in seconds",
		},
		[]string{"hostname"},
	)
	cpuTempCelsius = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: prometheusPrefix + "temp_celsius",
			Help: "CPU temperature in degrees Celsius",
		},
		[]string{"hostname"},
	)
)
