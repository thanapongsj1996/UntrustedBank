import expect from 'expect'
import ganache from 'ganache-cli'
import { AbiItem } from 'web3-utils'
import { beforeEach, describe, it } from '@jest/globals'
import { TomToken } from '../types/web3-v1-contracts/TomToken'
import TomTokenContract from '../build/contracts/TomToken.json'
import { JerryToken } from '../types/web3-v1-contracts/JerryToken'
import JerryTokenContract from '../build/contracts/JerryToken.json'
import { Pool } from '../types/web3-v1-contracts/Pool'
import PoolContract from '../build/contracts/Pool.json'
import { testActors } from '../utils/test_util'
import { getWeb3, fromWei, toWei } from '../utils/util'

const ganacheOptions = {}
const web3 = getWeb3(ganache.provider(ganacheOptions))

let tomToken: TomToken
let jerryToken: JerryToken
let pool: Pool

beforeEach(async () => {
    const actors = await testActors(web3)

    // deploy TomToken contract
    const tomTokenAbiItems: AbiItem[] = TomTokenContract.abi as AbiItem[]
    const tomTokenByteCode: string = TomTokenContract.bytecode
    tomToken = new web3.eth.Contract(tomTokenAbiItems) as any as TomToken
    tomToken = await tomToken.deploy({ data: tomTokenByteCode, arguments: [] }).send(actors.ownerTx) as any as TomToken

    // deploy JerryToken contract
    const jerryTokenAbiItems: AbiItem[] = JerryTokenContract.abi as AbiItem[]
    const jerryTokenByteCode: string = JerryTokenContract.bytecode
    jerryToken = new web3.eth.Contract(jerryTokenAbiItems) as any as JerryToken
    jerryToken = await jerryToken.deploy({ data: jerryTokenByteCode, arguments: [] }).send(actors.ownerTx) as any as JerryToken

    const tomTokenAddr = tomToken.options.address
    const jerryTokenAddr = jerryToken.options.address

    // deploy Pool contract
    const poolAbiItems: AbiItem[] = PoolContract.abi as AbiItem[]
    const poolByteCode: string = PoolContract.bytecode
    pool = new web3.eth.Contract(poolAbiItems) as any as Pool
    pool = await pool.deploy({ data: poolByteCode, arguments: [tomTokenAddr, jerryTokenAddr] }).send(actors.ownerTx) as any as Pool
})

describe('the pool', () => {
    it('check lp amount add/remove', async () => {
        const actors = await testActors(web3)
        const poolAddr = pool.options.address

        // tom : jerry = 1 : 2
        const tomAmount = toWei('1000000')
        const jerryAmount = toWei('2000000')

        // mint tom & jerry and add to liquidity pool
        await tomToken.methods.mint(tomAmount).send(actors.ownerTx)
        await jerryToken.methods.mint(jerryAmount).send(actors.ownerTx)
        await tomToken.methods.approve(poolAddr, tomAmount).send(actors.ownerTx)
        await jerryToken.methods.approve(poolAddr, jerryAmount).send(actors.ownerTx)
        await pool.methods.addLiquidity(tomAmount, jerryAmount).send(actors.ownerTx)

        // check LP token for owner
        let lpAmount = await pool.methods.balanceOf(actors.ownerAddr).call()
        expect(fromWei(lpAmount)).toEqual('1000000')

        // remove half of lp token then get tom&jerry
        await pool.methods.removeLiquidity(toWei('500000')).send(actors.ownerTx)
        lpAmount = await pool.methods.balanceOf(actors.ownerAddr).call()
        expect(fromWei(lpAmount)).toEqual('500000')

        const tomReceived = await tomToken.methods.balanceOf(actors.ownerAddr).call()
        const jerryReceived = await jerryToken.methods.balanceOf(actors.ownerAddr).call()
        expect(fromWei(tomReceived)).toEqual('500000')
        expect(fromWei(jerryReceived)).toEqual('1000000')
    })
    it('can swap', async () => {
        const actors = await testActors(web3)
        const poolAddr = pool.options.address

        // tom : jerry = 1 : 2
        const tomAmount = toWei('1000000')
        const jerryAmount = toWei('2000000')

        // mint tom & jerry and add to liquidity pool
        await tomToken.methods.mint(tomAmount).send(actors.ownerTx)
        await jerryToken.methods.mint(jerryAmount).send(actors.ownerTx)
        await tomToken.methods.approve(poolAddr, tomAmount).send(actors.ownerTx)
        await jerryToken.methods.approve(poolAddr, jerryAmount).send(actors.ownerTx)
        await pool.methods.addLiquidity(tomAmount, jerryAmount).send(actors.ownerTx)

        // swap 10 tom -> 19.96 jerry
        const tomIn = toWei('10')
        const quotePrice = await pool.methods.getJerryAmountByTom(tomIn).call()
        const jerryOut = parseFloat(fromWei(quotePrice)) * 0.998
        await tomToken.methods.mint(tomIn).send(actors.user1Tx)
        await tomToken.methods.approve(poolAddr, tomIn).send(actors.user1Tx)
        await pool.methods.swapTomforJerry(tomIn, toWei(`${jerryOut}`)).send(actors.user1Tx)
        const receivedJerry = await jerryToken.methods.balanceOf(actors.user1Addr).call()
        expect(fromWei(receivedJerry)).toEqual(`${jerryOut}`)
    })
})