import checkForNewKey from "../helper/apikey/checkForNewKey";
import DetaSingleton from "../helper/shared/DetaSingleton";
import { FunctionEvent } from "../helper/types/functionEvent";

export async function handler(event: FunctionEvent, context: any, callback: any) {
    try {
        await DetaSingleton.Instance.init();
        const deta = DetaSingleton.Instance.deta;
        let key = (event.headers as any)["Authorization"].split(" ")[1];
        let hasNewer = await checkForNewKey(key ?? "no key", deta!);
        let body = {
            statusCode: 200,
            headers:{"Content-Type":"application/json"},
            body: {hasNewer: hasNewer},
        };
        console.info("Request successfull")
        return body;

    } catch (e) {
        console.error("Something went wrong: " + e)
        return {
            headers: {"Content-Type":"text/plain"},
            statusCode: 500,
            body: "Something went wrong",
        }
    }

}