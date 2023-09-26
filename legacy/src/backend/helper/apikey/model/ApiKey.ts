import { assert } from "console";
import { ApiKeyT } from "../../types/key";

class ApiKeyModel{
    public apiKey:string;
    public expirationDate:Date;

    constructor(keyProps:ApiKeyT | undefined, jsonProps:{key:string,__expire:string}| undefined){
        if(keyProps){
        this.expirationDate = keyProps.expirationDate;
        this.apiKey = keyProps.key;
        }else {
            this.apiKey = jsonProps!.key;
            this.expirationDate = new Date(jsonProps!.__expire)
        }
        assert(keyProps!=undefined || jsonProps !=undefined)
    }
    
    toJson():{key:string,__expire:string}{
        return {
            key: this.apiKey,
            __expire: this.expirationDate.toISOString()
        }
    }
}

export default ApiKeyModel;