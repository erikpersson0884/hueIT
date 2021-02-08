import {postRequest} from "./requests";

export function setAllLights(hsb) {
    return postRequest("/lamps", {
        on: true,
        hsb: hsb
    })
}