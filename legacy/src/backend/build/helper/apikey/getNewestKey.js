"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
async function getNewestKey(deta) {
    const db = deta.Base("keys");
    let allKeys = await db.fetch();
    if (allKeys.count > 0) {
        let newestKey = allKeys.items.reduce((prevKey, currentKey) => {
            //TODO: Consider if using the raw numbers is enough instead of converting it to a date
            if (new Date(currentKey.__expires) > new Date(prevKey.__expires)) {
                return currentKey;
            }
            else {
                return prevKey;
            }
        });
        return newestKey.key;
    }
    else {
        throw "No keys found! Something is wrong!";
    }
}
exports.default = getNewestKey;
