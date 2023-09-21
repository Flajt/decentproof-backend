import Deta from "deta/dist/types/deta";

async function checkForNewKey(key:string,deta:Deta):Promise<boolean>{
    try{
        const db = deta.Base("keys")
        let currentKey = await db.get(key)
        if(currentKey){
            let expire = currentKey!.__expires as number;
            let keys = await db.fetch();
            let hasNewer = false;
            for(let key of keys.items){
                if(new Date(key.__expires! as number) > new Date(expire)){
                    hasNewer = true;
                    break;
                }
            }
            if(hasNewer)
                return true;
        }
        return true;
          
    }catch(e){
        return true;
    }
}
export default checkForNewKey;