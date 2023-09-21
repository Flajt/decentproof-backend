"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const verifyApikey_1 = __importDefault(require("../apikey/verifyApikey"));
const DetaSingleton_1 = __importDefault(require("../shared/DetaSingleton"));
async function authCheck(headers) {
    if (process.env.NODE_ENV === "production") {
        await DetaSingleton_1.default.Instance.init();
        const deta = DetaSingleton_1.default.Instance.deta;
        let authHeader = headers["Authorization"];
        if (authHeader !== null && authHeader !== undefined && authHeader.length > 5) {
            let apiKey = authHeader.split(" ")[1];
            let isValid = await (0, verifyApikey_1.default)(apiKey, deta);
            if (isValid)
                return true;
            return false;
        }
        else {
            return false;
        }
    }
    else {
        return true;
    }
}
exports.default = authCheck;
