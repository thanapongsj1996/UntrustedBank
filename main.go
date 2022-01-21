package main

import (
	"github.com/joho/godotenv"
	"github.com/thanapongsj1996/untrustedbank/consts"
	"github.com/thanapongsj1996/untrustedbank/services"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("error while getting environment variables..")
	}

	config := consts.NewConfig()
	if err := services.NewDeployer(config).Deploy(); err != nil {
		panic("error while deploying..")
	}
}
