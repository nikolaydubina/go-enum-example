package color

type ColorBig uint

const (
	RedBig ColorBig = iota
	GreenBig
	BlueBig
	PurpleBig
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
