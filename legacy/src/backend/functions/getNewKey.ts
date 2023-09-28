import { app } from "firebase-admin";
import getNewestKey from "../helper/apikey/getNewestKey";
import appCheck from "../helper/auth/appCheck";
import DetaSingleton from "../helper/shared/DetaSingleton";
import { FunctionEvent } from "../helper/types/functionEvent";

export async function handler(event: FunctionEvent, context: any, callback: any) {
    try {
        await DetaSingleton.Instance.init();
        //TODO: Find out why X-AppCheck is converted to X-Appcheck! something seems broken
        let isDeviceValid = await appCheck(event.headers["X-Appcheck"] ?? "No token", DetaSingleton.Instance);
        if (!isDeviceValid)
            return {
                statusCode: 401,
                body: "Invalid device",
                headers: { "Content-Type": "text/plain" }
            }

        const deta = DetaSingleton.Instance.deta!;
        let newestKey = await getNewestKey(deta);
        console.info("Request successfulls")
        return { statusCode: 200, headers:{"Content-Type":"application/json"} ,body: JSON.stringify({ key: newestKey }) }
    } catch (e) {
        console.error("Something went wrong: " + e)
        return {
            statusCode: 500, body: "Something went wrong", headers: { "Content-Type": "text/plain" }
        }
    }
}