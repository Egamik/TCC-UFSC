"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getGatewayAndContract = getGatewayAndContract;
const fabric_network_1 = require("fabric-network");
const fs_1 = __importDefault(require("fs"));
const wallet_js_1 = require("../utils/wallet.js");
const networkConfig_js_1 = require("../config/networkConfig.js");
const ccp = JSON.parse(fs_1.default.readFileSync(networkConfig_js_1.config.connectionProfile, 'utf8'));
async function getGatewayAndContract(userId, channelName, chaincodeName) {
    const wallet = await (0, wallet_js_1.getWallet)();
    const identity = await wallet.get(userId);
    if (!identity)
        throw new Error(`Identity for user ${userId} not found`);
    const gateway = new fabric_network_1.Gateway();
    await gateway.connect(ccp, {
        wallet,
        identity: userId,
        discovery: { enabled: true, asLocalhost: false }
    });
    const network = await gateway.getNetwork(channelName);
    const contract = network.getContract(chaincodeName);
    return { contract, gateway };
}
