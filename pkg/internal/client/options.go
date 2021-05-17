package client

// Option defines a single option function.
type Option func(o *Options)

// Options defines the available options for this package.
type Options struct {
	Username string
	Password string
}

// newOptions initializes the available default options.
func newOptions(opts ...Option) Options {
	opt := Options{}

	for _, o := range opts {
		o(&opt)
	}

	return opt
}

// WithUsername provides a function to set the username option.
func WithUsername(val string) Option {
	return func(o *Options) {
		o.Username = val
	}
}

// WithPassword provides a function to set the password option.
func WithPassword(val string) Option {
	return func(o *Options) {
		o.Password = val
	}
}
