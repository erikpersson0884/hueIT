import axios from "axios";

let path = "/api"

export function getRequest(endpoint) {
    let headers = {};
    return axios.get(path + endpoint, {headers});
}

export function postRequest(endpoint, data) {
    let headers = {};
    return axios.post(path + endpoint, data, {headers});
}

export function putRequest(endpoint, data) {
    let headers = {};
    return axios.put(path + endpoint, data, {headers});
}