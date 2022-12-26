"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const jsrsasign_1 = require("jsrsasign");
//Based on: https://stackoverflow.com/questions/30174078/how-to-define-singleton-in-typescript#36978360
class KeySingleton {
    constructor() {
        this.privKey = null;
        this.pubKey = null;
    }
    async init(drive) {
        let privResp = await drive.get("priv.pem");
        let pubResp = await drive.get("pub.pem");
        let privBuff = await (privResp === null || privResp === void 0 ? void 0 : privResp.arrayBuffer());
        let pubBuff = await (pubResp === null || pubResp === void 0 ? void 0 : pubResp.arrayBuffer());
        this.privKey = new jsrsasign_1.KJUR.crypto.ECDSA({ curve: "secp256r1", prv: Buffer.from(privBuff).toString("utf-8") });
        this.pubKey = new jsrsasign_1.KJUR.crypto.ECDSA({ curve: "secp256r1", pub: Buffer.from(pubBuff).toString("utf-8") });
    }
    static get Instance() {
        return this._instance || (this._instance = new this());
    }
}
exports.default = KeySingleton;
