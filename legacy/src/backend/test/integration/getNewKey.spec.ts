import {handler} from "../../functions/getNewKey";
import {assert} from "chai";
const dotenv = require("dotenv");

dotenv.config();

describe("getNewKey",()=>{
    it("should return status code 401 on invalid token",async ()=>{
        let resp = await handler({
            pathParameters: {},
            queryStringParameters: {},
            body: "",
            headers: {"X-AppCheck": "An invalid token string here"},
            method: "",
            isBase64Encoded: false
        },null,null);
        assert(resp.statusCode === 401);
    })
})