import fs from 'fs'
import FabricCAServices from 'fabric-ca-client'
import { X509Identity } from 'fabric-network'
import { getWallet } from '../utils/wallet.js'
import { config } from '../config/networkConfig.js'
import path from 'path'

const activePath = path.join(__dirname, '../../fabric/activeIdentity.json')

const ca = new FabricCAServices(config.caURL, { trustedRoots: [], verify: false }, config.caName);

// Enrolls default Admin user's identity
async function enrollAdmin() {
    const wallet = await getWallet()
    const adminExists = await wallet.get(config.adminUserId)
    if (adminExists) return

    const enrollment = await ca.enroll({
        enrollmentID: config.adminUserId,
        enrollmentSecret: config.adminUserPass
    })

    const identity: X509Identity = {
        credentials: {
            certificate: enrollment.certificate,
            privateKey: enrollment.key.toBytes()
        },
        mspId: config.mspId,
        type: 'X.509'
    }

    await wallet.put(config.adminUserId, identity)                                                                                                                                                                                                                                                                                                                                                                                                        
}

// Create new identity for user with given ID and role
async function registerAndEnrollUser(userId: string, role: string) {
    const wallet = await getWallet()
    const userExists = await wallet.get(userId);
    if (userExists) return;

    const adminIdentity = await wallet.get(config.adminUserId);
    if (!adminIdentity) throw new Error('Admin identity not enrolled');

    const provider = wallet.getProviderRegistry().getProvider(adminIdentity.type);
    const adminUser = await provider.getUserContext(adminIdentity, config.adminUserId);

    const secret = await ca.register({
        affiliation: 'org0.example.com',
        enrollmentID: userId,
        attrs: [
            { name: 'role', value: role, ecert: true },
            { name: 'personID', value: userId, ecert: true }
        ]
    }, adminUser);

    const enrollment = await ca.enroll({ enrollmentID: userId, enrollmentSecret: secret });

    const userIdentity: X509Identity = {
        credentials: {
        certificate: enrollment.certificate,
        privateKey: enrollment.key.toBytes()
        },
        mspId: config.mspId,
        type: 'X.509'
    };

    await wallet.put(userId, userIdentity);
}

// Sets active ID in activeIndetity.json file
async function setActiveID(id: string) {
    const wallet = await getWallet()
    const identity = await wallet.get(id)
    if (!identity) {
        throw new Error(`Identity: ${identity} does not exist.`)
    }

    fs.writeFileSync(activePath, JSON.stringify({active: id}, null, 2))
}

// Gets active ID in activeIdentity.json file
function getActiveID() {
    try {
        if (!fs.existsSync(activePath)) return null
        const content = JSON.parse(fs.readFileSync(activePath, 'utf8'))
        return content.active
    } catch (err) {
        console.log('Error [getActiveIdentity]: ', err)
        return null
    }
}

async function listIDs() {
    const wallet = await getWallet()
    const ids = await wallet.list()
    return ids
}

export {
    enrollAdmin,
    registerAndEnrollUser,
    setActiveID,
    getActiveID,
    listIDs
}