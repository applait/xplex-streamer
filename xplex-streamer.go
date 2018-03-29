package main

import (
	"flag"
	"log"
	"os"

	"github.com/applait/xplex-streamer/common"
	"github.com/applait/xplex-streamer/rtmp"
)

func run(conf common.JSONConfig) {
	rtmp.Start(conf.RTMPPort)
}

func main() {
	confPath := flag.String("conf", "config.json", "Path to configuration JSON file relative to current directory.")
	createConfig := flag.Bool("createConfig", false, "Set this flag to generate a dummy config file and exit")

	flag.Parse()

	if *createConfig {
		_, err := common.CreateConfig(*confPath)
		if err != nil {
			log.Fatalf("Unable to generate config file. Reason: %s\n", err)
		}
		log.Printf("Created config file at %s\n", *confPath)
		os.Exit(0)
	}

	conf, err := common.ParseConfig(*confPath)
	if err != nil {
		log.Fatalf("Error loading config. Reason: %s", err)
	}

	// Start
	run(conf)
}
