import Deta from "deta/dist/types/deta";

async function verifyApiKey(apiKey: string, deta: Deta) {
    const db = deta.Base("keys")
    let keys = await db.fetch()
    if(apiKey === null|| apiKey === undefined)
        return false;
    for (let key of keys.items) {
        if (key.key === apiKey)
            return true;
    }
    return false;

}
export default verifyApiKey;