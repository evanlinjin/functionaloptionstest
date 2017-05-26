package hello

import "github.com/pkg/errors"

type OptFunc func(g *Greeter) error

func Greeting(greeting string) OptFunc {
	return func(g *Greeter) error {
		if g == nil {
			return errors.New("empty greeter")
		}
		g.greeting = greeting
		return nil
	}
}

func Name(name string) OptFunc {
	return func(g *Greeter) error {
		if g == nil {
			return errors.New("empty greeter")
		}
		g.name = name
		return nil
	}
}

func Count(count int) OptFunc {
	return func(g *Greeter) error {
		if g == nil {
			return errors.New("empty greeter")
		}
		if count < 0 {
			return errors.New("count can not be negative")
		}
		g.count = count
		return nil
	}
}