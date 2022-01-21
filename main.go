package main

import (
	"github.com/thanapongsj1996/untrustedbank/consts"
	"github.com/thanapongsj1996/untrustedbank/services"
)

func main() {
	config := consts.NewConfig()

	deployer := services.NewDeployer(config)
	deployer.Deploy()
}
