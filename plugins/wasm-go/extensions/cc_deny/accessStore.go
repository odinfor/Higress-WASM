package cc_deny

import "time"

type accessStore struct {
	store map[string]access
}

type access struct {
	// 用户上一次访问时间
	lastAcc time.Time

	// 访问是否被锁定
	blocked bool

	// 统计秒内起始时间
	baseSec time.Time

	// 秒级别内的访问次数
	accInSec int

	// 统计分钟内起始时间
	baseMin time.Time

	// 分级别内的访问次数
	accInMin int

	// 统计天内访问次数起始时间
	baseDay time.Time

	// 天级别内的访问次数
	accInDay int

	// 锁定访问多久,单位秒
	blockSec int

	// 被锁定的时间
	blockTime time.Time

	// 释放锁定时间
	freeTime time.Time
}
