package gobf

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
)

func Test_bruteforce_Do(t *testing.T) {
	type args struct {
		ctx context.Context
		fn  func(pattern string)
	}
	type field struct {
		size       int
		numEnabled bool
	}
	type test struct {
		name      string
		args      args
		field     field
		want      error
		checkFunc func() error
	}

	tests := []test{
		func() test {
			var found int32
			return test{
				name: "return nil and find 123 when numEnabled is true and size is 3",
				args: args{
					ctx: context.Background(),
					fn: func(pattern string) {
						if pattern == "123" {
							atomic.AddInt32(&found, 1)
						}
					},
				},
				field: field{
					size:       3,
					numEnabled: true,
				},
				want: nil,
				checkFunc: func() error {
					found := int(atomic.LoadInt32(&found))
					switch {
					case found == 0:
						return errors.New("not found")
					case found == 1:
						return nil
					default:
						return fmt.Errorf("found duplicates: %d", found)
					}
				},
			}
		}(),
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			bf := &bruteforce{
				size:       test.field.size,
				numEnabled: test.field.numEnabled,
			}

			bf.Do(test.args.ctx, test.args.fn)
			if test.checkFunc != nil {
				if err := test.checkFunc(); err != nil {
					t.Error(err)
				}
			}
		})
	}
}

func Test_bruteforce_dicts(t *testing.T) {
	t.Parallel()
	type field struct {
		numEnabled   bool
		lowerEnabled bool
		upperEnabled bool
	}
	type test struct {
		name  string
		field field
		want  []string
	}

	tests := []test{
		{
			name: "return numbers when numEnabled is true",
			field: field{
				numEnabled: true,
			},
			want: number,
		},
		{
			name: "return lower charactors slice when lowerEnabled is true",
			field: field{
				lowerEnabled: true,
			},
			want: lower,
		},
		{
			name: "return upper charactors slice when upperEnabled is true",
			field: field{
				upperEnabled: true,
			},
			want: upper,
		},
		{
			name: "return lower and upper charactors and numbers when upperEnabled is true",
			field: field{
				lowerEnabled: true,
				upperEnabled: true,
				numEnabled:   true,
			},
			want: append(number, append(lower, upper...)...),
		},
		{
			name:  "return empty when all are false",
			field: field{},
			want:  []string{},
		},
	}

	for _, tc := range tests {
		test := tc
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			bf := &bruteforce{
				numEnabled:   test.field.numEnabled,
				lowerEnabled: test.field.lowerEnabled,
				upperEnabled: test.field.upperEnabled,
			}
			got := bf.dict()

			if len(test.want) != len(got) {
				t.Errorf("size is wrong. want: %d, but got: %d", len(test.want), len(got))
				return
			}

			for i, want := range test.want {
				if want, got := want, got[i]; want != got {
					t.Errorf("want: %v, but got: %v", want, got)
					return
				}
			}
		})
	}
}
