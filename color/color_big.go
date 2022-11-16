package color

type ColorBig struct{ c uint }

var (
	UndefinedBig = ColorBig{}
	RedBig       = ColorBig{1}
	GreenBig     = ColorBig{2}
	BlueBig      = ColorBig{3}
	PurpleBig    = ColorBig{4}
)

func (c ColorBig) String() string {
	switch c {
	case RedBig:
		return "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfas-1"
	case GreenBig:
		return "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfas-2"
	case BlueBig:
		return "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfas-3"
	case PurpleBig:
		return "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfas-4"
	default:
		return "asdfasdfasdfasdfasdfasdfasdfasdfasdfasdfasasdfasdfasdfasdfasdfasdfasdfasdfasdfasdfas-5"
	}
}
