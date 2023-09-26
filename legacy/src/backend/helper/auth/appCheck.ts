import { initializeApp, getApps } from "firebase-admin/app";
import AppcheckWrapper from "../appcheck/appcheck";
import { cert } from "firebase-admin/app";
import DetaSingleton from "../shared/DetaSingleton";

export default async function appCheck(appCheckToken: string,detaSingleton:DetaSingleton) {
    if (process.env.NODE_ENV === "production") {
         getApps().length === 0 ? initializeApp({
            credential: cert(detaSingleton.fireBaseKey!)
        }) : null;
        const appcheckWrapper = new AppcheckWrapper();
        if (!appCheckToken || appCheckToken.length <= 1) {
            console.info("no token")
            return false;
        } else {
            let isValid = await appcheckWrapper.verifyToken(appCheckToken);
            if (isValid)
                return true;
            else
                return false;
        }
    } else {
        return true;
    }
}
