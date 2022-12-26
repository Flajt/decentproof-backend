import Deta from "deta/dist/types/deta";

async function getNewestKey(deta: Deta): Promise<string> {
    const db = deta.Base("keys")
    let allKeys = await db.fetch();
    if (allKeys.count > 0) {
        let newestKey = allKeys.items.reduce((prevKey, currentKey) => {
            //TODO: Consider if using the raw numbers is enough instead of converting it to a date
            if (new Date(currentKey.__expires as number) > new Date(prevKey.__expires as number)) {
                return currentKey;
            } else {
                return prevKey;
            }
        })
        return newestKey.key as string;
    } else {
        throw "No keys found! Something is wrong!"
    }
}
export default getNewestKey;