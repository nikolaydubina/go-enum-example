package color

type ColorBig string

const (
	RedBig    ColorBig = "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfas-1"
	GreenBig  ColorBig = "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfas-2"
	BlueBig   ColorBig = "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfas-3"
	PurpleBig ColorBig = "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfas-4"
)

func (c ColorBig) String() string { return string(c) }
