package iotmaker_platform_coordinate

import (
	commonTypes "github.com/helmutkemper/iotmaker.santa_isabel_theater.commonTypes"
)

type IDensity interface {
	Set(value commonTypes.Number)
	SetInt(value commonTypes.Number)
	Add(value commonTypes.Number)
	Sub(value commonTypes.Number)
	SetDensityFactor(value commonTypes.Number)
	Int() commonTypes.Number
	Float64() commonTypes.Number
	Float32() commonTypes.Number
	Float() commonTypes.Number
	String() string
}
