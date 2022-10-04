package models

import (
	"testing"
)

func Test_fixFilename(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "good filename",
			args: args{
				s: "image \t\nd-^    .png",
			},
			want: "image+%09%0Ad-%5E++++.png",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fixFilename(tt.args.s); got != tt.want {
				t.Errorf("fixFilename() = %v, want %v", got, tt.want)
			}
		})
	}
}
