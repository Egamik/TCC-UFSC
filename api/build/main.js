"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const app_js_1 = require("./config/app.js");
const https_1 = __importDefault(require("https"));
const fs_1 = __importDefault(require("fs"));
const port = 8080;
// const app = express()
const options = {
    key: fs_1.default.readFileSync(''),
    cert: fs_1.default.readFileSync(''),
    ca: [fs_1.default.readFileSync('')],
};
const server = https_1.default.createServer(options, app_js_1.app);
server.listen(port, () => {
    console.log(`Server running on port ${port}`);
});
