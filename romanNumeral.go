package lab1

const (
	I romanSymbol = 1
	V romanSymbol = 5
	X romanSymbol = 10
	L romanSymbol = 50
	C romanSymbol = 100
	D romanSymbol = 500
	M romanSymbol = 1000
)

type (
	romanSymbol  int16
	RomanNumeral []romanSymbol
)

func (n RomanNumeral) String() string {
	var str string
	for _, symbol := range n {
		switch symbol {
		case I:
			str += "I"
		case V:
			str += "V"
		case X:
			str += "X"
		case L:
			str += "L"
		case C:
			str += "C"
		case D:
			str += "D"
		case M:
			str += "M"
		}
	}
	return str
}
