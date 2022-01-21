package services

import (
	"fmt"

	"github.com/thanapongsj1996/untrustedbank/consts"
)

type Deployer struct {
	config consts.IConfig
}

func NewDeployer(config consts.IConfig) *Deployer {
	return &Deployer{config: config}
}

func (svc *Deployer) Deploy() error {
	fmt.Printf("deploying ... \n")
	return nil
}
