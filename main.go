package main

import (
	"github.com/joho/godotenv"
	"github.com/thanapongsj1996/untrustedbank/consts"
	"github.com/thanapongsj1996/untrustedbank/services"
)

func main() {
	godotenv.Load(".env")

	config := consts.NewConfig()
	if err := services.NewDeployer(config).Deploy(); err != nil {
		panic("error while deploying..")
	}
}
