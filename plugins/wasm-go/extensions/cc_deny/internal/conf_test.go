package internal

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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashKey(tt.args.userAgent, tt.args.XForwarded); got != tt.want {
				fmt.Println(got)
				t.Errorf("hashKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
