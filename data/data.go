package data

// DataType is the type of value in Data
type DataType map[string][]string

// Data consists of the main set of fake information
var Data = map[string]DataType{
	"animal":    Animals,
}