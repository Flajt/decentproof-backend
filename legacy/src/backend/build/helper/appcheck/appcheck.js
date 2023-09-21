"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const firebase_admin_1 = require("firebase-admin");
class AppcheckWrapper {
    async verifyToken(token) {
        try {
            let isValid = await (0, firebase_admin_1.appCheck)().verifyToken(token);
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
