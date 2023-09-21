import {appCheck}  from "firebase-admin";

class AppcheckWrapper {
    async verifyToken(token: string): Promise<boolean> {
        try {
            let isValid = await appCheck().verifyToken(token);
            if (isValid)
                return true;
            return false
        } catch (e) {
            console.error("Cant verify token due to: "+e)
            return false;
        }
    }
}
export default AppcheckWrapper;