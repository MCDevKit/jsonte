package functions

type Docs struct {
	Summary   string
	Arguments []Argument
	Example   string
}

type Argument struct {
	Name     string
	Summary  string
	Optional bool
}
