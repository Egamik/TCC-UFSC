"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.enrollAdmin = enrollAdmin;
exports.registerAndEnrollUser = registerAndEnrollUser;
exports.setActiveID = setActiveID;
exports.getActiveID = getActiveID;
exports.listIDs = listIDs;
const fs_1 = __importDefault(require("fs"));
const fabric_ca_client_1 = __importDefault(require("fabric-ca-client"));
const wallet_js_1 = require("../utils/wallet.js");
const networkConfig_js_1 = require("../config/networkConfig.js");
const path_1 = __importDefault(require("path"));
const activePath = path_1.default.join(__dirname, '../../fabric/activeIdentity.json');
const ca = new fabric_ca_client_1.default(networkConfig_js_1.config.caURL, { trustedRoots: [], verify: false }, networkConfig_js_1.config.caName);
// Enrolls default Admin user's identity
async function enrollAdmin() {
    const wallet = await (0, wallet_js_1.getWallet)();
    const adminExists = await wallet.get(networkConfig_js_1.config.adminUserId);
    if (adminExists)
        return;
    const enrollment = await ca.enroll({
        enrollmentID: networkConfig_js_1.config.adminUserId,
        enrollmentSecret: networkConfig_js_1.config.adminUserPass
    });
    const identity = {
        credentials: {
            certificate: enrollment.certificate,
            privateKey: enrollment.key.toBytes()
        },
        mspId: networkConfig_js_1.config.mspId,
        type: 'X.509'
    };
    await wallet.put(networkConfig_js_1.config.adminUserId, identity);
}
// Create new identity for user with given ID and role
async function registerAndEnrollUser(userId, role) {
    const wallet = await (0, wallet_js_1.getWallet)();
    const userExists = await wallet.get(userId);
    if (userExists)
        return;
    const adminIdentity = await wallet.get(networkConfig_js_1.config.adminUserId);
    if (!adminIdentity)
        throw new Error('Admin identity not enrolled');
    const provider = wallet.getProviderRegistry().getProvider(adminIdentity.type);
    const adminUser = await provider.getUserContext(adminIdentity, networkConfig_js_1.config.adminUserId);
    const secret = await ca.register({
        affiliation: 'org0.example.com',
        enrollmentID: userId,
        attrs: [
            { name: 'role', value: role, ecert: true },
            { name: 'personID', value: userId, ecert: true }
        ]
    }, adminUser);
    const enrollment = await ca.enroll({ enrollmentID: userId, enrollmentSecret: secret });
    const userIdentity = {
        credentials: {
            certificate: enrollment.certificate,
            privateKey: enrollment.key.toBytes()
        },
        mspId: networkConfig_js_1.config.mspId,
        type: 'X.509'
    };
    await wallet.put(userId, userIdentity);
}
// Sets active ID in activeIndetity.json file
async function setActiveID(id) {
    const wallet = await (0, wallet_js_1.getWallet)();
    const identity = await wallet.get(id);
    if (!identity) {
        throw new Error(`Identity: ${identity} does not exist.`);
    }
    fs_1.default.writeFileSync(activePath, JSON.stringify({ active: id }, null, 2));
}
// Gets active ID in activeIdentity.json file
function getActiveID() {
    try {
        if (!fs_1.default.existsSync(activePath))
            return null;
        const content = JSON.parse(fs_1.default.readFileSync(activePath, 'utf8'));
        return content.active;
    }
    catch (err) {
        console.log('Error [getActiveIdentity]: ', err);
        return null;
    }
}
async function listIDs() {
    const wallet = await (0, wallet_js_1.getWallet)();
    const ids = await wallet.list();
    return ids;
}
