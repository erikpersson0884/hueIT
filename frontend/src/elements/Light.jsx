import "../App.css"
import React from "react"
import {toHslString} from "../utility";

export const Light = props => {
    let hsb = {
        h: 0,
        s: 100,
        b: 100
    }
    if (props.color.hsb) {
        hsb = props.color.hsb
    }

    return (
        <div className="Light" style={{
            backgroundColor: toHslString(hsb)
        }}/>
    )
}