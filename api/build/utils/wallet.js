"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.getWallet = getWallet;
// Wallet singleton
const fabric_network_1 = require("fabric-network");
const path_1 = __importDefault(require("path"));
let walletInstance = null;
async function getWallet() {
    if (walletInstance) {
        return walletInstance;
    }
    try {
        const walletPath = path_1.default.resolve(__dirname, '..', '..', 'wallet');
        walletInstance = await fabric_network_1.Wallets.newFileSystemWallet(walletPath);
        return walletInstance;
    }
    catch (err) {
        console.log('Erro ao criar wallet!');
        throw err;
    }
}
