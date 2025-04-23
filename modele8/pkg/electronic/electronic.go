package electronic

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type Smartphone interface {
	OS() string
}

type StationPhone interface {
	ButtonsCount() int
}

////////// apple //////////

type applePhone struct {
	model string
}

func NewApplePhone(model string) applePhone {
	return applePhone{
		model: model,
	}
}

func (a *applePhone) Brand() string {
	return "apple"
}

func (a *applePhone) Model() string {
	return a.model
}

func (a *applePhone) Type() string {
	return "smartphone"
}

func (a *applePhone) OS() string {
	return "ios"
}

////////// android //////////

type androidPhone struct {
	brand string
	model string
}

func NewAndroidPhone(brand, model string) androidPhone {
	return androidPhone{
		brand: brand,
		model: model,
	}
}

func (a androidPhone) Brand() string {
	return a.brand
}

func (a androidPhone) Model() string {
	return a.model
}

func (a androidPhone) Type() string {
	return "smartphone"
}

func (a androidPhone) OS() string {
	return "android"
}

////////// radioPhone //////////

type radioPhone struct {
	brand   string
	model   string
	buttons int
}

func NewRadioPhone(brand, model string, buttons int) radioPhone {
	return radioPhone{
		brand:   brand,
		model:   model,
		buttons: buttons,
	}
}

func (r radioPhone) Brand() string {
	return r.brand
}

func (r radioPhone) Model() string {
	return r.model
}

func (r radioPhone) Type() string {
	return "station"
}

func (r radioPhone) ButtonsCount() int {
	return r.buttons
}
