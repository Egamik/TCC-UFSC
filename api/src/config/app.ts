import express, { Application } from 'express'
import walletRouter from '../router/wallet.router.js'
import fabricRouter from '../router/fabric.router.js'


const createApp = (): Application => {
    const app = express()

    // Middleware
    app.use(express.json())

    // Routes
    app.use('/wallet', walletRouter)
    app.use('/fabric', fabricRouter)

    return app
}

export const app = createApp()