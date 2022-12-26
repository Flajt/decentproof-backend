"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
async function checkForNewKey(key, deta) {
    try {
        const db = deta.Base("keys");
        let currentKey = await db.get(key);
        if (currentKey) {
            let expire = currentKey.__expires;
            let keys = await db.fetch();
            let hasNewer = false;
            for (let key of keys.items) {
                if (new Date(key.__expires) > new Date(expire)) {
                    hasNewer = true;
                    break;
                }
            }
            if (hasNewer)
                return true;
        }
        return true;
    }
    catch (e) {
        return true;
    }
}
exports.default = checkForNewKey;
