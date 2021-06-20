package iotmaker_platform_coordinate

import (
	commonTypes "github.com/helmutkemper/iotmaker.santa_isabel_theater.commonTypes"
	"strconv"
)

type Density struct {
	OriginalValue commonTypes.Number
	DensityFactor commonTypes.Number
	DensityValue  commonTypes.Number
}

func (el *Density) adjustDensity() {
	if el.DensityFactor == 0 {
		el.DensityFactor = 1
	}

	el.DensityValue = el.OriginalValue * el.DensityFactor
}

func (el *Density) Set(value commonTypes.Number) {
	el.OriginalValue = value
	el.adjustDensity()
}

func (el *Density) SetInt(value commonTypes.Number) {
	el.OriginalValue = value
	el.adjustDensity()
}

func (el *Density) Add(value commonTypes.Number) {
	el.OriginalValue += value
	el.adjustDensity()
}

func (el *Density) Sub(value commonTypes.Number) {
	el.OriginalValue -= value
	el.adjustDensity()
}

func (el *Density) SetDensityFactor(value commonTypes.Number) {

	//switch converted := value.(type) {
	//case float64:
	//	el.DensityFactor = int(converted)
	//case int:
	//	el.DensityFactor = converted
	//}

	el.DensityFactor = value
	el.adjustDensity()
}

func (el Density) Int() commonTypes.Number {
	return el.DensityValue
}

func (el Density) Float64() commonTypes.Number {
	return el.DensityValue
}

func (el Density) Float32() commonTypes.Number {
	return el.DensityValue
}

func (el Density) Float() commonTypes.Number {
	return el.DensityValue
}

func (el Density) String() string {
	return strconv.FormatInt(int64(el.OriginalValue), 10)
}
