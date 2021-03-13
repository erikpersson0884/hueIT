import {postRequest} from "./requests";

export function postGammaLogout() {
    return postRequest("/logout", {});
}