"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const chai_1 = require("chai");
const sign_1 = require("../../functions/sign");
const dotenv = require("dotenv");
dotenv.config();
describe("sign", () => {
    it("should prevent access with invalid Authorization header", async () => {
        let resp = await (0, sign_1.handler)({
            body: JSON.stringify({ "data": "abcdefghijklmnopqrstuvwzyz" }),
            method: "POST",
            pathParameters: {},
            isBase64Encoded: false,
            queryStringParameters: {},
            headers: {
                "Authorization": "lalala"
            }
        }, null, null);
        (0, chai_1.assert)(resp.statusCode == 401);
    });
});
