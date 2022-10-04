package api

import (
	"testing"
)

func Test_getMysqlQueryFromQueryString(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "q = ?",
			args: args{
				q: "sum_damage_amount+%3D+100",
			},
			want: "sum_damage_amount = ?",
		},
		{
			name: "sum_damage_amount+BETWEEN+100+AND+10000000",
			args: args{
				q: "sum_damage_amount+BETWEEN+100+AND+10000000",
			},
			want: "sum_damage_amount BETWEEN ? AND ?",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMysqlQueryFromQueryString(tt.args.q, nil); got != tt.want {
				t.Errorf("getMysqlQueryFromQueryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidRawQuery(t *testing.T) {
	type args struct {
		q string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "q1",
			args: args{
				q: "sum_damage_amount+%3D+1000000",
			},
			want: true,
		},
		{
			name: "q1",
			args: args{
				q: "sum_damage_amount+BETWEEN+10000+AND+100000000",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidRawQuery(tt.args.q); got != tt.want {
				t.Errorf("isValidRawQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
