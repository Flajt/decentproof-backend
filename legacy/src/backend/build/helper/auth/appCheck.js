"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const app_1 = require("firebase-admin/app");
const appcheck_1 = __importDefault(require("../appcheck/appcheck"));
const app_2 = require("firebase-admin/app");
async function appCheck(appCheckToken, detaSingleton) {
    if (process.env.NODE_ENV === "production") {
        (0, app_1.getApps)().length === 0 ? (0, app_1.initializeApp)({
            credential: (0, app_2.cert)(detaSingleton.fireBaseKey)
        }) : null;
        const appcheckWrapper = new appcheck_1.default();
        if (!appCheckToken || appCheckToken.length <= 1) {
            console.info("no token");
            return false;
        }
        else {
            let isValid = await appcheckWrapper.verifyToken(appCheckToken);
            if (isValid)
                return true;
            else
                return false;
        }
    }
    else {
        return true;
    }
}
exports.default = appCheck;
