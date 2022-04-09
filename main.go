package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var hostname = getHostname()

func main() {
	mux := http.NewServeMux()

	go uptime()
	go cpuTempMeasurement()

	mux.Handle("/metrics", promhttp.Handler())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("starting server on %s", srv.Addr)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

func uptime() {
	for {
		uptimeSeconds.WithLabelValues(hostname).Inc()
		time.Sleep(1 * time.Second)
	}
}

func cpuTempMeasurement() {
	for {
		content, err := ioutil.ReadFile(os.Getenv("FILE"))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(content))
		cpuTempCelsius.WithLabelValues(hostname).Set(getCpuTemp())
		time.Sleep(15 * time.Second)
	}
}

func getCpuTemp() float64 {
	cpuTempRaw := getCpuTempFromFile()
	cpuTemp := parseRawCpuTemp(cpuTempRaw)
	return cpuTemp
}

func parseRawCpuTemp(cpuTempRaw string) float64 {
	cpuTemp1 := strings.TrimPrefix(cpuTempRaw, "temp=")
	cpuTemp2 := strings.TrimSuffix(cpuTemp1, "'C\n")
	cpuTempString, err := strconv.ParseFloat(cpuTemp2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return cpuTempString
}

func getCpuTempFromFile() string {
	content, err := ioutil.ReadFile(os.Getenv("FILE"))
	if err != nil {
		fmt.Println(err)
	}
	return string(content)
}

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hostname: " + hostname)
	return hostname
}
