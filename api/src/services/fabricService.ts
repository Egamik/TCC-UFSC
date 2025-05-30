import { Gateway, Contract } from 'fabric-network'
import fs from 'fs'
import { getWallet } from '../utils/wallet.js'
import { config } from '../config/networkConfig.js'

const ccp = JSON.parse(fs.readFileSync(config.connectionProfile, 'utf8'))

export async function getGatewayAndContract(userId: string, channelName: string, chaincodeName: string): Promise<{ contract: Contract; gateway: Gateway }> {
    const wallet = await getWallet()
    const identity = await wallet.get(userId)
    if (!identity) throw new Error(`Identity for user ${userId} not found`)

    const gateway = new Gateway()
    await gateway.connect(ccp, {
        wallet,
        identity: userId,
        discovery: { enabled: true, asLocalhost: false }
    })

    const network = await gateway.getNetwork(channelName)
    const contract = network.getContract(chaincodeName)

    return { contract, gateway }
}
