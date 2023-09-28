"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const console_1 = require("console");
class ApiKeyModel {
    constructor(keyProps, jsonProps) {
        if (keyProps) {
            this.expirationDate = keyProps.expirationDate;
            this.apiKey = keyProps.key;
        }
        else {
            this.apiKey = jsonProps.key;
            this.expirationDate = new Date(jsonProps.__expire);
        }
        (0, console_1.assert)(keyProps != undefined || jsonProps != undefined);
    }
    toJson() {
        return {
            key: this.apiKey,
            __expire: this.expirationDate.toISOString()
        };
    }
}
exports.default = ApiKeyModel;
