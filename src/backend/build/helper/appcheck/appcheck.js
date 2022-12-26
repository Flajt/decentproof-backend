"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const firebase_admin_1 = require("firebase-admin");
class AppcheckWrapper {
    async verifyToken(token) {
        try {
            console.log(token);
            let isValid = await (0, firebase_admin_1.appCheck)().verifyToken(token);
            console.log(isValid);
            if (isValid)
                return true;
            return false;
        }
        catch (e) {
            console.error("Cant verify token due to: " + e);
            return false;
        }
    }
}
exports.default = AppcheckWrapper;
