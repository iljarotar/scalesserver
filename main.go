package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/iljarotar/scalesserver/server"
	"go.uber.org/zap"
)

func main() {
	var host string
	flag.StringVar(&host, "host", "0.0.0.0", "The host the server binds to")

	var port string
	flag.StringVar(&port, "port", "8080", "The port the server binds to")

	var maxRange int
	flag.IntVar(&maxRange, "max-range", 12, "The maximum input value for the range field")

	var maxNum int
	flag.IntVar(&maxNum, "max-num", 12, "The maximum input value for the notes field")

	flag.Parse()

	l, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("unable to create logger, %v", err)
		os.Exit(1)
	}

	logger := l.Sugar()

	serverConfig := &server.ServerConfig{
		Host:     host,
		Port:     port,
		MaxRange: maxRange,
		MaxNum:   maxNum,
		Logger:   logger,
	}

	scalesServer := server.NewServer(serverConfig)
	log.Fatal(scalesServer.Serve())
}
