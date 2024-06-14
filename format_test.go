package logi

import (
	"reflect"
	"testing"
)

func TestFormat(t *testing.T) {
	type args struct {
		v    any
		opts []OptionFormat
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Test Format with nil",
			args: args{
				v:    nil,
				opts: []OptionFormat{},
			},
			want: nil,
		},
		{
			name: "Test Format with nil and WithDiscardNil",
			args: args{
				v:    nil,
				opts: []OptionFormat{WithDiscardNil(true)},
			},
			want: nil,
		},
		{
			name: "Test Format with struct",
			args: args{
				v: struct {
					A     string `log:"false"`
					B     string `log:"true"`
					C     *string
					D     *string
					E     *string `log:"true"`
					Inner struct {
						A string `log:"false"`
						B string `log:"true"`
					}
				}{
					A: "A",
					B: "B",
					C: func(v string) *string { return &v }("C"),
					Inner: struct {
						A string `log:"false"`
						B string `log:"true"`
					}{
						B: "B",
					},
				},
				opts: []OptionFormat{},
			},
			want: struct {
				B     string `log:"true"`
				C     *string
				D     *string
				E     *string `log:"true"`
				Inner struct {
					B string `log:"true"`
				}
			}{
				B: "B",
				C: func(v string) *string { return &v }("C"),
				D: nil,
				E: nil,
				Inner: struct {
					B string `log:"true"`
				}{
					B: "B",
				},
			},
		},
		{
			name: "Test Format with struct and WithDiscardNil",
			args: args{
				v: struct {
					A string `log:"false"`
					B string `log:"true"`
					C *string
					D *string
					E *string `log:"true"`
				}{
					A: "A",
					B: "B",
					C: func(v string) *string { return &v }("C"),
				},
				opts: []OptionFormat{WithDiscardNil(true)},
			},
			want: struct {
				B string `log:"true"`
				C *string
			}{
				B: "B",
				C: func(v string) *string { return &v }("C"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.v, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Format() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
