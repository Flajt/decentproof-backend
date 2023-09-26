import { KJUR } from "jsrsasign";
import authCheck from "../helper/auth/authMiddleware";
import DetaSingleton from "../helper/shared/DetaSingleton";
import KeySingleton from "../helper/shared/KeySingleton";
import { FunctionEvent } from "../helper/types/functionEvent";

export async function handler(event: FunctionEvent, context: any, callback: any) {
    try {
        let isValid = await authCheck(event.headers);
        if (!isValid){
            console.info("Invalid API key")
            return {
                statusCode: 401,
                body: "Access Denied",
                headers:{"Content-Type":"text/plain"}
            }
        }
        await DetaSingleton.Instance.init();
        await KeySingleton.Instance.init(DetaSingleton.Instance.keys!);
        const singleton = KeySingleton.Instance;
        const signatureManagerPriv = new KJUR.crypto.Signature({ alg: "SHA224withECDSA" })
        signatureManagerPriv.init(singleton.privKey!)
        let body: any = JSON.parse(event.body as string);
        signatureManagerPriv.updateString(body.data)
        let signature = signatureManagerPriv.sign()
        let respBody =
        {
            statusCode: 200,
            headers:{"Content-Type":"application/json"},
            body: JSON.stringify({ "signature": signature }),
        }
        console.info("Request successfull")
        return respBody;
    } catch (e) {
        console.error("Somehting went wrong: " + e)
        return {
            statusCode: 500,
            headers: {"Content-Type":"text/plain"},
            body: "Something went wrong",
        }
    }
}