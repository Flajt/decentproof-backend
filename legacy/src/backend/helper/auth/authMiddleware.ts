import verifyApiKey from "../apikey/verifyApikey";
import DetaSingleton from "../shared/DetaSingleton";

export default async function authCheck(headers:any){
    if(process.env.NODE_ENV==="production"){
        await DetaSingleton.Instance.init();
        const deta = DetaSingleton.Instance.deta;
        let authHeader:string = headers["Authorization"];
        if(authHeader!== null && authHeader!==undefined && authHeader.length>5){
            let apiKey = authHeader.split(" ")[1];
            let isValid = await verifyApiKey(apiKey,deta!)
            if(isValid)
                return true;
            return false
        }else{
            return false;
        }
    }else{
        return true;
    }
}

