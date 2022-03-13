package compton

type Options struct {
	DatabasePath string
}

func NewOptions() *Options {
	return &Options{
		DatabasePath: "./",
	}
}
