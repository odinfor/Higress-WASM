package internal

import (
	"time"
)

const (
	HeaderDeny = 1
	CookieDeny = 2
)

type (
	HeaderDenyMap struct {
		enable bool
		deny   map[string]AccessInfo
	}

	CookieDenyMap struct {
		enable bool
		deny   map[string]AccessInfo
	}
)

type (
	// AccessInfo 请求控制本地存储
	AccessInfo struct {

		// 策略类型, HeaderDeny: 基于header策略防护, CookieDeny: 基于cookie策略防护
		DenyType int `json:"denyType"`

		// 策略防护采用的key
		DenyK string `json:"denyK"`

		// 策略防护采用的value
		DenyV string `json:"denyV"`

		// 是否处于锁定状态
		Blocked bool `json:"blocked"`

		// 被锁定的原因
		BlockReason string `json:"BlockReason"`

		// 被锁定后屏蔽请求秒数
		BlockSeconds int `json:"blockSeconds"`

		// 是否启用秒内访问策略限制
		secUse bool

		// 当前秒内访问策略存储
		sec secondStore

		// 是否启用分钟内访问策略限制
		minUse bool
		// 当前分钟内访问策略存储
		min minuteStore

		// 是否启用当天内访问策略限制
		dayUse bool

		// 当天内访问策略存储
		day dayStore
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

		// 被锁定的时间
		blockTime time.Time

		// 释放锁定时间
		freeTime time.Time
	}
)

func GetHDMap() *HeaderDenyMap {
	return &HeaderDenyMap{}
}

func GetCDMap() *CookieDenyMap {
	return &CookieDenyMap{}
}

func InitDeny() {
	rules := NewCCDenyDo().Rules()
	hdm := GetHDMap()
	cdm := GetCDMap()

	for _, v := range rules {
		var info AccessInfo
		info.BlockSeconds = v.BlockSeconds

		if v.QPS > 0 {
			info.secUse = true
			info.sec.top = v.QPS
		}

		if v.QPM > 0 {
			info.minUse = true
			info.min.top = v.QPM
		}

		if v.QPD > 0 {
			info.dayUse = true
			info.day.top = v.QPD
		}

		if len(v.Header) > 0 {
			info.DenyType = HeaderDeny
			info.DenyK = Header
			info.DenyV = v.Header
			hdm.enable = true
			hdm.deny[v.Header] = info
			continue
		}
		if len(v.Cookie) > 0 {
			info.DenyType = CookieDeny
			info.DenyK = Cookie
			info.DenyV = v.Cookie
			cdm.enable = true
			cdm.deny[v.Cookie] = info
			continue
		}
	}
}
