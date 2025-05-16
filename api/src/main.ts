import { app } from './config/app.js'
import https from 'https'
import fs from 'fs'

const port = 8080
// const app = express()

const options = {
    key: fs.readFileSync(''),
    cert: fs.readFileSync(''),
    ca: [fs.readFileSync('')],
}

const server = https.createServer(options, app)

server.listen(port, () => {
    console.log(`Server running on port ${port}`)
})