import {getRequest} from "./requests";

export function getLights() {
    return getRequest("/lamps")
}