import convert from "color-convert"

export function toHslString(hsb) {
    // let hsl = hsbToHsl(hsb)
    let hsl = convert.hsv.hsl(hsb.h, hsb.s, hsb.b)
    // console.log("HSB", hsb, "HSL", hsl)
    return `hsl(${hsl[0]}, ${hsl[1]}%, ${hsl[2]}%)`
}

function hsbToHsl(hsb) {
    let ll = (2 - hsb.s) * hsb.b
    let ss = hsb.s * hsb.b
    ss /= ll <= 1 ? ll : 2 - ll
    ll /= 2

    return {
        h: hsb.h,
        s: ss,
        l: ll
    }
}

export function toHsbString(hsb) {
    return `hsl(${hsb.h}, ${hsb.s}%, ${hsb.b}%)`
}

export function toRgbString(rgb) {
    return `rgb(${rgb.r}, ${rgb.g}, ${rgb.b})`
}

let setLoginUrl = _ => {console.log("Error?")}

export function initializeSetLoginUrl(setLoginUrlFunc) {
    setLoginUrl = setLoginUrlFunc
}

export async function AuthorizedApiCall(call) {
    return call()
    .then(response => {
        return {
            error: false,
            response: response
        }
    })
    .catch(error => {
        const headers = error.response.headers;
        if (error.response.status === 401 && headers.location) {
            setLoginUrl(headers.location)
        }

        return {
            error: true,
            errResponse: error
        }
    })
}