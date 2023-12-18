package main

import (
	"reflect"
	"testing"
)

func Test_parseDomain(t *testing.T) {
	type args struct {
		domain string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "with google.com",
			args: args{
				domain: "google.com",
			},
			want: []byte("\x06google\x03com\x00"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseDomain(tt.args.domain); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
