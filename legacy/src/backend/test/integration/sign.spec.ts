import { assert } from "chai";
import {handler} from "../../functions/sign"

const dotenv = require("dotenv");
dotenv.config();
describe("sign",()=>{
    it("should prevent access with invalid Authorization header",async ()=>{
        let resp = await handler({
            body: JSON.stringify({"data":"abcdefghijklmnopqrstuvwzyz"}),
            method: "POST",
            pathParameters: {},
            isBase64Encoded: false,
            queryStringParameters: {},
            headers: {
                "Authorization": "lalala"
            }
        },null,null);
        assert(resp.statusCode==401);
    });
});