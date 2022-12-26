import {appCheck}  from "firebase-admin";

class AppcheckWrapper {
    async verifyToken(token: string): Promise<boolean> {
        try {
            console.log(token)
            let isValid = await appCheck().verifyToken(token);
            console.log(isValid);
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