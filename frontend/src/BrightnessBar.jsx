import React, {useMemo, useRef} from "react"
import {toHslString} from "./utility";
import {toColor} from "react-color-palette";

// Code copied / modified from https://github.com/Wondermarin/react-color-palette/
export const BrightnessBar = props => {
    const {width, color, onChange} = props
    const brightnessRef = useRef(null);
    const position = useMemo(() => {
        return getCoordinatesByBrightness(color.hsb.b, width);
    }, [color.hsb.b, width]);

    const moveCursor = (x, shiftX) => {
        const [newX] = moveAt({
                                  value: x,
                                  shift: shiftX,
                                  min: 0,
                                  max: width
                              });

        const newB = getBrightnessByCoordinates(newX, width);

        onChange(toColor("hsb", {...color.hsb, b: newB}));
    };

    const onMouseDown = (e) => {
        if (brightnessRef.current) {
            if (e.button !== 0) return;

            const { current: hue } = brightnessRef;

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
    <div ref={brightnessRef} className="BrightnessBar" onMouseDown={onMouseDown} style={{backgroundImage: brightnessColors(color.hsb)}}>
        <div className="BrightnessBarCursor" style={{ left: position, backgroundColor: toHslString(color.hsb) }} />
    </div>
    );
};

function getCoordinatesByBrightness(h, width) {
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

export function getBrightnessByCoordinates(x, width) {
    return (x / width) * 100;
}


function brightnessColors(hsb) {
    const zero = {
        h: hsb.h,
        s: hsb.s,
        b: 0
    }

    const color = {
        h: hsb.h,
        s: hsb.s,
        b: 50
    }

    return `linear-gradient( to left, ${toHslString(color)}, ${toHslString(zero)})`
}
