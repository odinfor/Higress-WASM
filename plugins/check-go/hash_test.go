package check_go

import (
	"fmt"
	"testing"
)

func Test_hashKey(t *testing.T) {
	type args struct {
		userAgent  string
		XForwarded string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test hash",
			args: args{
				userAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
				XForwarded: "172.16.1.6:80",
			},
			want: "2504fda9c530eb9774445537e4b833291fe2b9428f4f7aa028a8511236fa3ad7",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hashKey(tt.args.userAgent, tt.args.XForwarded)
			fmt.Println(got)
			if got != tt.want {
				t.Errorf("hashKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bloomKey(t *testing.T) {
	type args struct {
		userAgent  string
		XForwarded string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test hash",
			args: args{
				userAgent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36",
				XForwarded: "172.16.1.6:80",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bloomKey(tt.args.userAgent, tt.args.XForwarded); got != tt.want {
				t.Errorf("bloomKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
