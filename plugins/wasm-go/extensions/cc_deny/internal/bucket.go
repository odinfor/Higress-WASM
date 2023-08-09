package internal

import (
	"cc_deny/internal/types"
	"context"
	xrate "golang.org/x/time/rate"
	"sync"
	"time"
)

/* 实现一个令牌桶 */

const (
	tokenFormat     = "{%s}.tokens"
	timestampFormat = "{%s}.ts"
	pingInterval    = time.Millisecond * 100
)

// A TokenLimiter controls how frequently events are allowed to happen with in one second.
type TokenLimiter struct {
	// 令牌速率
	rate int

	// 令牌桶最大令牌数量
	burst int

	// 令牌key存储
	store *TokenQ

	//tokenKey     string
	//timestampKey string

	rescueLock sync.Mutex

	// 限制器
	rescueLimiter *xrate.Limiter
}

// NewTokenLimiter returns a new TokenLimiter that allows events up to rate and permits
// bursts of at most burst tokens.
func NewTokenLimiter() *TokenLimiter {

	c := types.NewConfDo()

	return &TokenLimiter{
		rate:  c.LimiterConf().Rate,
		burst: c.LimiterConf().Burst,
		store: NewTQ(c.LimiterConf().Burst),
		rescueLimiter: xrate.NewLimiter(
			xrate.Every(time.Second/time.Duration(c.LimiterConf().Rate)), c.LimiterConf().Burst,
		),
	}
}

func (t *TokenLimiter) Wait(ctx context.Context) error {
	return t.rescueLimiter.Wait(ctx)
}

func (t *TokenLimiter) WaitN(ctx context.Context, n int) error {
	return t.rescueLimiter.WaitN(ctx, n)
}

func (t *TokenLimiter) Allow() bool {
	return t.rescueLimiter.Allow()
}

func (t *TokenLimiter) AllowN(now time.Time, n int) bool {
	return t.rescueLimiter.AllowN(now, n)
}
