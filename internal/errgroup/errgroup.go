package errgroup

// NOTE: This is sample code.

type Group interface {
	Do()
}

type group struct{}

func New() Group {
	return new(group)
}

func (g *group) Do() {
}
