package surface

import (
	"fmt"
	"net/http"
	"strconv"
)

// 指定パラメータ
const (
	keyModel   = "modelType"
	keyColor   = "color"
	keyWidth   = "width"
	keyHeight  = "height"
	keyCells   = "cells"
	keyXYRange = "xyrange"
	keyZScale  = "zscale"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	for key, values := range r.URL.Query() {
		switch key {
		case keyModel:
			setModel(w, values[0])
		case keyColor:
			setColor(w, values[0])
		case keyWidth, keyHeight, keyCells, keyXYRange, keyZScale:
			setVal(w, key, values[0])
		}
	}
	PrintXML(w)
}

func setModel(w http.ResponseWriter, val string) {
	switch val {
	case "EggCase":
		Model = EGGCASE
	case "Moguls":
		Model = MOGULS
	case "Saddle":
		Model = SADDLE
	default:
		Model = DEFAULT
	}
}

func setColor(w http.ResponseWriter, val string) {
	if val == "gradient" {
		IsGradientColor = true
	} else {
		FillColor = val
		IsGradientColor = false
	}
}

func setVal(w http.ResponseWriter, key string, val string) {
	param, err := strconv.Atoi(val)
	if err != nil {
		fmt.Fprintf(w, "invalid value: %v\n", val)
	}

	switch key {
	case keyWidth:
		Width = float64(param)
	case keyHeight:
		Height = float64(param)
	case keyCells:
		Cells = param
	case keyXYRange:
		XYRange = float64(param)
	case keyZScale:
		ZScale = float64(param)
	}
}
