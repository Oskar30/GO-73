package dimension

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
	inch float64  = 2.54
	cm   float64  = 0.394
)

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		if u.T == Inch {
			return value * inch
		} else {
			return value * cm
		}
	}
	return value
}

func GetAll(d Dimensions, t UnitType) (length, width, height float64) {
	return d.Length().Get(t), d.Width().Get(t), d.Height().Get(t)
}

////////// euroSize //////////

type euroSize struct {
	length Unit
	width  Unit
	height Unit
}

func NewEuroSize(length, width, height float64) euroSize {
	return euroSize{
		length: Unit{Value: length, T: CM},
		width:  Unit{Value: width, T: CM},
		height: Unit{Value: height, T: CM},
	}
}

func (e euroSize) Length() Unit {
	return e.length
}

func (e euroSize) Width() Unit {
	return e.width
}

func (e euroSize) Height() Unit {
	return e.height
}

////////// usaSize //////////

type usaSize struct {
	length Unit
	width  Unit
	height Unit
}

func NewUsaSize(length, width, height float64) usaSize {
	return usaSize{
		length: Unit{Value: length, T: Inch},
		width:  Unit{Value: width, T: Inch},
		height: Unit{Value: height, T: Inch},
	}
}

func (u usaSize) Length() Unit {
	return u.length
}

func (u usaSize) Width() Unit {
	return u.width
}

func (u usaSize) Height() Unit {
	return u.height
}
