import {Deta as d,app as detaApp} from "deta";
import Deta from "deta/dist/types/deta";
import Drive from "deta/dist/types/drive";

//Based on: https://stackoverflow.com/questions/30174078/how-to-define-singleton-in-typescript#36978360
class DetaSingleton
{
    private static _instance: DetaSingleton;
    deta:Deta | undefined;
    keys:Drive | undefined;
    fireBaseKey:object | undefined;

    public async init():Promise<void>{
        this.deta = d(process.env.PROJECT_KEY);
        this.keys = this.deta.Drive("keys");
        let firebaseKeyBuf = await this.keys.get("firebaseKey.json");
        let buf = await firebaseKeyBuf?.arrayBuffer();
        this.fireBaseKey = JSON.parse(Buffer.from(buf!).toString("utf-8"));
    }

    public static get Instance()
    {  
        return this._instance || (this._instance = new this());
    }
}


export default DetaSingleton;