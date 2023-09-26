"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
async function verifyApiKey(apiKey, deta) {
    const db = deta.Base("keys");
    let keys = await db.fetch();
    if (apiKey === null || apiKey === undefined)
        return false;
    for (let key of keys.items) {
        if (key.key === apiKey)
            return true;
    }
    return false;
}
exports.default = verifyApiKey;
