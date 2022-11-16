package color

type Color uint8

const (
	Red Color = iota
	Green
	Blue
	Purple
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
