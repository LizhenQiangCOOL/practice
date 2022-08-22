package utils

import "log"

var (
	_closers []Closer
)

type Closer func() error

func AppendToClosers(c Closer) {
	_closers = append(_closers, c)
}

func CloserAll() {
	for i := range _closers {
		if err := _closers[i](); err != nil {
			log.Panicln(err)
		}
	}
}
