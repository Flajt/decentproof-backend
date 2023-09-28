"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const getNewKey_1 = require("../../functions/getNewKey");
const chai_1 = require("chai");
const dotenv = require("dotenv");
dotenv.config();
describe("getNewKey", () => {
    it("should return status code 401 on invalid token", async () => {
        let resp = await (0, getNewKey_1.handler)({
            pathParameters: {},
            queryStringParameters: {},
            body: "",
            headers: { "X-AppCheck": "An invalid token string here" },
            method: "",
            isBase64Encoded: false
        }, null, null);
        (0, chai_1.assert)(resp.statusCode === 401);
    });
});
