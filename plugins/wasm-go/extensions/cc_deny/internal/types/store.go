package types

import (
	"sync"
	"time"
)

type (
	// AccessStore 请求控制本地存储
	AccessStore struct {
		mutex sync.Mutex
		sec   secondStore
		min   minuteStore
		day   dayStore
	}

	secondStore struct {

		// 访问是否被锁定
		blocked bool

		// 每秒最多访问次数
		top int

		// 用户上一次访问时间
		lastAcc time.Time

		// 当前统计周期的起始时间
		cycleStart time.Time

		// 当前统计问周期内的访问次数
		account int

		// 每次被锁定后屏蔽请求秒数
		blockNum int

		// 被锁定的时间
		blockTime time.Time

		// 释放锁定时间
		freeTime time.Time
	}

	minuteStore struct {

		// 访问是否被锁定
		blocked bool

		// 每分钟最多访问次数
		top int

		// 用户上一次访问时间
		lastAcc time.Time

		// 当前统计周期的起始时间
		cycleStart time.Time

		// 当前统计问周期内的访问次数
		account int

		// 每次被锁定后屏蔽请求分钟数
		blockNum int

		// 被锁定的时间
		blockTime time.Time

		// 释放锁定时间
		freeTime time.Time
	}

	dayStore struct {

		// 访问是否被锁定
		blocked bool

		// 每天最多访问次数
		top int

		// 用户上一次访问时间
		lastAcc time.Time

		// 当前统计周期的起始时间
		cycleStart time.Time

		// 当前统计问周期内的访问次数
		account int

		// 每次被锁定后屏蔽请求天数
		blockNum int

		// 被锁定的时间
		blockTime time.Time

		// 释放锁定时间
		freeTime time.Time
	}
)

func NewStore() *AccessStore {
	//c := NewConfDo()
	return &AccessStore{}
}
