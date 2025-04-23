package main

import (
	"fmt"
	"module8/pkg/electronic"
)

func printCharacteristics(p electronic.Phone) string {
	if smart, ok := p.(electronic.Smartphone); !ok {
		if station, ok := p.(electronic.StationPhone); !ok {
			return "invalid type assertion"
		} else {
			return fmt.Sprintf("Brand: %s \nModel: %s\nType: %s\nButtons: %d", p.Brand(), p.Model(), p.Type(), station.ButtonsCount())
		}
	} else {
		return fmt.Sprintf("Brand: %s \nModel: %s\nType: %s\nOS: %s", p.Brand(), p.Model(), p.Type(), smart.OS())
	}
}

func main() {

	iphone := electronic.NewApplePhone("iphone 16")
	android := electronic.NewAndroidPhone("samsung", "s21 ultra")
	station := electronic.NewRadioPhone("panasonic", "KX-TGJ320", 23)

	fmt.Println(printCharacteristics(&iphone), "\n ")
	fmt.Println(printCharacteristics(&android), "\n ")
	fmt.Println(printCharacteristics(&station))

}
