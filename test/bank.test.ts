import expect from 'expect'
import ganache from 'ganache-cli'
import { AbiItem } from 'web3-utils'
import { beforeEach, describe, it } from '@jest/globals'
import { Bank } from '../types/web3-v1-contracts/Bank'
import BankContract from '../build/contracts/Bank.json'
import { Token } from '../types/web3-v1-contracts/Token'
import TokenContract from '../build/contracts/Token.json'
import { testActors } from '../utils/test_util'
import { getWeb3, fromWei, toWei } from '../utils/util'

const ganacheOptions = {}
const web3 = getWeb3(ganache.provider(ganacheOptions))

let bank: Bank
let token: Token

beforeEach(async () => {
    const actors = await testActors(web3)

    // deploy token contract
    const tokenAbiItems: AbiItem[] = TokenContract.abi as AbiItem[]
    const tokenByteCode: string = TokenContract.bytecode
    token = new web3.eth.Contract(tokenAbiItems) as any as Token
    token = await token.deploy({ data: tokenByteCode, arguments: [] }).send(actors.ownerTx) as any as Token

    // deploy bank contract
    const bankAbiItems: AbiItem[] = BankContract.abi as AbiItem[]
    const bankByteCode: string = BankContract.bytecode
    bank = new web3.eth.Contract(bankAbiItems) as any as Bank
    bank = await bank.deploy({ data: bankByteCode, arguments: [token.options.address] }).send(actors.ownerTx) as any as Bank

    // mint 1000 token for user
    const amount = toWei('1000')
    await token.methods.mint(actors.user1Addr, amount).send(actors.ownerTx)
})

describe('Test deposit and withdraw function', () => {
    it('check total supply', async () => {
        const total = await token.methods.totalSupply().call()
        expect(fromWei(total)).toEqual('1000')
    })
})