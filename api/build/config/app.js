"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.app = void 0;
const express_1 = __importDefault(require("express"));
const wallet_router_js_1 = __importDefault(require("../router/wallet.router.js"));
const caService_js_1 = require("../services/caService.js");
const createApp = () => {
    const app = (0, express_1.default)();
    // Config Fabric client
    (0, caService_js_1.enrollAdmin)().then(() => {
        console.log('Admin enrolled successfully!');
    }).catch((err) => {
        console.log('Error enrolling admin: ', err);
        process.exit(1);
    });
    // Middleware
    app.use(express_1.default.json());
    // Routes
    app.use('/wallet', wallet_router_js_1.default);
    // app.use('/app', appRouter)
    return app;
};
exports.app = createApp();
