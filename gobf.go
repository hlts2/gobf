package gobf

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
)

type BruteForce interface {
	Do(ctx context.Context, fn func(pattern string)) error
}

type bruteforce struct {
	size         int
	limit        int
	numEnabled   bool
	lowerEnabled bool
	upperEnabled bool

	wg sync.WaitGroup
}

// New returns BruteForce implementation (*bruteforce) if no error occurs.
func New(opts ...Option) (BruteForce, error) {
	bf := new(bruteforce)

	for _, opt := range append(defaultOpts, opts...) {
		if err := opt(bf); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}
	return bf, nil
}

func (bf *bruteforce) Do(ctx context.Context, fn func(pattern string)) error {
	return bf.do(ctx, bf.dict(), bf.size, fn)
}

// results creates password candidates from a dictionary of given characters.
func (bf *bruteforce) do(ctx context.Context, chars []string, size int, fn func(pattern string)) error {
	var cnt int64
	var body func(pos int, part string)
	body = func(pos int, part string) {
		if pos == size {
			fn(part)
			return
		}
		for _, char := range chars {
			if int(atomic.LoadInt64(&cnt)) >= bf.limit {
				select {
				case <-ctx.Done():
				default:
					body(pos+1, char+part)
				}
			} else {
				bf.wg.Add(1)
				atomic.AddInt64(&cnt, 1)
				go func(char string) {
					defer bf.wg.Done()
					defer atomic.AddInt64(&cnt, -1)
					select {
					case <-ctx.Done():
					default:
						body(pos+1, char+part)
					}
				}(char)
			}
		}
	}
	body(0, "")
	bf.wg.Wait()
	return ctx.Err()
}

// dict creates a dictionary of characters.
func (bf *bruteforce) dict() (chars []string) {
	if bf.numEnabled {
		chars = append(chars, number...)
	}

	if bf.lowerEnabled {
		chars = append(chars, lower...)
	}

	if bf.upperEnabled {
		chars = append(chars, upper...)
	}

	return
}

// dictionary

var (
	number = []string{
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}

	upper = []string{
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
		"K",
		"L",
		"M",
		"N",
		"O",
		"P",
		"Q",
		"R",
		"S",
		"T",
		"U",
		"V",
		"W",
		"X",
		"Y",
		"Z",
	}

	lower = []string{
		"a",
		"b",
		"c",
		"d",
		"e",
		"f",
		"g",
		"h",
		"i",
		"j",
		"k",
		"l",
		"m",
		"n",
		"o",
		"p",
		"q",
		"r",
		"s",
		"t",
		"u",
		"v",
		"w",
		"x",
		"y",
		"z",
	}
)
