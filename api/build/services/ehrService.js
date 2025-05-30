"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.createRecord = createRecord;
exports.addPrescription = addPrescription;
exports.addAppointment = addAppointment;
exports.addProcedure = addProcedure;
exports.readRecord = readRecord;
exports.readPrescriptions = readPrescriptions;
exports.readAppointments = readAppointments;
exports.readProcedures = readProcedures;
exports.recordExists = recordExists;
async function createRecord(contract, ownerID) {
    return await contract.submitTransaction('CreateRecord', ownerID);
}
async function addPrescription(contract, ownerID, prescriptionJSON) {
    try {
        return await contract.submitTransaction('AddPrescription', ownerID, prescriptionJSON);
    }
    catch (err) {
        console.log('Error submiting transaction: ', err);
        throw err;
    }
}
async function addAppointment(contract, ownerID, appointmentJSON) {
    try {
        return await contract.submitTransaction('AddAppointment', ownerID, appointmentJSON);
    }
    catch (err) {
        console.log('Error submiting transaction: ', err);
        throw err;
    }
}
async function addProcedure(contract, ownerID, procedureJSON) {
    try {
        return await contract.submitTransaction('AddProcedure', ownerID, procedureJSON);
    }
    catch (err) {
        console.log('Error submiting transaction: ', err);
        throw err;
    }
}
async function readRecord(contract, ownerID) {
    try {
        return await contract.evaluateTransaction('ReadRecord', ownerID);
    }
    catch (err) {
        console.log('Error submiting transaction: ', err);
        throw err;
    }
}
async function readPrescriptions(contract, ownerID, professionalID) {
    try {
        return await contract.evaluateTransaction('ReadPrescription', ownerID, professionalID);
    }
    catch (err) {
        console.log('Error evaluating transaction: ', err);
        throw err;
    }
}
async function readAppointments(contract, ownerID, professionalID) {
    try {
        return await contract.evaluateTransaction('ReadAppointments', ownerID, professionalID);
    }
    catch (err) {
        console.log('Error evaluating transaction: ', err);
        throw err;
    }
}
async function readProcedures(contract, ownerID, professionalID) {
    try {
        return await contract.evaluateTransaction('ReadProcedures', ownerID, professionalID);
    }
    catch (err) {
        console.log('Error evaluating transaction: ', err);
        throw err;
    }
}
async function recordExists(contract, ownerID) {
    try {
        return await contract.evaluateTransaction('RecordExists', ownerID);
    }
    catch (err) {
        console.log('Error evaluating transaction: ', err);
        throw err;
    }
}
