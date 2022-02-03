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
	if [ ! -d bindings/tom-token ]; then mkdir bindings/tom-token ; fi
	./abigen --abi=abigenBindings/abi/TomToken.abi --bin=abigenBindings/bin/TomToken.bin --pkg tomtoken --out bindings/tom-token/bindings.go
	if [ ! -d bindings/jerry-token ]; then mkdir bindings/jerry-token ; fi
	./abigen --abi=abigenBindings/abi/JerryToken.abi --bin=abigenBindings/bin/JerryToken.bin --pkg jerrytoken --out bindings/jerry-token/bindings.go
	if [ ! -d bindings/lp-token ]; then mkdir bindings/lp-token ; fi
	./abigen --abi=abigenBindings/abi/LPToken.abi --bin=abigenBindings/bin/LPToken.bin --pkg lptoken --out bindings/lp-token/bindings.go
	if [ ! -d bindings/bank ]; then mkdir bindings/bank ; fi
	./abigen --abi=abigenBindings/abi/Bank.abi --bin=abigenBindings/bin/Bank.bin --pkg bank --out bindings/bank/bindings.go
	if [ ! -d bindings/pool ]; then mkdir bindings/pool ; fi
	./abigen --abi=abigenBindings/abi/Pool.abi --bin=abigenBindings/bin/Pool.bin --pkg pool --out bindings/pool/bindings.go
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
