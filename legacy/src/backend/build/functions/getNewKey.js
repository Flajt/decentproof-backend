"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.handler = void 0;
const getNewestKey_1 = __importDefault(require("../helper/apikey/getNewestKey"));
const appCheck_1 = __importDefault(require("../helper/auth/appCheck"));
const DetaSingleton_1 = __importDefault(require("../helper/shared/DetaSingleton"));
async function handler(event, context, callback) {
    var _a;
    try {
        await DetaSingleton_1.default.Instance.init();
        //TODO: Find out why X-AppCheck is converted to X-Appcheck! something seems broken
        let isDeviceValid = await (0, appCheck_1.default)((_a = event.headers["X-Appcheck"]) !== null && _a !== void 0 ? _a : "No token", DetaSingleton_1.default.Instance);
        if (!isDeviceValid)
            return {
                statusCode: 401,
                body: "Invalid device",
                headers: { "Content-Type": "text/plain" }
            };
        const deta = DetaSingleton_1.default.Instance.deta;
        let newestKey = await (0, getNewestKey_1.default)(deta);
        console.info("Request successfulls");
        return { statusCode: 200, headers: { "Content-Type": "application/json" }, body: JSON.stringify({ key: newestKey }) };
    }
    catch (e) {
        console.error("Something went wrong: " + e);
        return {
            statusCode: 500, body: "Something went wrong", headers: { "Content-Type": "text/plain" }
        };
    }
}
exports.handler = handler;
