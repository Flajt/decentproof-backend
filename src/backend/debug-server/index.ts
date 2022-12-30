import express from "express";
import { handler } from "../functions/hasNewKey";
import { FunctionEvent } from "../helper/types/functionEvent";
import signingHandler from "./signingHandlerExporter";
import getKeyHandler from "./getKeyHandlerExporter";
const dotenv = require("dotenv")

dotenv.config();
const app = express();


app.get("/has-new",async (req,resp)=>{
    let event:FunctionEvent = {
        headers: req.headers,
        body: "",
        pathParameters: {},
        queryStringParameters: {},
        method: "GET",
        isBase64Encoded: false
    };
    let respBody = await handler(event,null,null);
    return resp.status(respBody.statusCode).json(respBody.body)
})

app.post("/sign",async (req,resp)=>{
    let event:FunctionEvent = {
        pathParameters: "",
        queryStringParameters: {},
        body: req.body,
        headers: req.headers,
        method: "POST",
        isBase64Encoded: false
    } 
    let respBody = await signingHandler(event,null,null);

})
app.get("/new-key",async (req,resp)=>{
    let event:FunctionEvent = {
        pathParameters: "",
        queryStringParameters: {},
        body: req.body,
        headers: {"X-Appcheck":req.headers["x-appcheck"]},
        method: "POST",
        isBase64Encoded: false
    } 
    let respBody = await getKeyHandler(event,null,null);
    resp.set("Content-Type","application/json");
    return resp.status(respBody.statusCode).json(respBody.body);
});

app.get("/test",(req,resp)=>{
    return resp.status(200).json(JSON.stringify({data:"ok"}))
})
app.listen(8000,()=>console.info("Running on PORT 8000"))