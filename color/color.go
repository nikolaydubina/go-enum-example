package color

type Color struct{ c uint }

var (
	Undefined = Color{}
	Red       = Color{1}
	Green     = Color{2}
	Blue      = Color{3}
	Purple    = Color{4}
)

func (c Color) String() string {
	switch c {
	case Red:
		return "red"
	case Green:
		return "green"
	case Blue:
		return "blue"
	case Purple:
		return "purple"
	default:
		return "unknown"
	}
}
