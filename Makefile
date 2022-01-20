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
	yarn add web3
	yarn add bn.js
truffle-compile:
	truffle compile
	yarn run typechain --target=web3-v1 'build/contracts/*.json' 
test-contract:
	yarn run jest --testTimeout=60000