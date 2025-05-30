"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createAsset = createAsset;
exports.addIdentity = addIdentity;
exports.removeIdentity = removeIdentity;
exports.isIdentityApproved = isIdentityApproved;
exports.getIdentityList = getIdentityList;
async function createAsset(contract, ownerID) {
    try {
        contract.submitTransaction('CreateAsset', ownerID);
    }
    catch (err) {
        console.log('Error submiting transaction: ', err);
        throw err;
    }
}
async function addIdentity(contract, ownerID, professionalID) {
    try {
        contract.submitTransaction('AddIdentity', ownerID, professionalID);
    }
    catch (err) {
        console.log('Error submiting transaction: ', err);
        throw err;
    }
}
async function removeIdentity(contract, ownerID, professionalID) {
    try {
        contract.submitTransaction('RemoveIdentity', ownerID, professionalID);
    }
    catch (err) {
        console.log('Error submiting transaction: ', err);
        throw err;
    }
}
async function isIdentityApproved(contract, ownerID, professionalID) {
    try {
        contract.evaluateTransaction('IsIdentityApproved', ownerID, professionalID);
    }
    catch (err) {
        console.log('Error evaluating transaction: ', err);
        throw err;
    }
}
async function getIdentityList(contract, ownerID) {
    try {
        contract.evaluateTransaction('GetIdentityList', ownerID);
    }
    catch (err) {
        console.log('Error evaluating transaction: ', err);
        throw err;
    }
}
