yarn-init:
	yarn add -D ganache-cli
	yarn add -D typechain
	yarn add -D @typechain/web3-v1
	yarn add -D jest @types/jest
	yarn add -D @types/web3
	yarn add -D @types/node
	yarn add -D ts-node
	yarn add -D ts-jest
	yarn add -D typescript
	yarn add -D @openzeppelin/contracts@4.4.0
	yarn add -D @chainsafe/truffle-plugin-abigen 
	yarn add web3
	yarn add bn.js
truffle-compile:
	truffle compile
	yarn run typechain --target=web3-v1 'build/contracts/*.json' 
gen-abi-binding:
	if [ ! -d abigenBindings ]; then mkdir abigenBindings ; fi
	truffle run abigen
gen-go-contract-binding:
	if [ ! -d bindings ]; then mkdir bindings ; fi
	if [ ! -d bindings/token ]; then mkdir bindings/token ; fi
	./abigen --abi=abigenBindings/abi/Token.abi --bin=abigenBindings/bin/Token.bin --pkg token --out bindings/token/bindings.go
	if [ ! -d bindings/bank ]; then mkdir bindings/bank ; fi
	./abigen --abi=abigenBindings/abi/Bank.abi --bin=abigenBindings/bin/Bank.bin --pkg bank --out bindings/bank/bindings.go
deploy-contract:
	go run main.go
test-contract:
	yarn run jest --testTimeout=60000
go-init:
	go mod init github.com/thanapongsj1996/untrustedbank
go-get:
	go get
go-build:
	go build .
