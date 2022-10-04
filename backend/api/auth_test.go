package api

import "testing"

func Test_replaceProvince(t *testing.T) {
	type args struct {
		cur string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "t1",
			args: args{cur: "061,062"},
			want: "061,062,811,810",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceProvince(tt.args.cur); got != tt.want {
				t.Errorf("replaceProvince() = %v, want %v", got, tt.want)
			}
		})
	}
}
