const generateApiKey = require("./generateApiKey")
async function updateApiKeys(deta) {
    let db = await deta.Base("keys");
    let newKey = generateApiKey();
    await db.insert({ key: newKey }, undefined, { expireIn: 1814400 })//3 Weeks in seconds
    console.log("Keys updated!")
}
module.exports = updateApiKeys;