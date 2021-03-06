package lru

type Options struct {
	Size          int
	Lookup        LookupFunc
	MultiThreaded bool
}

type Option func(*Options)

const (
	DefaultSize = 50
)

func WithLookup(fn LookupFunc) Option {
	return func(o *Options) {
		o.Lookup = fn
	}
}

func WithSize(s int) Option {
	return func(o *Options) {
		o.Size = s
	}
}

func WithMultiThreaded(o *Options) {
	o.MultiThreaded = true
}

func newDefaultOptions() *Options {
	return &Options{
		Size: DefaultSize,
	}
}
