package services

import (
	"fmt"

	bank "github.com/thanapongsj1996/untrustedbank/bindings/bank"
	jerryToken "github.com/thanapongsj1996/untrustedbank/bindings/jerry-token"
	pool "github.com/thanapongsj1996/untrustedbank/bindings/pool"
	tomToken "github.com/thanapongsj1996/untrustedbank/bindings/tom-token"
	"github.com/thanapongsj1996/untrustedbank/consts"
	"github.com/thanapongsj1996/untrustedbank/utils"
)

type Deployer struct {
	config consts.IConfig
}

func NewDeployer(config consts.IConfig) *Deployer {
	return &Deployer{config: config}
}

func (svc *Deployer) Deploy() error {
	fmt.Printf("deploying ... \n")

	config := svc.config
	network := config.Network()

	client, err := consts.GetClient(network)
	if err != nil {
		fmt.Printf("error while getting client : %s\n", err.Error())
		return err
	}

	tomTokenAddr, tomTokenTx, _, err := tomToken.DeployTomtoken(utils.MySendOpt(client, network), client)
	if err != nil {
		return err
	}
	jerryTokenAddr, jerryTokenTx, _, err := jerryToken.DeployJerrytoken(utils.MySendOpt(client, network), client)
	if err != nil {
		return err
	}
	bankAddr, bankTx, _, err := bank.DeployBank(utils.MySendOpt(client, network), client, tomTokenAddr)
	if err != nil {
		return err
	}
	poolAddr, poolTx, _, err := pool.DeployPool(utils.MySendOpt(client, network), client, tomTokenAddr, jerryTokenAddr)
	if err != nil {
		return err
	}

	fmt.Printf("TOM_TOKEN_TX=%s\n", tomTokenTx.Hash().Hex())
	fmt.Printf("JERRY_TOKEN_TX=%s\n", jerryTokenTx.Hash().Hex())
	fmt.Printf("BANK_TX=%s\n", bankTx.Hash().Hex())
	fmt.Printf("POOL_TX=%s\n", poolTx.Hash().Hex())
	fmt.Printf("-------------------\n")
	fmt.Printf("TOM_TOKEN_ADDR=%s\n", tomTokenAddr.String())
	fmt.Printf("JERRY_TOKEN_ADDR=%s\n", jerryTokenAddr.String())
	fmt.Printf("BANK_ADDR=%s\n", bankAddr.String())
	fmt.Printf("POOL_ADDR=%s\n", poolAddr.String())

	return nil
}
