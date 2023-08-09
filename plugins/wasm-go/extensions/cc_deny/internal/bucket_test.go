package internal

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewTokenLimiter(t *testing.T) {
	type args struct {
		rate  int
		burst int
		key   string
	}
	tests := []struct {
		name string
		args args
		want *TokenLimiter
	}{
		// TODO: Add test cases.
		{
			name: "test rate way",
			args: args{
				key:   "rate",
				rate:  10,
				burst: 20,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lim := NewTokenLimiter()

			for i := 0; i < 30; i++ {
				if lim.Allow() {
					fmt.Printf("success, %f \n", lim.rescueLimiter.Tokens())
				} else {
					fmt.Printf("fail, %f \n", lim.rescueLimiter.Tokens())
				}
			}

			time.Sleep(time.Second)

			fmt.Printf("wait time out, start get limiter \n")

			for i := 0; i < 5; i++ {
				if lim.Allow() {
					fmt.Printf("success, %f \n", lim.rescueLimiter.Tokens())
				} else {
					fmt.Printf("fail, %f \n", lim.rescueLimiter.Tokens())
				}
			}

			fmt.Println("use wait get 10")
			ctx, _ := context.WithTimeout(context.Background(), time.Second*30)
			if err := lim.WaitN(ctx, 10); err != nil {
				fmt.Println("waitN found err")
			}

		})
	}
}
