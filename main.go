package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/aristidesneto/exporter-backup-mysql/config"
	"github.com/aristidesneto/exporter-backup-mysql/metrics"
	"github.com/aristidesneto/exporter-backup-mysql/parser"

	"github.com/prometheus/client_golang/prometheus"
)

var reg *prometheus.Registry

func init() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error finding executable path: %s", err)
	}
	exeDir := filepath.Dir(exePath)
	configDir := filepath.Join(exeDir, "config")
	config.Configuration(configDir)

	// Prometheus metrics
	reg = prometheus.NewRegistry()
	log.Println("Registry initialized")
	metrics.NewMetrics(reg)
	if metrics.M == nil {
		log.Fatal("Metrics not initialized")
	} else {
		log.Println("Metrics initialized successfully")
	}
}


func main()  {
	var logPath string
	flag.StringVar(&logPath, "logpath", "backup.log", "Informe o caminho do arquivo de backup de log")
	flag.Parse()

	if reg == nil {
		log.Fatal("Registry is nil before setting up HTTP handler")
	}

	// Loading backup file
	parser.LoadFile(logPath)

	// serverPort := viper.GetString("server.port")
	// http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	// log.Printf("Server is running on port %s", serverPort)
	// err := http.ListenAndServe(":" + serverPort, nil)
	// if err != nil {
	// 	log.Fatalf("Failed to start HTTP server: %v", err)
	// }
}