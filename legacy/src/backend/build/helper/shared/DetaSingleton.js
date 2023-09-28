"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const deta_1 = require("deta");
//Based on: https://stackoverflow.com/questions/30174078/how-to-define-singleton-in-typescript#36978360
class DetaSingleton {
    async init() {
        this.deta = (0, deta_1.Deta)(process.env.PROJECT_KEY);
        this.keys = this.deta.Drive("keys");
        let firebaseKeyBuf = await this.keys.get("firebaseKey.json");
        let buf = await (firebaseKeyBuf === null || firebaseKeyBuf === void 0 ? void 0 : firebaseKeyBuf.arrayBuffer());
        this.fireBaseKey = JSON.parse(Buffer.from(buf).toString("utf-8"));
    }
    static get Instance() {
        return this._instance || (this._instance = new this());
    }
}
exports.default = DetaSingleton;
