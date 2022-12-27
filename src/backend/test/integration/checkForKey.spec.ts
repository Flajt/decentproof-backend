import { assert } from "chai";
import { describe, it } from "mocha";
import {handler} from "../../functions/hasNewKey";
const dotenv = require("dotenv");

dotenv.config();

describe("hasNewKey.handler",()=>{
    it("should return true if provided with a random input",async ()=>{
        let testKey = "avampkbapkvpnbnqblnebl"; //random not correct key
        let requestData = {
            pathParameters: {},
            queryStringParameters: {},
            body: "",
            headers: {"Authorization": "bearer "+testKey},
            method: "",
            isBase64Encoded: false
        } 
        let respBody = await handler(
            requestData
        ,null,()=>{});
        assert((respBody.body.valueOf() as any).hasNewer==true) // not motivated enough to parse into an object
    })
})