package main

import (
	"github.com/thanapongsj1996/untrustedbank/consts"
	"github.com/thanapongsj1996/untrustedbank/services"
)

func main() {
	config := consts.NewConfig()
	if err := services.NewDeployer(config).Deploy(); err != nil {
		panic("error while deploying..")
	}
}
