const {randomBytes} = require("crypto")

function generateApiKey(){
    let randomBuffer = randomBytes(30)
    return randomBuffer.toString("base64");
}

module.exports = generateApiKey;