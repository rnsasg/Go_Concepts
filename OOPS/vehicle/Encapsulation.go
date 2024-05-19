package vehicle

type Color int

const (
	UNKNOWN = iota
	RED
	BLACK
	BLUE
)

type vechile struct {
	noOfWheel int
	color     Color
}

func NewVechile(noOfWheel int, color Color) *vechile {
	return &vechile{noOfWheel: noOfWheel, color: color}
}

func (v *vechile) GetNoOfWheel() int {
	return v.noOfWheel
}
func (v *vechile) SetNoOfWheel(noOfWheel int) {
	v.noOfWheel = noOfWheel
}
