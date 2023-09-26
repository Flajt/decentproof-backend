"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.handler = void 0;
const jsrsasign_1 = require("jsrsasign");
const authMiddleware_1 = __importDefault(require("../helper/auth/authMiddleware"));
const DetaSingleton_1 = __importDefault(require("../helper/shared/DetaSingleton"));
const KeySingleton_1 = __importDefault(require("../helper/shared/KeySingleton"));
async function handler(event, context, callback) {
    try {
        let isValid = await (0, authMiddleware_1.default)(event.headers);
        if (!isValid) {
            console.info("Invalid API key");
            return {
                statusCode: 401,
                body: "Access Denied",
                headers: { "Content-Type": "text/plain" }
            };
        }
        await DetaSingleton_1.default.Instance.init();
        await KeySingleton_1.default.Instance.init(DetaSingleton_1.default.Instance.keys);
        const singleton = KeySingleton_1.default.Instance;
        const signatureManagerPriv = new jsrsasign_1.KJUR.crypto.Signature({ alg: "SHA224withECDSA" });
        signatureManagerPriv.init(singleton.privKey);
        let body = JSON.parse(event.body);
        signatureManagerPriv.updateString(body.data);
        let signature = signatureManagerPriv.sign();
        let respBody = {
            statusCode: 200,
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ "signature": signature }),
        };
        console.info("Request successfull");
        return respBody;
    }
    catch (e) {
        console.error("Somehting went wrong: " + e);
        return {
            statusCode: 500,
            headers: { "Content-Type": "text/plain" },
            body: "Something went wrong",
        };
    }
}
exports.handler = handler;
