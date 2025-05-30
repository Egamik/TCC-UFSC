import { Router } from "express"
import {createIdentity, getActiveIdentity, setActiveIdentity, listIdentities} from '../controller/wallets.controller.js'

const router = Router()

router.post('/', createIdentity)
router.post('/active', setActiveIdentity)
router.get('/active', getActiveIdentity)
router.get('/list', listIdentities)

export default router