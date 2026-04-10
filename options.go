package htmltable

type options struct {}

type Option func(*options)

func applyOptions(opts []Option) options {
	o := options{}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
