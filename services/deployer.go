package services

import (
	"fmt"

	"github.com/thanapongsj1996/untrustedbank/bindings/bank"
	"github.com/thanapongsj1996/untrustedbank/bindings/token"
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
		fmt.Errorf("error while getting client : ", err.Error())
		return err
	}

	tokenAddr, tokenTx, _, err := token.DeployToken(utils.MySendOpt(client, network), client)
	if err != nil {
		return err
	}

	bankAddr, bankTx, _, err := bank.DeployBank(utils.MySendOpt(client, network), client, tokenAddr)
	if err != nil {
		return err
	}

	fmt.Printf("TOKEN_TX=%s\n", tokenTx.Hash().Hex())
	fmt.Printf("BANK_TX=%s\n", bankTx.Hash().Hex())
	fmt.Printf("-------------------\n")
	fmt.Printf("TOKEN_ADDR=%s\n", tokenAddr.String())
	fmt.Printf("BANK_ADDR=%s\n", bankAddr.String())

	return nil
}
