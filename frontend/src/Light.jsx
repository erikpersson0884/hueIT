import "./App.css"
import React from "react"

export const Light = props => {
    const {color} = props
    return (
        <div className="Light" style={{
            backgroundColor: color
        }}/>
    )
}