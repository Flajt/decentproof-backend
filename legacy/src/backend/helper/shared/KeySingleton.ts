import Drive from "deta/dist/types/drive";
import { KJUR } from "jsrsasign";

//Based on: https://stackoverflow.com/questions/30174078/how-to-define-singleton-in-typescript#36978360
class KeySingleton
{
    private static _instance: KeySingleton;
    //Looks like an odd workaround, TODO: Find way of creating it without use of null
    pubKey: KJUR.crypto.ECDSA | null;
    privKey: KJUR.crypto.ECDSA | null;

    private constructor(){
        this.privKey = null;
        this.pubKey = null;
    }
    public async init(drive:Drive):Promise<void>{
        let privResp = await drive.get("priv.pem")
        let pubResp = await drive.get("pub.pem")
        let privBuff = await privResp?.arrayBuffer();
        let pubBuff = await pubResp?.arrayBuffer();
        this.privKey = new KJUR.crypto.ECDSA({curve:"secp256r1", prv: Buffer.from(privBuff!).toString("utf-8")})
        this.pubKey = new KJUR.crypto.ECDSA({curve:"secp256r1",pub: Buffer.from(pubBuff!).toString("utf-8")}) 
    }

    public static get Instance()
    {
        return this._instance || (this._instance = new this());
    }
}


export default KeySingleton;