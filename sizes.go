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
)

// Standards maps standard names to instances of standards
var Standards = map[string]*Standard{
	"ISO-126": StandardISO216,
	"ISO-269": StandardISO269,
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
}
