import './App.css';
import ColorPicker, {useColor} from "react-color-palette";
import React, {useEffect, useState} from "react";
import {Light} from "./Light";
import {getLights} from "./api/get.Lights.api";
import {toHslString} from "./utility";
import {BrightnessBar} from "./BrightnessBar";
import {setLight} from "./api/post.Light.api";
import {setAllLights} from "./api/post.AllLights.api";
import {SaturationBar} from "./SaturationBar";

const defaultLights = [
  {
    id: 0,
    x: 0,
    y: 0,
    color: "#ffffff"
  },
  {
    id: 1,
    x: 1,
    y: 0,
    color: "#ffffff"
  },
  {
    id: 2,
    x: 0,
    y: 1,
    color: "#ffffff"
  },
  {
    id: 5,
    x: 1,
    y: 1,
    color: "#ffffff"
  },
  {
    id: 7,
    x: 0,
    y: 2,
    color: "#ffffff"
  },
  {
    id: 11,
    x: 1,
    y: 2,
    color: "#ffffff"
  },
  {
    id: 14,
    x: 0,
    y: 3,
    color: "#ffffff"
  },
  {
    id: 99,
    x: 1,
    y: 3,
    color: "#ffffff"
  }
]

function App() {
  const [color, setColor] = useColor("hex", "#ffffff")
  const [lights, setLights] = useState(defaultLights)
  const [error, setError] = useState(undefined)

  useEffect(() => {
    getLights()
    .then(response => {
      setLights(response.data.lights.map(obj => {
        const l = obj.light
        const s = obj.state
        return {
          id: l.id,
          x: l.x,
          y: l.y,
          color: {
            hsb: {
              h: s.h,
              s: s.s,
              b: s.b
            }
          }
        }
      }))
    })
    .catch(error => {
      console.error("Failed to retrieve lights, error: ", error)
      setError("Something went wrong :(")
    })
  }, [])

  return (
    <div className="App">
      <div className="Row">
        <div>
          <ColorPicker width={400} height={400} color={color} onChange={setColor}/>
          {color.hsb && (
          <div className="BarsContainer">
            <BrightnessBar width={400-32} color={color} onChange={setColor}/>
            <SaturationBar width={400-32} color={color} onChange={setColor}/>
          </div>
          )}
          <button className="SetAllButton" onClick={() => setLights(updateAllLights(lights, color, setError))}>
            SET ALL
          </button>
        </div>
        { error ? (
        <div className="ErrorContainerContainer">
          <div className="ErrorContainer">
            <p className="ErrorText">{error}</p>
          </div>
        </div>
        ) : (
        <div className="Fixed">
          { lights.map((light, index) => (
            <button className="LightContainer" style={{
              left: light.x * 100 + "px",
              top: light.y * 100 + "px",
            }} key={index} onClick={() => {
              setLights(updateLight(light.id, lights, color, setError))
            }}>
              <Light color={light.color}/>
            </button>
          ))}
        </div>
        )}
      </div>
    </div>
  );
}

function updateAllLights(lights, color, setError) {
  setAllLights(color.hsb)
  .catch(error => {
    console.log("failed to set all lights due to error: ", error)
    setError("Failed to set lights, please reload the page and try again")
  })
  return lights.map(light => {
    return {
      ...light,
      color: color
    }
  })
}

function updateLight(id, lights, color, setError) {
  setLight(id, color.hsb)
  .catch(error => {
    console.log("failed to set light due to error: ", error)
    setError("Failed to set light, please reload the page and try again.")
  })

  return lights.map(light => {
    if (light.id === id) {
      return {
        ...light,
        color: color
      }
    }
    return light
  })
}

export default App;
