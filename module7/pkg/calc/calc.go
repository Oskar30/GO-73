package calc

import "fmt"

type calculator struct{}

func New() calculator {
	return calculator{}
}

func (c calculator) Calculate(x, y float64, opr string) float64 {
	switch opr {
	case "+":
		return c.sum(x, y)
	case "-":
		return c.diff(x, y)
	case "*":
		return c.multi(x, y)
	case "/":
		return c.div(x, y)
	}
	return 0
}

func (c calculator) sum(x, y float64) float64 {
	return x + y
}

func (c calculator) diff(x, y float64) float64 {
	return x - y
}

func (c calculator) multi(x, y float64) float64 {
	return x * y
}

func (c calculator) div(x, y float64) float64 {
	if y == 0 {
		fmt.Println("Ошибка: деление на ноль!")
	}
	return x / y
}
