// Packer papersizes provides mapping between many paper size standards
// and their dimensions.
// Wikipedia article https://en.wikipedia.org/wiki/Paper_size has been
// the primary source for this package.

package papersizes

type Standard struct {
	// Name of the standard
	Name string
}

var (
	StandardISO216 = &Standard{"ISO-126"}
	StandardISO269 = &Standard{"ISO-269"}
	StandardDIN476 = &Standard{"DIN 476"}
)

// Standards maps standard names to instances of standards
var Standards = map[string]*Standard{
	"ISO-126": StandardISO216,
	"ISO-269": StandardISO269,
	"DIN 476": StandardDIN476,
}

// Size describes the attributes of a paper size
type Size struct {
	// Canonical name as defined by the relevant standard
	Name string

	// Standard defining this paper size
	Standard *Standard

	// Width of the paper in exact millimeters
	Width int

	// Height of the paper in exact millimeters
	Height int
}

var (
	SizeA0  = &Size{"A0", StandardISO216, 841, 1189}
	SizeA1  = &Size{"A1", StandardISO216, 594, 841}
	SizeA2  = &Size{"A2", StandardISO216, 420, 594}
	SizeA3  = &Size{"A3", StandardISO216, 297, 420}
	SizeA4  = &Size{"A4", StandardISO216, 210, 297}
	SizeA5  = &Size{"A5", StandardISO216, 148, 210}
	SizeA6  = &Size{"A6", StandardISO216, 105, 148}
	SizeA7  = &Size{"A7", StandardISO216, 74, 105}
	SizeA8  = &Size{"A8", StandardISO216, 52, 74}
	SizeA9  = &Size{"A9", StandardISO216, 37, 52}
	SizeA10 = &Size{"A10", StandardISO216, 26, 37}
	SizeA11 = &Size{"A11", StandardDIN476, 18, 26}
	SizeA12 = &Size{"A12", StandardDIN476, 13, 18}
	SizeA13 = &Size{"A13", StandardDIN476, 9, 13}

	SizeB0  = &Size{"B0", StandardISO216, 1000, 1414}
	SizeB1  = &Size{"B1", StandardISO216, 707, 1000}
	SizeB2  = &Size{"B2", StandardISO216, 500, 707}
	SizeB3  = &Size{"B3", StandardISO216, 353, 500}
	SizeB4  = &Size{"B4", StandardISO216, 250, 353}
	SizeB5  = &Size{"B5", StandardISO216, 176, 250}
	SizeB6  = &Size{"B6", StandardISO216, 125, 176}
	SizeB7  = &Size{"B7", StandardISO216, 88, 125}
	SizeB8  = &Size{"B8", StandardISO216, 62, 88}
	SizeB9  = &Size{"B9", StandardISO216, 44, 62}
	SizeB10 = &Size{"B10", StandardISO216, 31, 44}
	SizeB11 = &Size{"B11", StandardDIN476, 22, 31}
	SizeB12 = &Size{"B12", StandardDIN476, 15, 22}
	SizeB13 = &Size{"B13", StandardDIN476, 11, 15}

	SizeC0  = &Size{"C0", StandardISO269, 917, 1297}
	SizeC1  = &Size{"C1", StandardISO269, 648, 917}
	SizeC2  = &Size{"C2", StandardISO269, 458, 648}
	SizeC3  = &Size{"C3", StandardISO269, 324, 458}
	SizeC4  = &Size{"C4", StandardISO269, 229, 324}
	SizeC5  = &Size{"C5", StandardISO269, 162, 229}
	SizeC6  = &Size{"C6", StandardISO269, 114, 162}
	SizeC7  = &Size{"C7", StandardISO269, 81, 114}
	SizeC8  = &Size{"C8", StandardISO269, 57, 81}
	SizeC9  = &Size{"C9", StandardISO269, 40, 57}
	SizeC10 = &Size{"C10", StandardISO269, 28, 40}

	SizeD0 = &Size{"D0", StandardDIN476, 771, 1090}
	SizeD1 = &Size{"D1", StandardDIN476, 545, 771}
	SizeD2 = &Size{"D2", StandardDIN476, 385, 545}
	SizeD3 = &Size{"D3", StandardDIN476, 272, 385}
	SizeD4 = &Size{"D4", StandardDIN476, 192, 272}
	SizeD5 = &Size{"D5", StandardDIN476, 136, 192}
	SizeD6 = &Size{"D6", StandardDIN476, 96, 136}
	SizeD7 = &Size{"D7", StandardDIN476, 68, 96}
	SizeD8 = &Size{"D8", StandardDIN476, 48, 68}

	Size4A0 = &Size{"4A0", StandardDIN476, 1682, 2378}
	Size2A0 = &Size{"2A0", StandardDIN476, 1189, 1682}
)

// Sizes maps a paper size name to one or more Size instances
// The first slice element is the most common paper size
// with that name in case of multiple ones.
var Sizes = map[string][]*Size{
	"A0":  {SizeA0},
	"A1":  {SizeA1},
	"A2":  {SizeA2},
	"A3":  {SizeA3},
	"A4":  {SizeA4},
	"A5":  {SizeA5},
	"A6":  {SizeA6},
	"A7":  {SizeA7},
	"A8":  {SizeA8},
	"A9":  {SizeA9},
	"A10": {SizeA10},
	"A11": {SizeA11},
	"A12": {SizeA12},
	"A13": {SizeA13},

	"B0":  {SizeB0},
	"B1":  {SizeB1},
	"B2":  {SizeB2},
	"B3":  {SizeB3},
	"B4":  {SizeB4},
	"B5":  {SizeB5},
	"B6":  {SizeB6},
	"B7":  {SizeB7},
	"B8":  {SizeB8},
	"B9":  {SizeB9},
	"B10": {SizeB10},
	"B11": {SizeB11},
	"B12": {SizeB12},
	"B13": {SizeB13},

	"C0":  {SizeC0},
	"C1":  {SizeC1},
	"C2":  {SizeC2},
	"C3":  {SizeC3},
	"C4":  {SizeC4},
	"C5":  {SizeC5},
	"C6":  {SizeC6},
	"C7":  {SizeC7},
	"C8":  {SizeC8},
	"C9":  {SizeC9},
	"C10": {SizeC10},

	"D0": {SizeD0},
	"D1": {SizeD1},
	"D2": {SizeD2},
	"D3": {SizeD3},
	"D4": {SizeD4},
	"D5": {SizeD5},
	"D6": {SizeD6},
	"D7": {SizeD7},
	"D8": {SizeD8},

	"4A0": {Size4A0},
	"2A0": {Size2A0},
}
