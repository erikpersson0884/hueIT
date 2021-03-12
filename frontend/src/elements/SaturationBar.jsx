import React, {useMemo, useRef} from "react"
import {toHsbString, toRgbString} from "../utility";
import {toColor} from "react-color-palette";

// Code copied / modified from https://github.com/Wondermarin/react-color-palette/
export const SaturationBar = props => {
    const {width, color, onChange} = props
    const saturationRef = useRef(null);
    const position = useMemo(() => {
        return getCoordinatesBySaturation(color.hsb.s, width);
    }, [color.hsb.s, width]);

    const moveCursor = (x, shiftX) => {
        const [newX] = moveAt({
                                  value: x,
                                  shift: shiftX,
                                  min: 0,
                                  max: width
                              });

        const newS = getSaturationByCoordinates(newX, width);

        onChange(toColor("hsb", {...color.hsb, s: newS}));
    };

    const onMouseDown = (e) => {
        if (saturationRef.current) {
            if (e.button !== 0) return;

            const { current: hue } = saturationRef;

            document.getSelection()?.empty();

            const { left: shiftX } = hue.getBoundingClientRect();

            moveCursor(e.clientX, shiftX);

            const onMouseMove = (e) => {
                moveCursor(e.clientX, shiftX);
            };

            const onMouseUp = () => {
                document.removeEventListener("mousemove", onMouseMove, false);
                document.removeEventListener("mouseup", onMouseUp, false);
            };

            document.addEventListener("mousemove", onMouseMove, false);
            document.addEventListener("mouseup", onMouseUp, false);
        }
    };

    return (
    <div ref={saturationRef} className="BrightnessBar" onMouseDown={onMouseDown} style={{backgroundImage: saturationColors(color.hsb), width: width + "px"}}>
        <div className="BrightnessBarCursor" style={{ left: position, backgroundColor: toRgbString(color.rgb) }} />
    </div>
    );
};

function getCoordinatesBySaturation(h, width) {
    return (h / 100) * width;
}

function moveAt(x, y) {
    const X = x.value - x.shift;
    const newX = X < x.min ? x.min : X > x.max ? x.max : X;

    if (y) {
        const Y = y.value - y.shift;
        const newY = Y < y.min ? y.min : Y > y.max ? y.max : Y;

        return [newX, newY];
    }

    return [newX];
}

export function getSaturationByCoordinates(x, width) {
    return (x / width) * 100;
}


function saturationColors(hsb) {
    const zero = {
        h: hsb.h,
        s: 0,
        b: hsb.b / 2
    }

    const color = {
        h: hsb.h,
        s: 100,
        b: hsb.b / 2
    }

    return `linear-gradient( to left, ${toHsbString(color)}, ${toHsbString(zero)})`
}
