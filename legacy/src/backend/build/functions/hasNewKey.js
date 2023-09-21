"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.handler = void 0;
const checkForNewKey_1 = __importDefault(require("../helper/apikey/checkForNewKey"));
const DetaSingleton_1 = __importDefault(require("../helper/shared/DetaSingleton"));
async function handler(event, context, callback) {
    try {
        await DetaSingleton_1.default.Instance.init();
        const deta = DetaSingleton_1.default.Instance.deta;
        let key = event.headers["Authorization"].split(" ")[1];
        let hasNewer = await (0, checkForNewKey_1.default)(key !== null && key !== void 0 ? key : "no key", deta);
        let body = {
            statusCode: 200,
            headers: { "Content-Type": "application/json" },
            body: { hasNewer: hasNewer },
        };
        console.info("Request successfull");
        return body;
    }
    catch (e) {
        console.error("Something went wrong: " + e);
        return {
            headers: { "Content-Type": "text/plain" },
            statusCode: 500,
            body: "Something went wrong",
        };
    }
}
exports.handler = handler;
