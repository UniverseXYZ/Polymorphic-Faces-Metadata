package main

import (
	"github.com/polymorph-metadata/app/interface/dlt/ethereum"
	log "github.com/sirupsen/logrus"
	"os"
)

func connectToNodes() (*ethereum.EthereumClient, *ethereum.EthereumClient) {

	nodeURL := os.Getenv("NODE_URL")

	nodeURLPolygon := os.Getenv("NODE_URL_POLYGON")

	client, err := ethereum.NewEthereumClient(nodeURL)

	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Successfully connected to ethereum client")

	clientPolygon, err := ethereum.NewEthereumClient(nodeURLPolygon)

	if err != nil {
		log.Fatal(err)
	}

	log.Infoln("Successfully connected to polygon client")

	return client, clientPolygon
}
