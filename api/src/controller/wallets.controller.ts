import { Request, Response } from "express"
import { registerAndEnrollUser, getActiveID, setActiveID, listIDs } from "../services/caService.js"

const createIdentity = async (req: Request, res: Response) => {
    const { userId, role } = req.body
    if (!userId || !role) {
        res.status(400).json({error: 'Faltam argumentos!'})
        return
    }
    try {
        await registerAndEnrollUser(userId, role)
        res.status(200).json({message: 'User enrolled', userId})
        return
    } catch (err) {
        res.status(500).json({error: err})
    }
}

const setActiveIdentity = async (req: Request, res: Response) => {
    const { id } = req.body
    try {
        await setActiveID(id)
        res.status(200).json({ message: `Identity ${id} is now active.` })
    } catch (err) {
        console.log('Error [SetActiveIdentity]: ', err)
        res.status(500).json({error: err})
    }
}

const getActiveIdentity = async (req: Request, res: Response) => {
    try {
        const active = await getActiveID()
        res.status(200).json({ message: active})
        return
    } catch (err) {
        console.log('Error [getActiveIdentity]: ', err)
        res.status(500).json({error: err})
    }
}

const listIdentities = async (req: Request, res: Response) => {
    try {
        const ids = await listIDs()
        res.status(200).json({ message: ids })
    } catch (err) {
        console.log('Error [listIdentities]: ', err)
        res.status(500).json({ error: err})
    }
}

export {
    createIdentity,
    setActiveIdentity,
    getActiveIdentity,
    listIdentities
}