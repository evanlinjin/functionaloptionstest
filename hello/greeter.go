package hello

import "fmt"

// Greeter represents a greeter.
type Greeter struct {
	greeting string
	name     string
	count    int
}

// NewGreeter creates a new Greeter.
func NewGreeter(opts ...func(*Greeter) error) (*Greeter, error) {
	g := &Greeter{
		greeting: "Hello",
		name: "Anonymous",
		count: 1,
	}
	for _, opt := range opts {
		if e := opt(g); e != nil {
			return nil, e
		}
	}
	return g, nil
}

// Greet greets like a boss.
func (g *Greeter) Greet() {
	for i := 0; i < g.count; i++ {
		fmt.Printf("%s %s!\n", g.greeting, g.name)
	}
}