package gobf

import "fmt"

type Option func(*bruteforce) error

var (
	defaultOpts = []Option{
		WithConcrencyLimit(500000),
	}
)

func WithConcrencyLimit(limit int) Option {
	return func(bf *bruteforce) error {
		if limit < 0 {
			return fmt.Errorf("invalid size: %d", limit)
		}

		bf.limit = limit
		return nil
	}
}

func WithSize(size int) Option {
	return func(bf *bruteforce) error {
		if size < 0 {
			return fmt.Errorf("invalid size: %d", size)
		}

		bf.size = size
		return nil
	}
}

func WithNumber(enabled bool) Option {
	return func(bf *bruteforce) error {
		bf.numEnabled = enabled
		return nil
	}
}

func WithUpper(enabled bool) Option {
	return func(bf *bruteforce) error {
		bf.upperEnabled = enabled
		return nil
	}
}

func WithLower(enabled bool) Option {
	return func(bf *bruteforce) error {
		bf.lowerEnabled = enabled
		return nil
	}
}
