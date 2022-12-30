"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const hasNewKey_1 = require("../functions/hasNewKey");
const signingHandlerExporter_1 = __importDefault(require("./signingHandlerExporter"));
const getKeyHandlerExporter_1 = __importDefault(require("./getKeyHandlerExporter"));
const dotenv = require("dotenv");
dotenv.config();
const app = (0, express_1.default)();
app.get("/has-new", async (req, resp) => {
    let event = {
        headers: req.headers,
        body: "",
        pathParameters: {},
        queryStringParameters: {},
        method: "GET",
        isBase64Encoded: false
    };
    let respBody = await (0, hasNewKey_1.handler)(event, null, null);
    return resp.status(respBody.statusCode).json(respBody.body);
});
app.post("/sign", async (req, resp) => {
    let event = {
        pathParameters: "",
        queryStringParameters: {},
        body: req.body,
        headers: req.headers,
        method: "POST",
        isBase64Encoded: false
    };
    let respBody = await (0, signingHandlerExporter_1.default)(event, null, null);
});
app.get("/new-key", async (req, resp) => {
    let event = {
        pathParameters: "",
        queryStringParameters: {},
        body: req.body,
        headers: { "X-Appcheck": req.headers["x-appcheck"] },
        method: "POST",
        isBase64Encoded: false
    };
    let respBody = await (0, getKeyHandlerExporter_1.default)(event, null, null);
    resp.set("Content-Type", "application/json");
    return resp.status(respBody.statusCode).json(respBody.body);
});
app.get("/test", (req, resp) => {
    return resp.status(200).json(JSON.stringify({ data: "ok" }));
});
app.listen(8000, () => console.info("Running on PORT 8000"));
