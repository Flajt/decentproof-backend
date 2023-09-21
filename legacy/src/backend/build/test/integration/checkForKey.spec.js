"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const chai_1 = require("chai");
const mocha_1 = require("mocha");
const hasNewKey_1 = require("../../functions/hasNewKey");
const dotenv = require("dotenv");
dotenv.config();
(0, mocha_1.describe)("hasNewKey.handler", () => {
    (0, mocha_1.it)("should return true if provided with a random input", async () => {
        let testKey = "avampkbapkvpnbnqblnebl"; //random not correct key
        let requestData = {
            pathParameters: {},
            queryStringParameters: {},
            body: "",
            headers: { "Authorization": "bearer " + testKey },
            method: "",
            isBase64Encoded: false
        };
        let respBody = await (0, hasNewKey_1.handler)(requestData, null, () => { });
        (0, chai_1.assert)(respBody.body.valueOf().hasNewer == true); // not motivated enough to parse into an object
    });
});
