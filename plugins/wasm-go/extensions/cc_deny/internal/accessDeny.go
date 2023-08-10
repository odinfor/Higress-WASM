package internal

import (
	"time"
)

type Deny struct {
	as    *AccessStore
	allow bool
	err   error
}

type DenyDo interface {
	Access(now time.Time)
	Allow() bool
	Err() error
	Clear()
}

func NewDeny() DenyDo {
	return &Deny{}
}

func (d *Deny) name() {

}

func (d *Deny) Access(now time.Time) {
	// 检查今日访问次数限制

	// 检查分钟访问次数限制

	// 检查秒钟访问次数限制

}

func (d *Deny) Allow() bool {
	return d.allow
}

func (d *Deny) Err() error {
	if d.err == nil {
		return nil
	}
	return d.err
}

// Clear 防护缓存清理
func (d *Deny) Clear() {

}
