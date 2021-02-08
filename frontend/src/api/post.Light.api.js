import {postRequest} from "./requests";

export function setLight(id, hsb) {

    return postRequest("/lamps/" + id, {
        on: true,
        hsb: hsb
    })
}