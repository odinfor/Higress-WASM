package cc_deny

import (
	"fmt"
	"github.com/redis/go-redis/v9"
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
	rate           int
	burst          int
	store          *redis.Client
	tokenKey       string
	timestampKey   string
	rescueLock     sync.Mutex
	redisAlive     uint32
	monitorStarted bool
	rescueLimiter  *xrate.Limiter
}

// NewTokenLimiter returns a new TokenLimiter that allows events up to rate and permits
// bursts of at most burst tokens.
func NewTokenLimiter(rate, burst int, store *redis.Client, key string) *TokenLimiter {
	tokenKey := fmt.Sprintf(tokenFormat, key)
	timestampKey := fmt.Sprintf(timestampFormat, key)

	return &TokenLimiter{
		rate:          rate,
		burst:         burst,
		store:         store,
		tokenKey:      tokenKey,
		timestampKey:  timestampKey,
		redisAlive:    1,
		rescueLimiter: xrate.NewLimiter(xrate.Every(time.Second/time.Duration(rate)), burst),
	}
}
