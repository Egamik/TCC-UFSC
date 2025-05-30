"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.listIdentities = exports.getActiveIdentity = exports.setActiveIdentity = exports.createIdentity = void 0;
const caService_js_1 = require("../services/caService.js");
const createIdentity = async (req, res) => {
    const { userId, role } = req.body;
    if (!userId || !role) {
        res.status(400).json({ error: 'Faltam argumentos!' });
        return;
    }
    try {
        await (0, caService_js_1.registerAndEnrollUser)(userId, role);
        res.status(200).json({ message: 'User enrolled', userId });
        return;
    }
    catch (err) {
        res.status(500).json({ error: err });
    }
};
exports.createIdentity = createIdentity;
const setActiveIdentity = async (req, res) => {
    const { id } = req.body;
    try {
        await (0, caService_js_1.setActiveID)(id);
        res.status(200).json({ message: `Identity ${id} is now active.` });
    }
    catch (err) {
        console.log('Error [SetActiveIdentity]: ', err);
        res.status(500).json({ error: err });
    }
};
exports.setActiveIdentity = setActiveIdentity;
const getActiveIdentity = async (req, res) => {
    try {
        const active = await (0, caService_js_1.getActiveID)();
        res.status(200).json({ message: active });
        return;
    }
    catch (err) {
        console.log('Error [getActiveIdentity]: ', err);
        res.status(500).json({ error: err });
    }
};
exports.getActiveIdentity = getActiveIdentity;
const listIdentities = async (req, res) => {
    try {
        const ids = await (0, caService_js_1.listIDs)();
        res.status(200).json({ message: ids });
    }
    catch (err) {
        console.log('Error [listIdentities]: ', err);
        res.status(500).json({ error: err });
    }
};
exports.listIdentities = listIdentities;
