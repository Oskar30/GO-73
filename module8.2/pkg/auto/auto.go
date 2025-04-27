package auto

import (
	"module8.2/pkg/dimension"
)

type Auto interface {
	Brand() string
	Model() string
	Dimensions() dimension.Dimensions
	MaxSpeed() int
	EnginePower() int
}

type car struct {
	brand       string
	model       string
	dimensions  dimension.Dimensions
	maxSpeed    int
	enginePower int
}

func NewCar(brand, model string, dimensions dimension.Dimensions, maxSpeed, enginePower int) car {
	return car{
		brand:       brand,
		model:       model,
		dimensions:  dimensions,
		maxSpeed:    maxSpeed,
		enginePower: enginePower,
	}
}

func (c car) Brand() string {
	return c.brand
}

func (c car) Model() string {
	return c.model
}

func (c car) Dimensions() dimension.Dimensions {
	return c.dimensions
}

func (c car) MaxSpeed() int {
	return c.maxSpeed
}

func (c car) EnginePower() int {
	return c.enginePower
}
