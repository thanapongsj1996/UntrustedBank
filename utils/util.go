package utils

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/thanapongsj1996/untrustedbank/consts"
)

func MySendOpt(client *ethclient.Client, network string) *bind.TransactOpts {
	myWallet, err := MyWallet(network)
	if err != nil {
		return nil
	}

	path := hdwallet.MustParseDerivationPath(MyAccountPath(network))
	myAccount, err := myWallet.Derive(path, false)
	if err != nil {
		return nil
	}
	privateKey, err := myWallet.PrivateKey(myAccount)
	if err != nil {
		return nil
	}

	nonce, err := client.PendingNonceAt(context.Background(), myAccount.Address)
	if err != nil {
		return nil
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // wei
	if network == consts.GanacheCLI {
		auth.GasLimit = uint64(6721975)
	} else {
		auth.GasLimit = uint64(8000000)
	}

	gasPrice, _ := client.SuggestGasPrice(context.Background())
	auth.GasPrice = gasPrice

	return auth
}

func MyAccountPath(network string) string {
	return fmt.Sprintf("m/44'/60'/0'/0/%d", 0)
}

func MyWallet(network string) (*hdwallet.Wallet, error) {
	mnemonic := ""
	if network == consts.GanacheCLI {
		mnemonic = os.Getenv("SEED_GANACHE")
	} else if network == consts.BSCTest {
		mnemonic = os.Getenv("SEED_BSCTEST")
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		fmt.Printf("error while getting wallet : %s\n", err.Error())
		return nil, err
	}

	return wallet, nil
}
