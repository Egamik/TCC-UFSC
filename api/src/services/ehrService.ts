import { Contract } from "fabric-network";

async function createRecord(contract: Contract, ownerID: string) {
    return await contract.submitTransaction('CreateRecord', ownerID)
}

async function addPrescription(contract: Contract, ownerID: string, prescriptionJSON: string) {
    try {
        return await contract.submitTransaction('AddPrescription', ownerID, prescriptionJSON)
    } catch (err) {
        console.log('Error submiting transaction: ', err)
        throw err
    }
}

async function addAppointment(contract: Contract, ownerID: string, appointmentJSON: string) {
    try {
        return await contract.submitTransaction('AddAppointment', ownerID, appointmentJSON)
    } catch (err) {
        console.log('Error submiting transaction: ', err)
        throw err
    }
}

async function addProcedure(contract: Contract, ownerID: string, procedureJSON: string) {
    try {
        return await contract.submitTransaction('AddProcedure', ownerID, procedureJSON)
    } catch (err) {
        console.log('Error submiting transaction: ', err)
        throw err
    }
}

async function readRecord(contract: Contract, ownerID: string) {
    try {
        return await contract.evaluateTransaction('ReadRecord', ownerID)
    } catch (err) {
        console.log('Error submiting transaction: ', err)
        throw err
    }
}

async function readPrescriptions(contract: Contract, ownerID: string, professionalID: string) {
    try {
        return await contract.evaluateTransaction('ReadPrescription', ownerID, professionalID)
    } catch (err) {
        console.log('Error evaluating transaction: ', err)
        throw err
    }
}

async function readAppointments(contract: Contract, ownerID: string, professionalID: string) {
    try {
        return await contract.evaluateTransaction('ReadAppointments', ownerID, professionalID)
    } catch (err) {
        console.log('Error evaluating transaction: ', err)
        throw err
    }
}

async function readProcedures(contract: Contract, ownerID: string, professionalID: string) {
    try {
        return await contract.evaluateTransaction('ReadProcedures', ownerID, professionalID)
    } catch (err) {
        console.log('Error evaluating transaction: ', err)
        throw err
    }
}

async function recordExists(contract: Contract, ownerID: string) {
    try {
        return await contract.evaluateTransaction('RecordExists', ownerID)
    } catch (err) {
        console.log('Error evaluating transaction: ', err)
        throw err
    }
}

export {
    createRecord,
    addPrescription,
    addAppointment,
    addProcedure,
    readRecord,
    readPrescriptions,
    readAppointments,
    readProcedures,
    recordExists,
}