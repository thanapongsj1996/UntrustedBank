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
    await token.methods.mint(amount).send(actors.user1Tx)
    await token.methods.mint(amount).send(actors.user2Tx)
})

describe('Test deposit and withdraw function', () => {
    it('correct total supply', async () => {
        const total = await token.methods.totalSupply().call()
        expect(fromWei(total)).toEqual('2000')
    })
    it('can deposit and withdraw', async () => {
        const actors = await testActors(web3)
        const bankAddr = bank.options.address

        // user deposit 200 token, check bank balance and user balance
        let amount = toWei('200')
        await token.methods.approve(bankAddr, amount).send(actors.user1Tx)
        await bank.methods.deposit(amount).send(actors.user1Tx)
        let balBank = await token.methods.balanceOf(bankAddr).call()
        let balUser = await token.methods.balanceOf(actors.user1Addr).call()
        let balUserInBank = await bank.methods.checkUserBalance().call(actors.user1Tx)

        expect(fromWei(balBank)).toEqual('200')
        expect(fromWei(balUserInBank)).toEqual('200')
        expect(fromWei(balUser)).toEqual('800')

        amount = toWei('50')
        await bank.methods.withdraw(amount).send(actors.user1Tx)
        balBank = await token.methods.balanceOf(bankAddr).call()
        balUser = await token.methods.balanceOf(actors.user1Addr).call()

        expect(fromWei(balBank)).toEqual('150')
        expect(fromWei(balUser)).toEqual('850')
    })
    it('can tranfer token from user to the other user', async () => {
        const actors = await testActors(web3)

        const amount = toWei('50')
        await token.methods.transfer(actors.user2Addr, amount).send(actors.user1Tx)
        const balUser1 = await token.methods.balanceOf(actors.user1Addr).call()
        const balUser2 = await token.methods.balanceOf(actors.user2Addr).call()

        expect(fromWei(balUser1)).toEqual('950')
        expect(fromWei(balUser2)).toEqual('1050')
    })
})