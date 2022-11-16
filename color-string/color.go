package color

type Color string

const (
	Red    Color = "red"
	Green  Color = "green"
	Blue   Color = "blue"
	Purple Color = "purple"
)

func (c Color) String() string { return string(c) }
