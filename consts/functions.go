package consts

import (
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func GetClient(network string) (*ethclient.Client, error) {
	networkURL := ""
	if network == GanacheCLI {
		networkURL = os.Getenv("LOCAL_NETWORK")
	} else if network == BSCTest {
		networkURL = os.Getenv("BSC_NETWORK")
	}

	client, err := ethclient.Dial(networkURL)
	if err != nil {
		return nil, err
	}

	return client, nil
}
