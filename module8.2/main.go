package main

import (
	"fmt"

	"module8.2/pkg/auto"
	"module8.2/pkg/dimension"
)

func main() {

	BMW := auto.NewCar("BMW", "X7", dimension.NewEuroSize(515.1, 200, 180.5), 250, 381)
	Mercedes := auto.NewCar("Mercedes", "W465", dimension.NewEuroSize(481.7, 193.1, 196.9), 240, 585)
	Dodge := auto.NewCar("Dodge", "Challenger 2024", dimension.NewUsaSize(197.5, 75.7, 55.7), 257, 305)

	//////////car struct//////////
	fmt.Println(Mercedes)
	fmt.Println(Dodge)
	fmt.Println()

	//////////Dimensions interface//////////
	fmt.Println(BMW.Dimensions().Length())
	fmt.Println(Mercedes.Dimensions().Width())
	fmt.Println(Dodge.Dimensions().Height())
	fmt.Println()

	//////////Get//////////
	fmt.Println(Mercedes.Dimensions().Length().Get(dimension.Inch))
	fmt.Println(Dodge.Dimensions().Height().Get(dimension.CM))
	fmt.Println()

	//////////+ GetAll//////////
	fmt.Println(dimension.GetAll(BMW.Dimensions(), dimension.Inch))
	fmt.Println(dimension.GetAll(Dodge.Dimensions(), dimension.CM))
}
