import {postRequest} from "./requests";

export function postGammaAuth(code) {
    return postRequest("/auth", {code: code});
}